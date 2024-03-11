# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DeviceSubnetRoutesArgs', 'DeviceSubnetRoutes']

@pulumi.input_type
class DeviceSubnetRoutesArgs:
    def __init__(__self__, *,
                 device_id: pulumi.Input[str],
                 routes: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a DeviceSubnetRoutes resource.
        :param pulumi.Input[str] device_id: The device to set subnet routes for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] routes: The subnet routes that are enabled to be routed by a device
        """
        pulumi.set(__self__, "device_id", device_id)
        pulumi.set(__self__, "routes", routes)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        The device to set subnet routes for
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The subnet routes that are enabled to be routed by a device
        """
        return pulumi.get(self, "routes")

    @routes.setter
    def routes(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "routes", value)


@pulumi.input_type
class _DeviceSubnetRoutesState:
    def __init__(__self__, *,
                 device_id: Optional[pulumi.Input[str]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DeviceSubnetRoutes resources.
        :param pulumi.Input[str] device_id: The device to set subnet routes for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] routes: The subnet routes that are enabled to be routed by a device
        """
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if routes is not None:
            pulumi.set(__self__, "routes", routes)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        The device to set subnet routes for
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def routes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The subnet routes that are enabled to be routed by a device
        """
        return pulumi.get(self, "routes")

    @routes.setter
    def routes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "routes", value)


class DeviceSubnetRoutes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The device_subnet_routes resource allows you to configure subnet routes for your Tailscale devices. See https://tailscale.com/kb/1019/subnets for more information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_device = tailscale.get_device(name="device.example.com")
        sample_routes = tailscale.DeviceSubnetRoutes("sampleRoutes",
            device_id=sample_device.id,
            routes=[
                "10.0.1.0/24",
                "1.2.0.0/16",
                "2.0.0.0/24",
            ])
        sample_exit_node = tailscale.DeviceSubnetRoutes("sampleExitNode",
            device_id=sample_device.id,
            routes=[
                "0.0.0.0/0",
                "::/0",
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The device to set subnet routes for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] routes: The subnet routes that are enabled to be routed by a device
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceSubnetRoutesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The device_subnet_routes resource allows you to configure subnet routes for your Tailscale devices. See https://tailscale.com/kb/1019/subnets for more information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_device = tailscale.get_device(name="device.example.com")
        sample_routes = tailscale.DeviceSubnetRoutes("sampleRoutes",
            device_id=sample_device.id,
            routes=[
                "10.0.1.0/24",
                "1.2.0.0/16",
                "2.0.0.0/24",
            ])
        sample_exit_node = tailscale.DeviceSubnetRoutes("sampleExitNode",
            device_id=sample_device.id,
            routes=[
                "0.0.0.0/0",
                "::/0",
            ])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param DeviceSubnetRoutesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceSubnetRoutesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 routes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceSubnetRoutesArgs.__new__(DeviceSubnetRoutesArgs)

            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            if routes is None and not opts.urn:
                raise TypeError("Missing required property 'routes'")
            __props__.__dict__["routes"] = routes
        super(DeviceSubnetRoutes, __self__).__init__(
            'tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_id: Optional[pulumi.Input[str]] = None,
            routes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'DeviceSubnetRoutes':
        """
        Get an existing DeviceSubnetRoutes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The device to set subnet routes for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] routes: The subnet routes that are enabled to be routed by a device
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceSubnetRoutesState.__new__(_DeviceSubnetRoutesState)

        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["routes"] = routes
        return DeviceSubnetRoutes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        The device to set subnet routes for
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter
    def routes(self) -> pulumi.Output[Sequence[str]]:
        """
        The subnet routes that are enabled to be routed by a device
        """
        return pulumi.get(self, "routes")

