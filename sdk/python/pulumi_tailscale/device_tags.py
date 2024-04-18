# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DeviceTagsArgs', 'DeviceTags']

@pulumi.input_type
class DeviceTagsArgs:
    def __init__(__self__, *,
                 device_id: pulumi.Input[str],
                 tags: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        The set of arguments for constructing a DeviceTags resource.
        :param pulumi.Input[str] device_id: The device to set tags for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags to apply to the device
        """
        pulumi.set(__self__, "device_id", device_id)
        pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        The device to set tags for
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The tags to apply to the device
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _DeviceTagsState:
    def __init__(__self__, *,
                 device_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering DeviceTags resources.
        :param pulumi.Input[str] device_id: The device to set tags for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags to apply to the device
        """
        if device_id is not None:
            pulumi.set(__self__, "device_id", device_id)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        The device to set tags for
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The tags to apply to the device
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class DeviceTags(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The device_tags resource is used to apply tags to Tailscale devices. See https://tailscale.com/kb/1068/acl-tags/ for more details.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_device = tailscale.get_device(name="device.example.com")
        sample_tags = tailscale.DeviceTags("sample_tags",
            device_id=sample_device.id,
            tags=["room:bedroom"])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The device to set tags for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags to apply to the device
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceTagsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The device_tags resource is used to apply tags to Tailscale devices. See https://tailscale.com/kb/1068/acl-tags/ for more details.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_device = tailscale.get_device(name="device.example.com")
        sample_tags = tailscale.DeviceTags("sample_tags",
            device_id=sample_device.id,
            tags=["room:bedroom"])
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param DeviceTagsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceTagsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceTagsArgs.__new__(DeviceTagsArgs)

            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            if tags is None and not opts.urn:
                raise TypeError("Missing required property 'tags'")
            __props__.__dict__["tags"] = tags
        super(DeviceTags, __self__).__init__(
            'tailscale:index/deviceTags:DeviceTags',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_id: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'DeviceTags':
        """
        Get an existing DeviceTags resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The device to set tags for
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: The tags to apply to the device
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceTagsState.__new__(_DeviceTagsState)

        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["tags"] = tags
        return DeviceTags(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        The device to set tags for
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Sequence[str]]:
        """
        The tags to apply to the device
        """
        return pulumi.get(self, "tags")

