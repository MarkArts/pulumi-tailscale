# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetDevicesDeviceResult',
]

@pulumi.output_type
class GetDevicesDeviceResult(dict):
    def __init__(__self__, *,
                 addresses: Sequence[str],
                 id: str,
                 name: str,
                 user: str):
        """
        :param Sequence[str] addresses: Tailscale IPs for the device
        :param str id: The unique identifier of the device
        :param str name: The name of the device
        :param str user: The user associated with the device
        """
        pulumi.set(__self__, "addresses", addresses)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[str]:
        """
        Tailscale IPs for the device
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The unique identifier of the device
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
    def user(self) -> str:
        """
        The user associated with the device
        """
        return pulumi.get(self, "user")


