// Copyright 2014 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package firewaller

import (
	"sort"

	"github.com/juju/errors"
	"github.com/juju/loggo"
	"github.com/juju/names/v4"

	"github.com/juju/juju/apiserver/common"
	"github.com/juju/juju/apiserver/common/cloudspec"
	"github.com/juju/juju/apiserver/common/firewall"
	apiservererrors "github.com/juju/juju/apiserver/errors"
	"github.com/juju/juju/apiserver/facade"
	"github.com/juju/juju/apiserver/params"
	corefirewall "github.com/juju/juju/core/firewall"
	"github.com/juju/juju/core/status"
	"github.com/juju/juju/state"
	"github.com/juju/juju/state/watcher"
)

var logger = loggo.GetLogger("juju.apiserver.firewaller")

// FirewallerAPIV3 provides access to the Firewaller v3 API facade.
type FirewallerAPIV3 struct {
	*common.LifeGetter
	*common.ModelWatcher
	*common.AgentEntityWatcher
	*common.UnitsWatcher
	*common.ModelMachinesWatcher
	*common.InstanceIdGetter
	cloudspec.CloudSpecAPI

	st                State
	resources         facade.Resources
	authorizer        facade.Authorizer
	accessUnit        common.GetAuthFunc
	accessApplication common.GetAuthFunc
	accessMachine     common.GetAuthFunc
	accessModel       common.GetAuthFunc
}

// FirewallerAPIV4 provides access to the Firewaller v4 API facade.
type FirewallerAPIV4 struct {
	*FirewallerAPIV3
	*common.ControllerConfigAPI
}

// FirewallerAPIV5 provides access to the Firewaller v5 API facade.
type FirewallerAPIV5 struct {
	*FirewallerAPIV4
}

// NewStateFirewallerAPIV3 creates a new server-side FirewallerAPIV3 facade.
func NewStateFirewallerAPIV3(context facade.Context) (*FirewallerAPIV3, error) {
	st := context.State()

	m, err := st.Model()
	if err != nil {
		return nil, errors.Trace(err)
	}

	cloudSpecAPI := cloudspec.NewCloudSpec(
		context.Resources(),
		cloudspec.MakeCloudSpecGetterForModel(st),
		cloudspec.MakeCloudSpecWatcherForModel(st),
		cloudspec.MakeCloudSpecCredentialWatcherForModel(st),
		cloudspec.MakeCloudSpecCredentialContentWatcherForModel(st),
		common.AuthFuncForTag(m.ModelTag()),
	)
	return NewFirewallerAPI(stateShim{st: st, State: firewall.StateShim(st, m)}, context.Resources(), context.Auth(), cloudSpecAPI)
}

// NewStateFirewallerAPIV4 creates a new server-side FirewallerAPIV4 facade.
func NewStateFirewallerAPIV4(context facade.Context) (*FirewallerAPIV4, error) {
	facadev3, err := NewStateFirewallerAPIV3(context)
	if err != nil {
		return nil, err
	}
	return &FirewallerAPIV4{
		ControllerConfigAPI: common.NewStateControllerConfig(context.State()),
		FirewallerAPIV3:     facadev3,
	}, nil
}

// NewStateFirewallerAPIV5 creates a new server-side FirewallerAPIV5 facade.
func NewStateFirewallerAPIV5(context facade.Context) (*FirewallerAPIV5, error) {
	facadev4, err := NewStateFirewallerAPIV4(context)
	if err != nil {
		return nil, err
	}
	return &FirewallerAPIV5{
		FirewallerAPIV4: facadev4,
	}, nil
}

// NewFirewallerAPI creates a new server-side FirewallerAPIV3 facade.
func NewFirewallerAPI(
	st State,
	resources facade.Resources,
	authorizer facade.Authorizer,
	cloudSpecAPI cloudspec.CloudSpecAPI,
) (*FirewallerAPIV3, error) {
	if !authorizer.AuthController() {
		// Firewaller must run as a controller.
		return nil, apiservererrors.ErrPerm
	}
	// Set up the various authorization checkers.
	accessModel := common.AuthFuncForTagKind(names.ModelTagKind)
	accessUnit := common.AuthFuncForTagKind(names.UnitTagKind)
	accessApplication := common.AuthFuncForTagKind(names.ApplicationTagKind)
	accessMachine := common.AuthFuncForTagKind(names.MachineTagKind)
	accessRelation := common.AuthFuncForTagKind(names.RelationTagKind)
	accessUnitApplicationOrMachineOrRelation := common.AuthAny(accessUnit, accessApplication, accessMachine, accessRelation)

	// Life() is supported for units, applications or machines.
	lifeGetter := common.NewLifeGetter(
		st,
		accessUnitApplicationOrMachineOrRelation,
	)
	// ModelConfig() and WatchForModelConfigChanges() are allowed
	// with unrestricted access.
	modelWatcher := common.NewModelWatcher(
		st,
		resources,
		authorizer,
	)
	// Watch() is supported for applications only.
	entityWatcher := common.NewAgentEntityWatcher(
		st,
		resources,
		accessApplication,
	)
	// WatchUnits() is supported for machines.
	unitsWatcher := common.NewUnitsWatcher(st,
		resources,
		accessMachine,
	)
	// WatchModelMachines() is allowed with unrestricted access.
	machinesWatcher := common.NewModelMachinesWatcher(
		st,
		resources,
		authorizer,
	)
	// InstanceId() is supported for machines.
	instanceIdGetter := common.NewInstanceIdGetter(
		st,
		accessMachine,
	)

	return &FirewallerAPIV3{
		LifeGetter:           lifeGetter,
		ModelWatcher:         modelWatcher,
		AgentEntityWatcher:   entityWatcher,
		UnitsWatcher:         unitsWatcher,
		ModelMachinesWatcher: machinesWatcher,
		InstanceIdGetter:     instanceIdGetter,
		CloudSpecAPI:         cloudSpecAPI,
		st:                   st,
		resources:            resources,
		authorizer:           authorizer,
		accessUnit:           accessUnit,
		accessApplication:    accessApplication,
		accessMachine:        accessMachine,
		accessModel:          accessModel,
	}, nil
}

// WatchOpenedPorts returns a new StringsWatcher for each given
// model tag.
func (f *FirewallerAPIV3) WatchOpenedPorts(args params.Entities) (params.StringsWatchResults, error) {
	result := params.StringsWatchResults{
		Results: make([]params.StringsWatchResult, len(args.Entities)),
	}
	if len(args.Entities) == 0 {
		return result, nil
	}
	canWatch, err := f.accessModel()
	if err != nil {
		return params.StringsWatchResults{}, errors.Trace(err)
	}
	for i, entity := range args.Entities {
		tag, err := names.ParseTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
			continue
		}
		if !canWatch(tag) {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
			continue
		}
		watcherId, initial, err := f.watchOneModelOpenedPorts(tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		result.Results[i].StringsWatcherId = watcherId
		result.Results[i].Changes = initial
	}
	return result, nil
}

func (f *FirewallerAPIV3) watchOneModelOpenedPorts(tag names.Tag) (string, []string, error) {
	// NOTE: tag is ignored, as there is only one model in the
	// state DB. Once this changes, change the code below accordingly.
	watch := f.st.WatchOpenedPorts()
	// Consume the initial event and forward it to the result.
	if changes, ok := <-watch.Changes(); ok {
		return f.resources.Register(watch), changes, nil
	}
	return "", nil, watcher.EnsureErr(watch)
}

// GetMachinePorts returns the port ranges opened on a machine across all
// subnets as a map mapping port ranges to the tags of the units that opened
// them.
func (f *FirewallerAPIV3) GetMachinePorts(args params.MachinePortsParams) (params.MachinePortsResults, error) {
	result := params.MachinePortsResults{
		Results: make([]params.MachinePortsResult, len(args.Params)),
	}
	canAccess, err := f.accessMachine()
	if err != nil {
		return params.MachinePortsResults{}, err
	}
	for i, param := range args.Params {
		machineTag, err := names.ParseMachineTag(param.MachineTag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		// Pre 2.9 controllers always open ports in all subnets. The
		// per-subnet functionality was implemented by the controller
		// but never exposed via the API. As such, we can change this
		// method so that always returns *all* open port ranges and
		// simply return an error if a subnet is provided
		if param.SubnetTag != "" {
			if _, err = names.ParseSubnetTag(param.SubnetTag); err == nil {
				err = errors.NotSupportedf("retrieving machine ports for specific subnets")
			}
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		machine, err := f.getMachine(canAccess, machineTag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		machPortRanges, err := machine.OpenedPortRanges()
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		// Emulate old behavior and return opened ports for all endpoints.
		var rangeList []params.MachinePortRange
		for unitName, unitPortRanges := range machPortRanges.ByUnit() {
			unitTag := names.NewUnitTag(unitName).String()
			for _, pr := range unitPortRanges.UniquePortRanges() {
				rangeList = append(rangeList, params.MachinePortRange{
					UnitTag:   unitTag,
					PortRange: params.FromNetworkPortRange(pr),
				})
			}
		}

		sort.Slice(rangeList, func(i, j int) bool {
			return rangeList[i].PortRange.NetworkPortRange().LessThan(
				rangeList[j].PortRange.NetworkPortRange(),
			)
		})
		result.Results[i].Ports = rangeList
	}
	return result, nil
}

// GetMachineActiveSubnets returns the tags of the all subnets that each machine
// (in args) has open ports on.
func (f *FirewallerAPIV3) GetMachineActiveSubnets(args params.Entities) (params.StringsResults, error) {
	result := params.StringsResults{
		Results: make([]params.StringsResult, len(args.Entities)),
	}
	canAccess, err := f.accessMachine()
	if err != nil {
		return params.StringsResults{}, err
	}
	for i, entity := range args.Entities {
		machineTag, err := names.ParseMachineTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		machine, err := f.getMachine(canAccess, machineTag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}

		// Pre 2.9 controllers always open ports in all subnets. If
		// at least one port range is open, return the wildcard subnet
		machPorts, err := machine.OpenedPortRanges()
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		if len(machPorts.UniquePortRanges()) != 0 {
			result.Results[i].Result = append(result.Results[i].Result, "")
		}
	}
	return result, nil
}

// GetExposed returns the exposed flag value for each given application.
func (f *FirewallerAPIV3) GetExposed(args params.Entities) (params.BoolResults, error) {
	result := params.BoolResults{
		Results: make([]params.BoolResult, len(args.Entities)),
	}
	canAccess, err := f.accessApplication()
	if err != nil {
		return params.BoolResults{}, err
	}
	for i, entity := range args.Entities {
		tag, err := names.ParseApplicationTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
			continue
		}
		application, err := f.getApplication(canAccess, tag)
		if err == nil {
			result.Results[i].Result = application.IsExposed()
		}
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}

// GetAssignedMachine returns the assigned machine tag (if any) for
// each given unit.
func (f *FirewallerAPIV3) GetAssignedMachine(args params.Entities) (params.StringResults, error) {
	result := params.StringResults{
		Results: make([]params.StringResult, len(args.Entities)),
	}
	canAccess, err := f.accessUnit()
	if err != nil {
		return params.StringResults{}, err
	}
	for i, entity := range args.Entities {
		tag, err := names.ParseUnitTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(apiservererrors.ErrPerm)
			continue
		}
		unit, err := f.getUnit(canAccess, tag)
		if err == nil {
			var machineId string
			machineId, err = unit.AssignedMachineId()
			if err == nil {
				result.Results[i].Result = names.NewMachineTag(machineId).String()
			}
		}
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}

func (f *FirewallerAPIV3) getEntity(canAccess common.AuthFunc, tag names.Tag) (state.Entity, error) {
	if !canAccess(tag) {
		return nil, apiservererrors.ErrPerm
	}
	return f.st.FindEntity(tag)
}

func (f *FirewallerAPIV3) getUnit(canAccess common.AuthFunc, tag names.UnitTag) (*state.Unit, error) {
	entity, err := f.getEntity(canAccess, tag)
	if err != nil {
		return nil, err
	}
	// The authorization function guarantees that the tag represents a
	// unit.
	return entity.(*state.Unit), nil
}

func (f *FirewallerAPIV3) getApplication(canAccess common.AuthFunc, tag names.ApplicationTag) (*state.Application, error) {
	entity, err := f.getEntity(canAccess, tag)
	if err != nil {
		return nil, err
	}
	// The authorization function guarantees that the tag represents a
	// application.
	return entity.(*state.Application), nil
}

func (f *FirewallerAPIV3) getMachine(canAccess common.AuthFunc, tag names.MachineTag) (*state.Machine, error) {
	entity, err := f.getEntity(canAccess, tag)
	if err != nil {
		return nil, err
	}
	// The authorization function guarantees that the tag represents a
	// machine.
	return entity.(*state.Machine), nil
}

// WatchEgressAddressesForRelations creates a watcher that notifies when addresses, from which
// connections will originate for the relation, change.
// Each event contains the entire set of addresses which are required for ingress for the relation.
func (f *FirewallerAPIV4) WatchEgressAddressesForRelations(relations params.Entities) (params.StringsWatchResults, error) {
	return firewall.WatchEgressAddressesForRelations(f.resources, f.st, relations)
}

// WatchIngressAddressesForRelations creates a watcher that returns the ingress networks
// that have been recorded against the specified relations.
func (f *FirewallerAPIV4) WatchIngressAddressesForRelations(relations params.Entities) (params.StringsWatchResults, error) {
	results := params.StringsWatchResults{
		make([]params.StringsWatchResult, len(relations.Entities)),
	}

	one := func(tag string) (id string, changes []string, _ error) {
		logger.Debugf("Watching ingress addresses for %+v from model %v", tag, f.st.ModelUUID())

		relationTag, err := names.ParseRelationTag(tag)
		if err != nil {
			return "", nil, errors.Trace(err)
		}
		rel, err := f.st.KeyRelation(relationTag.Id())
		if err != nil {
			return "", nil, errors.Trace(err)
		}
		w := rel.WatchRelationIngressNetworks()
		changes, ok := <-w.Changes()
		if !ok {
			return "", nil, apiservererrors.ServerError(watcher.EnsureErr(w))
		}
		return f.resources.Register(w), changes, nil
	}

	for i, e := range relations.Entities {
		watcherId, changes, err := one(e.Tag)
		if err != nil {
			results.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		results.Results[i].StringsWatcherId = watcherId
		results.Results[i].Changes = changes
	}
	return results, nil
}

// MacaroonForRelations returns the macaroon for the specified relations.
func (f *FirewallerAPIV4) MacaroonForRelations(args params.Entities) (params.MacaroonResults, error) {
	var result params.MacaroonResults
	result.Results = make([]params.MacaroonResult, len(args.Entities))
	for i, entity := range args.Entities {
		relationTag, err := names.ParseRelationTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		mac, err := f.st.GetMacaroon(relationTag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		result.Results[i].Result = mac
	}
	return result, nil
}

// SetRelationsStatus sets the status for the specified relations.
func (f *FirewallerAPIV4) SetRelationsStatus(args params.SetStatus) (params.ErrorResults, error) {
	var result params.ErrorResults
	result.Results = make([]params.ErrorResult, len(args.Entities))
	for i, entity := range args.Entities {
		relationTag, err := names.ParseRelationTag(entity.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		rel, err := f.st.KeyRelation(relationTag.Id())
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		err = rel.SetStatus(status.StatusInfo{
			Status:  status.Status(entity.Status),
			Message: entity.Info,
		})
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}

// FirewallRules returns the firewall rules for the specified well known service types.
func (f *FirewallerAPIV4) FirewallRules(args params.KnownServiceArgs) (params.ListFirewallRulesResults, error) {
	var result params.ListFirewallRulesResults
	for _, knownService := range args.KnownServices {
		rule, err := f.st.FirewallRule(corefirewall.WellKnownServiceType(knownService))
		if err != nil && !errors.IsNotFound(err) {
			return result, apiservererrors.ServerError(err)
		}
		if err != nil {
			continue
		}
		result.Rules = append(result.Rules, params.FirewallRule{
			KnownService:   knownService,
			WhitelistCIDRS: rule.WhitelistCIDRs(),
		})
	}
	return result, nil
}

// AreManuallyProvisioned returns whether each given entity is
// manually provisioned or not. Only machine tags are accepted.
func (f *FirewallerAPIV5) AreManuallyProvisioned(args params.Entities) (params.BoolResults, error) {
	result := params.BoolResults{
		Results: make([]params.BoolResult, len(args.Entities)),
	}
	canAccess, err := f.accessMachine()
	if err != nil {
		return result, err
	}
	for i, arg := range args.Entities {
		machineTag, err := names.ParseMachineTag(arg.Tag)
		if err != nil {
			result.Results[i].Error = apiservererrors.ServerError(err)
			continue
		}
		machine, err := f.getMachine(canAccess, machineTag)
		if err == nil {
			result.Results[i].Result, err = machine.IsManual()
		}
		result.Results[i].Error = apiservererrors.ServerError(err)
	}
	return result, nil
}
