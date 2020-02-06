// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/juju/juju/apiserver/facades/agent/instancemutater (interfaces: ModelCache,ModelCacheMachine,ModelCacheApplication,ModelCacheUnit,ModelCacheCharm)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	instancemutater "github.com/juju/juju/apiserver/facades/agent/instancemutater"
	cache "github.com/juju/juju/core/cache"
	instance "github.com/juju/juju/core/instance"
	lxdprofile "github.com/juju/juju/core/lxdprofile"
)

// MockModelCache is a mock of ModelCache interface
type MockModelCache struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheMockRecorder
}

// MockModelCacheMockRecorder is the mock recorder for MockModelCache
type MockModelCacheMockRecorder struct {
	mock *MockModelCache
}

// NewMockModelCache creates a new mock instance
func NewMockModelCache(ctrl *gomock.Controller) *MockModelCache {
	mock := &MockModelCache{ctrl: ctrl}
	mock.recorder = &MockModelCacheMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCache) EXPECT() *MockModelCacheMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockModelCache) Application(arg0 string) (instancemutater.ModelCacheApplication, error) {
	ret := m.ctrl.Call(m, "Application", arg0)
	ret0, _ := ret[0].(instancemutater.ModelCacheApplication)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Application indicates an expected call of Application
func (mr *MockModelCacheMockRecorder) Application(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockModelCache)(nil).Application), arg0)
}

// Charm mocks base method
func (m *MockModelCache) Charm(arg0 string) (instancemutater.ModelCacheCharm, error) {
	ret := m.ctrl.Call(m, "Charm", arg0)
	ret0, _ := ret[0].(instancemutater.ModelCacheCharm)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Charm indicates an expected call of Charm
func (mr *MockModelCacheMockRecorder) Charm(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Charm", reflect.TypeOf((*MockModelCache)(nil).Charm), arg0)
}

// Machine mocks base method
func (m *MockModelCache) Machine(arg0 string) (instancemutater.ModelCacheMachine, error) {
	ret := m.ctrl.Call(m, "Machine", arg0)
	ret0, _ := ret[0].(instancemutater.ModelCacheMachine)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Machine indicates an expected call of Machine
func (mr *MockModelCacheMockRecorder) Machine(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Machine", reflect.TypeOf((*MockModelCache)(nil).Machine), arg0)
}

// Name mocks base method
func (m *MockModelCache) Name() string {
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockModelCacheMockRecorder) Name() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModelCache)(nil).Name))
}

// WatchMachines mocks base method
func (m *MockModelCache) WatchMachines() (cache.StringsWatcher, error) {
	ret := m.ctrl.Call(m, "WatchMachines")
	ret0, _ := ret[0].(cache.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchMachines indicates an expected call of WatchMachines
func (mr *MockModelCacheMockRecorder) WatchMachines() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchMachines", reflect.TypeOf((*MockModelCache)(nil).WatchMachines))
}

// MockModelCacheMachine is a mock of ModelCacheMachine interface
type MockModelCacheMachine struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheMachineMockRecorder
}

// MockModelCacheMachineMockRecorder is the mock recorder for MockModelCacheMachine
type MockModelCacheMachineMockRecorder struct {
	mock *MockModelCacheMachine
}

// NewMockModelCacheMachine creates a new mock instance
func NewMockModelCacheMachine(ctrl *gomock.Controller) *MockModelCacheMachine {
	mock := &MockModelCacheMachine{ctrl: ctrl}
	mock.recorder = &MockModelCacheMachineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCacheMachine) EXPECT() *MockModelCacheMachineMockRecorder {
	return m.recorder
}

// CharmProfiles mocks base method
func (m *MockModelCacheMachine) CharmProfiles() []string {
	ret := m.ctrl.Call(m, "CharmProfiles")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CharmProfiles indicates an expected call of CharmProfiles
func (mr *MockModelCacheMachineMockRecorder) CharmProfiles() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmProfiles", reflect.TypeOf((*MockModelCacheMachine)(nil).CharmProfiles))
}

// ContainerType mocks base method
func (m *MockModelCacheMachine) ContainerType() instance.ContainerType {
	ret := m.ctrl.Call(m, "ContainerType")
	ret0, _ := ret[0].(instance.ContainerType)
	return ret0
}

// ContainerType indicates an expected call of ContainerType
func (mr *MockModelCacheMachineMockRecorder) ContainerType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerType", reflect.TypeOf((*MockModelCacheMachine)(nil).ContainerType))
}

// InstanceId mocks base method
func (m *MockModelCacheMachine) InstanceId() (instance.Id, error) {
	ret := m.ctrl.Call(m, "InstanceId")
	ret0, _ := ret[0].(instance.Id)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InstanceId indicates an expected call of InstanceId
func (mr *MockModelCacheMachineMockRecorder) InstanceId() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InstanceId", reflect.TypeOf((*MockModelCacheMachine)(nil).InstanceId))
}

// Units mocks base method
func (m *MockModelCacheMachine) Units() ([]instancemutater.ModelCacheUnit, error) {
	ret := m.ctrl.Call(m, "Units")
	ret0, _ := ret[0].([]instancemutater.ModelCacheUnit)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Units indicates an expected call of Units
func (mr *MockModelCacheMachineMockRecorder) Units() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Units", reflect.TypeOf((*MockModelCacheMachine)(nil).Units))
}

// WatchContainers mocks base method
func (m *MockModelCacheMachine) WatchContainers() (cache.StringsWatcher, error) {
	ret := m.ctrl.Call(m, "WatchContainers")
	ret0, _ := ret[0].(cache.StringsWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchContainers indicates an expected call of WatchContainers
func (mr *MockModelCacheMachineMockRecorder) WatchContainers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchContainers", reflect.TypeOf((*MockModelCacheMachine)(nil).WatchContainers))
}

// WatchLXDProfileVerificationNeeded mocks base method
func (m *MockModelCacheMachine) WatchLXDProfileVerificationNeeded() (cache.NotifyWatcher, error) {
	ret := m.ctrl.Call(m, "WatchLXDProfileVerificationNeeded")
	ret0, _ := ret[0].(cache.NotifyWatcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WatchLXDProfileVerificationNeeded indicates an expected call of WatchLXDProfileVerificationNeeded
func (mr *MockModelCacheMachineMockRecorder) WatchLXDProfileVerificationNeeded() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WatchLXDProfileVerificationNeeded", reflect.TypeOf((*MockModelCacheMachine)(nil).WatchLXDProfileVerificationNeeded))
}

// MockModelCacheApplication is a mock of ModelCacheApplication interface
type MockModelCacheApplication struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheApplicationMockRecorder
}

// MockModelCacheApplicationMockRecorder is the mock recorder for MockModelCacheApplication
type MockModelCacheApplicationMockRecorder struct {
	mock *MockModelCacheApplication
}

// NewMockModelCacheApplication creates a new mock instance
func NewMockModelCacheApplication(ctrl *gomock.Controller) *MockModelCacheApplication {
	mock := &MockModelCacheApplication{ctrl: ctrl}
	mock.recorder = &MockModelCacheApplicationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCacheApplication) EXPECT() *MockModelCacheApplicationMockRecorder {
	return m.recorder
}

// CharmURL mocks base method
func (m *MockModelCacheApplication) CharmURL() string {
	ret := m.ctrl.Call(m, "CharmURL")
	ret0, _ := ret[0].(string)
	return ret0
}

// CharmURL indicates an expected call of CharmURL
func (mr *MockModelCacheApplicationMockRecorder) CharmURL() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CharmURL", reflect.TypeOf((*MockModelCacheApplication)(nil).CharmURL))
}

// MockModelCacheUnit is a mock of ModelCacheUnit interface
type MockModelCacheUnit struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheUnitMockRecorder
}

// MockModelCacheUnitMockRecorder is the mock recorder for MockModelCacheUnit
type MockModelCacheUnitMockRecorder struct {
	mock *MockModelCacheUnit
}

// NewMockModelCacheUnit creates a new mock instance
func NewMockModelCacheUnit(ctrl *gomock.Controller) *MockModelCacheUnit {
	mock := &MockModelCacheUnit{ctrl: ctrl}
	mock.recorder = &MockModelCacheUnitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCacheUnit) EXPECT() *MockModelCacheUnitMockRecorder {
	return m.recorder
}

// Application mocks base method
func (m *MockModelCacheUnit) Application() string {
	ret := m.ctrl.Call(m, "Application")
	ret0, _ := ret[0].(string)
	return ret0
}

// Application indicates an expected call of Application
func (mr *MockModelCacheUnitMockRecorder) Application() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Application", reflect.TypeOf((*MockModelCacheUnit)(nil).Application))
}

// MockModelCacheCharm is a mock of ModelCacheCharm interface
type MockModelCacheCharm struct {
	ctrl     *gomock.Controller
	recorder *MockModelCacheCharmMockRecorder
}

// MockModelCacheCharmMockRecorder is the mock recorder for MockModelCacheCharm
type MockModelCacheCharmMockRecorder struct {
	mock *MockModelCacheCharm
}

// NewMockModelCacheCharm creates a new mock instance
func NewMockModelCacheCharm(ctrl *gomock.Controller) *MockModelCacheCharm {
	mock := &MockModelCacheCharm{ctrl: ctrl}
	mock.recorder = &MockModelCacheCharmMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModelCacheCharm) EXPECT() *MockModelCacheCharmMockRecorder {
	return m.recorder
}

// LXDProfile mocks base method
func (m *MockModelCacheCharm) LXDProfile() lxdprofile.Profile {
	ret := m.ctrl.Call(m, "LXDProfile")
	ret0, _ := ret[0].(lxdprofile.Profile)
	return ret0
}

// LXDProfile indicates an expected call of LXDProfile
func (mr *MockModelCacheCharmMockRecorder) LXDProfile() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LXDProfile", reflect.TypeOf((*MockModelCacheCharm)(nil).LXDProfile))
}
