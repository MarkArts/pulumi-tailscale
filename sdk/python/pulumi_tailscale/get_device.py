# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDeviceResult',
    'AwaitableGetDeviceResult',
    'get_device',
    'get_device_output',
]

@pulumi.output_type
class GetDeviceResult:
    """
    A collection of values returned by getDevice.
    """
    def __init__(__self__, addresses=None, id=None, name=None, tags=None, user=None, wait_for=None):
        if addresses and not isinstance(addresses, list):
            raise TypeError("Expected argument 'addresses' to be a list")
        pulumi.set(__self__, "addresses", addresses)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if tags and not isinstance(tags, list):
            raise TypeError("Expected argument 'tags' to be a list")
        pulumi.set(__self__, "tags", tags)
        if user and not isinstance(user, str):
            raise TypeError("Expected argument 'user' to be a str")
        pulumi.set(__self__, "user", user)
        if wait_for and not isinstance(wait_for, str):
            raise TypeError("Expected argument 'wait_for' to be a str")
        pulumi.set(__self__, "wait_for", wait_for)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[str]:
        """
        The list of device's IPs
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the device
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> Sequence[str]:
        """
        The tags applied to the device
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def user(self) -> str:
        """
        The user associated with the device
        """
        return pulumi.get(self, "user")

    @property
    @pulumi.getter(name="waitFor")
    def wait_for(self) -> Optional[str]:
        """
        If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
        """
        return pulumi.get(self, "wait_for")


class AwaitableGetDeviceResult(GetDeviceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeviceResult(
            addresses=self.addresses,
            id=self.id,
            name=self.name,
            tags=self.tags,
            user=self.user,
            wait_for=self.wait_for)


def get_device(name: Optional[str] = None,
               wait_for: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeviceResult:
    """
    The device data source describes a single device in a tailnet

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    sample_device = tailscale.get_device(name="user1-device.example.com",
        wait_for="60s")
    ```


    :param str name: The name of the device
    :param str wait_for: If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['waitFor'] = wait_for
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale:index/getDevice:getDevice', __args__, opts=opts, typ=GetDeviceResult).value

    return AwaitableGetDeviceResult(
        addresses=pulumi.get(__ret__, 'addresses'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        tags=pulumi.get(__ret__, 'tags'),
        user=pulumi.get(__ret__, 'user'),
        wait_for=pulumi.get(__ret__, 'wait_for'))


@_utilities.lift_output_func(get_device)
def get_device_output(name: Optional[pulumi.Input[str]] = None,
                      wait_for: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeviceResult]:
    """
    The device data source describes a single device in a tailnet

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    sample_device = tailscale.get_device(name="user1-device.example.com",
        wait_for="60s")
    ```


    :param str name: The name of the device
    :param str wait_for: If specified, the provider will make multiple attempts to obtain the data source until the wait_for duration is reached. Retries are made every second so this value should be greater than 1s
    """
    ...
