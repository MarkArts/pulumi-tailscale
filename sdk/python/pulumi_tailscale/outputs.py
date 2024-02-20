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
    'GetDevicesDeviceResult',
]

@pulumi.output_type
class GetDevicesDeviceResult(dict):
    def __init__(__self__, *,
                 addresses: Sequence[str],
                 hostname: str,
                 id: str,
                 name: str,
                 tags: Sequence[str],
                 user: str):
        """
        :param Sequence[str] addresses: The list of device's IPs
        :param str hostname: The short hostname of the device
        :param str id: The unique identifier of the device
        :param str name: The full name of the device (e.g. `hostname.domain.ts.net`)
        :param Sequence[str] tags: The tags applied to the device
        :param str user: The user associated with the device
        """
        pulumi.set(__self__, "addresses", addresses)
        pulumi.set(__self__, "hostname", hostname)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "tags", tags)
        pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def addresses(self) -> Sequence[str]:
        """
        The list of device's IPs
        """
        return pulumi.get(self, "addresses")

    @property
    @pulumi.getter
    def hostname(self) -> str:
        """
        The short hostname of the device
        """
        return pulumi.get(self, "hostname")

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
        The full name of the device (e.g. `hostname.domain.ts.net`)
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


