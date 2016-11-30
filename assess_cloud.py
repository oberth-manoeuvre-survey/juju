#!/usr/bin/env python
from argparse import ArgumentParser
from textwrap import dedent
import yaml

from deploy_stack import (
    BootstrapManager,
    )
from fakejuju import (
    FakeBackend,
    FakeControllerState,
    )
from jujuconfig import get_juju_home
from jujupy import (
    EnvJujuClient,
    get_client_class,
    WaitMachineNotPresent,
    )
from utility import (
    add_basic_testing_arguments,
    configure_logging,
    )


def client_from_args(args):
    """Return client from args, as generated by parse_args.

    If the path given is FAKE, fake_juju_client() is used.  Otherwise, the
    client is determined based on the path and version.
    """
    if args.juju_bin == 'FAKE':
        client_class = EnvJujuClient
        controller_state = FakeControllerState()
        version = '2.0.0'
        backend = FakeBackend(controller_state, full_path=args.juju_bin,
                              version=version)
    else:
        version = EnvJujuClient.get_version(args.juju_bin)
        client_class = get_client_class(version)
        backend = None
    juju_home = get_juju_home()
    with open(args.clouds_file) as f:
        clouds = yaml.safe_load(f)
    juju_data = client_class.config_class.from_cloud_region(
        args.cloud, args.region, {}, clouds, juju_home)
    return client_class(juju_data, version, args.juju_bin, debug=args.debug,
                        soft_deadline=args.deadline, _backend=backend)


def assess_cloud_combined(bs_manager):
    """Assess several operations on a cloud.

    This tests bootstrap, deploy, remove-unit and destroy-controller.
    """
    client = bs_manager.client
    with bs_manager.booted_context(upload_tools=False):
        old_status = client.get_status()
        client.juju('deploy', ('ubuntu'))
        new_status = client.wait_for_started()
        new_machines = [k for k, v in new_status.iter_new_machines(old_status)]
        client.juju('remove-unit', 'ubuntu/0')
        new_status = client.wait_for([WaitMachineNotPresent(n)
                                      for n in new_machines])


def assess_cloud_kill_controller(bs_manager):
    client = bs_manager.client
    with bs_manager.booted_context(upload_tools=False):
        controller_client = client.get_controller_client()
        controller_client.juju('run', (
            '--machine', '0', 'sudo service jujud-machine-0 stop'),
            check=False)
        bs_manager.has_controller = False


def parse_args(args):
    parser = ArgumentParser(description=dedent("""\
        Test a specified cloud.

        This tests basic provider operations and charm store support.

        The cloud.yaml file must be provided, followed by the name of the
        cloud to test.
        """))
    subparsers = parser.add_subparsers(dest='test')
    for test in ['combined', 'kill-controller']:
        subparser = subparsers.add_parser(test)
        subparser.add_argument('clouds_file',
                               help='A clouds.yaml file to use for testing.')
        subparser.add_argument('cloud', help='Specific cloud to test.')
        add_basic_testing_arguments(subparser, env=False)
    return parser.parse_args(args)


def main():
    args = parse_args(None)
    configure_logging(args.verbose)
    client = client_from_args(args)
    bs_manager = BootstrapManager.from_client(args, client)
    if args.test == 'combined':
        assess_cloud_combined(bs_manager)
    else:
        assess_cloud_kill_controller(bs_manager)


if __name__ == '__main__':
    main()
