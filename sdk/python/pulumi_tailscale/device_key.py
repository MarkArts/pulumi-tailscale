# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['DeviceKeyArgs', 'DeviceKey']

@pulumi.input_type
class DeviceKeyArgs:
    def __init__(__self__, *,
                 device_id: pulumi.Input[str],
                 key_expiry_disabled: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a DeviceKey resource.
        :param pulumi.Input[str] device_id: The device to update the key properties of
        :param pulumi.Input[bool] key_expiry_disabled: Determines whether or not the device's key will expire
        """
        DeviceKeyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            device_id=device_id,
            key_expiry_disabled=key_expiry_disabled,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             device_id: Optional[pulumi.Input[str]] = None,
             key_expiry_disabled: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if device_id is None and 'deviceId' in kwargs:
            device_id = kwargs['deviceId']
        if device_id is None:
            raise TypeError("Missing 'device_id' argument")
        if key_expiry_disabled is None and 'keyExpiryDisabled' in kwargs:
            key_expiry_disabled = kwargs['keyExpiryDisabled']

        _setter("device_id", device_id)
        if key_expiry_disabled is not None:
            _setter("key_expiry_disabled", key_expiry_disabled)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Input[str]:
        """
        The device to update the key properties of
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether or not the device's key will expire
        """
        return pulumi.get(self, "key_expiry_disabled")

    @key_expiry_disabled.setter
    def key_expiry_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "key_expiry_disabled", value)


@pulumi.input_type
class _DeviceKeyState:
    def __init__(__self__, *,
                 device_id: Optional[pulumi.Input[str]] = None,
                 key_expiry_disabled: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering DeviceKey resources.
        :param pulumi.Input[str] device_id: The device to update the key properties of
        :param pulumi.Input[bool] key_expiry_disabled: Determines whether or not the device's key will expire
        """
        _DeviceKeyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            device_id=device_id,
            key_expiry_disabled=key_expiry_disabled,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             device_id: Optional[pulumi.Input[str]] = None,
             key_expiry_disabled: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if device_id is None and 'deviceId' in kwargs:
            device_id = kwargs['deviceId']
        if key_expiry_disabled is None and 'keyExpiryDisabled' in kwargs:
            key_expiry_disabled = kwargs['keyExpiryDisabled']

        if device_id is not None:
            _setter("device_id", device_id)
        if key_expiry_disabled is not None:
            _setter("key_expiry_disabled", key_expiry_disabled)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> Optional[pulumi.Input[str]]:
        """
        The device to update the key properties of
        """
        return pulumi.get(self, "device_id")

    @device_id.setter
    def device_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "device_id", value)

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether or not the device's key will expire
        """
        return pulumi.get(self, "key_expiry_disabled")

    @key_expiry_disabled.setter
    def key_expiry_disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "key_expiry_disabled", value)


class DeviceKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 key_expiry_disabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        The device_key resource allows you to update the properties of a device's key

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        example_device = tailscale.get_device(name="device.example.com")
        example_key = tailscale.DeviceKey("exampleKey",
            device_id=example_device.id,
            key_expiry_disabled=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The device to update the key properties of
        :param pulumi.Input[bool] key_expiry_disabled: Determines whether or not the device's key will expire
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DeviceKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The device_key resource allows you to update the properties of a device's key

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        example_device = tailscale.get_device(name="device.example.com")
        example_key = tailscale.DeviceKey("exampleKey",
            device_id=example_device.id,
            key_expiry_disabled=True)
        ```

        :param str resource_name: The name of the resource.
        :param DeviceKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DeviceKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            DeviceKeyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 device_id: Optional[pulumi.Input[str]] = None,
                 key_expiry_disabled: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DeviceKeyArgs.__new__(DeviceKeyArgs)

            if device_id is None and not opts.urn:
                raise TypeError("Missing required property 'device_id'")
            __props__.__dict__["device_id"] = device_id
            __props__.__dict__["key_expiry_disabled"] = key_expiry_disabled
        super(DeviceKey, __self__).__init__(
            'tailscale:index/deviceKey:DeviceKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            device_id: Optional[pulumi.Input[str]] = None,
            key_expiry_disabled: Optional[pulumi.Input[bool]] = None) -> 'DeviceKey':
        """
        Get an existing DeviceKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] device_id: The device to update the key properties of
        :param pulumi.Input[bool] key_expiry_disabled: Determines whether or not the device's key will expire
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DeviceKeyState.__new__(_DeviceKeyState)

        __props__.__dict__["device_id"] = device_id
        __props__.__dict__["key_expiry_disabled"] = key_expiry_disabled
        return DeviceKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deviceId")
    def device_id(self) -> pulumi.Output[str]:
        """
        The device to update the key properties of
        """
        return pulumi.get(self, "device_id")

    @property
    @pulumi.getter(name="keyExpiryDisabled")
    def key_expiry_disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether or not the device's key will expire
        """
        return pulumi.get(self, "key_expiry_disabled")

