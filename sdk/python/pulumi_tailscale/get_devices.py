# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetDevicesResult',
    'AwaitableGetDevicesResult',
    'get_devices',
    'get_devices_output',
]

@pulumi.output_type
class GetDevicesResult:
    """
    A collection of values returned by getDevices.
    """
    def __init__(__self__, devices=None, id=None, name_prefix=None):
        if devices and not isinstance(devices, list):
            raise TypeError("Expected argument 'devices' to be a list")
        pulumi.set(__self__, "devices", devices)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name_prefix and not isinstance(name_prefix, str):
            raise TypeError("Expected argument 'name_prefix' to be a str")
        pulumi.set(__self__, "name_prefix", name_prefix)

    @property
    @pulumi.getter
    def devices(self) -> Sequence['outputs.GetDevicesDeviceResult']:
        """
        The list of devices in the tailnet
        """
        return pulumi.get(self, "devices")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> Optional[str]:
        """
        Filters the device list to elements whose name has the provided prefix
        """
        return pulumi.get(self, "name_prefix")


class AwaitableGetDevicesResult(GetDevicesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDevicesResult(
            devices=self.devices,
            id=self.id,
            name_prefix=self.name_prefix)


def get_devices(name_prefix: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDevicesResult:
    """
    The devices data source describes a list of devices in a tailnet

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    sample_devices = tailscale.get_devices(name_prefix="example-")
    ```


    :param str name_prefix: Filters the device list to elements whose name has the provided prefix
    """
    __args__ = dict()
    __args__['namePrefix'] = name_prefix
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale:index/getDevices:getDevices', __args__, opts=opts, typ=GetDevicesResult).value

    return AwaitableGetDevicesResult(
        devices=pulumi.get(__ret__, 'devices'),
        id=pulumi.get(__ret__, 'id'),
        name_prefix=pulumi.get(__ret__, 'name_prefix'))


@_utilities.lift_output_func(get_devices)
def get_devices_output(name_prefix: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDevicesResult]:
    """
    The devices data source describes a list of devices in a tailnet

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    sample_devices = tailscale.get_devices(name_prefix="example-")
    ```


    :param str name_prefix: Filters the device list to elements whose name has the provided prefix
    """
    ...
