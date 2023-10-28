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
    'Get4Via6Result',
    'AwaitableGet4Via6Result',
    'get4_via6',
    'get4_via6_output',
]

@pulumi.output_type
class Get4Via6Result:
    """
    A collection of values returned by get4Via6.
    """
    def __init__(__self__, cidr=None, id=None, ipv6=None, site=None):
        if cidr and not isinstance(cidr, str):
            raise TypeError("Expected argument 'cidr' to be a str")
        pulumi.set(__self__, "cidr", cidr)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv6 and not isinstance(ipv6, str):
            raise TypeError("Expected argument 'ipv6' to be a str")
        pulumi.set(__self__, "ipv6", ipv6)
        if site and not isinstance(site, int):
            raise TypeError("Expected argument 'site' to be a int")
        pulumi.set(__self__, "site", site)

    @property
    @pulumi.getter
    def cidr(self) -> str:
        """
        The IPv4 CIDR to map
        """
        return pulumi.get(self, "cidr")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def ipv6(self) -> str:
        """
        The 4via6 mapped address
        """
        return pulumi.get(self, "ipv6")

    @property
    @pulumi.getter
    def site(self) -> int:
        """
        Site ID (between 0 and 255)
        """
        return pulumi.get(self, "site")


class AwaitableGet4Via6Result(Get4Via6Result):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return Get4Via6Result(
            cidr=self.cidr,
            id=self.id,
            ipv6=self.ipv6,
            site=self.site)


def get4_via6(cidr: Optional[str] = None,
              site: Optional[int] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGet4Via6Result:
    """
    The 4via6 data source is calculates an IPv6 prefix for a given site ID and IPv4 CIDR. See Tailscale documentation for [4via6 subnets](https://tailscale.com/kb/1201/4via6-subnets/) for more details.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    example = tailscale.get4_via6(cidr="10.1.1.0/24",
        site=7)
    ```


    :param str cidr: The IPv4 CIDR to map
    :param int site: Site ID (between 0 and 255)
    """
    __args__ = dict()
    __args__['cidr'] = cidr
    __args__['site'] = site
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('tailscale:index/get4Via6:get4Via6', __args__, opts=opts, typ=Get4Via6Result).value

    return AwaitableGet4Via6Result(
        cidr=pulumi.get(__ret__, 'cidr'),
        id=pulumi.get(__ret__, 'id'),
        ipv6=pulumi.get(__ret__, 'ipv6'),
        site=pulumi.get(__ret__, 'site'))


@_utilities.lift_output_func(get4_via6)
def get4_via6_output(cidr: Optional[pulumi.Input[str]] = None,
                     site: Optional[pulumi.Input[int]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[Get4Via6Result]:
    """
    The 4via6 data source is calculates an IPv6 prefix for a given site ID and IPv4 CIDR. See Tailscale documentation for [4via6 subnets](https://tailscale.com/kb/1201/4via6-subnets/) for more details.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_tailscale as tailscale

    example = tailscale.get4_via6(cidr="10.1.1.0/24",
        site=7)
    ```


    :param str cidr: The IPv4 CIDR to map
    :param int site: Site ID (between 0 and 255)
    """
    ...
