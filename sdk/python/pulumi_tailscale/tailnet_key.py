# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['TailnetKeyArgs', 'TailnetKey']

@pulumi.input_type
class TailnetKeyArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[int]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 recreate_if_invalid: Optional[pulumi.Input[str]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a TailnetKey resource.
        :param pulumi.Input[str] description: A description of the key consisting of alphanumeric characters. Defaults to `""`.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral. Defaults to `false`.
        :param pulumi.Input[int] expiry: The expiry of the key in seconds. Defaults to `7776000` (90 days).
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        :param pulumi.Input[str] recreate_if_invalid: Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ephemeral is not None:
            pulumi.set(__self__, "ephemeral", ephemeral)
        if expiry is not None:
            pulumi.set(__self__, "expiry", expiry)
        if preauthorized is not None:
            pulumi.set(__self__, "preauthorized", preauthorized)
        if recreate_if_invalid is not None:
            pulumi.set(__self__, "recreate_if_invalid", recreate_if_invalid)
        if reusable is not None:
            pulumi.set(__self__, "reusable", reusable)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the key consisting of alphanumeric characters. Defaults to `""`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is ephemeral. Defaults to `false`.
        """
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter
    def expiry(self) -> Optional[pulumi.Input[int]]:
        """
        The expiry of the key in seconds. Defaults to `7776000` (90 days).
        """
        return pulumi.get(self, "expiry")

    @expiry.setter
    def expiry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiry", value)

    @property
    @pulumi.getter
    def preauthorized(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        """
        return pulumi.get(self, "preauthorized")

    @preauthorized.setter
    def preauthorized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preauthorized", value)

    @property
    @pulumi.getter(name="recreateIfInvalid")
    def recreate_if_invalid(self) -> Optional[pulumi.Input[str]]:
        """
        Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        """
        return pulumi.get(self, "recreate_if_invalid")

    @recreate_if_invalid.setter
    def recreate_if_invalid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recreate_if_invalid", value)

    @property
    @pulumi.getter
    def reusable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is reusable or single-use. Defaults to `false`.
        """
        return pulumi.get(self, "reusable")

    @reusable.setter
    def reusable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reusable", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags to apply to the machines authenticated by the key.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _TailnetKeyState:
    def __init__(__self__, *,
                 created_at: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 expires_at: Optional[pulumi.Input[str]] = None,
                 expiry: Optional[pulumi.Input[int]] = None,
                 invalid: Optional[pulumi.Input[bool]] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 recreate_if_invalid: Optional[pulumi.Input[str]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering TailnetKey resources.
        :param pulumi.Input[str] created_at: The creation timestamp of the key in RFC3339 format
        :param pulumi.Input[str] description: A description of the key consisting of alphanumeric characters. Defaults to `""`.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral. Defaults to `false`.
        :param pulumi.Input[str] expires_at: The expiry timestamp of the key in RFC3339 format
        :param pulumi.Input[int] expiry: The expiry of the key in seconds. Defaults to `7776000` (90 days).
        :param pulumi.Input[bool] invalid: Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
        :param pulumi.Input[str] key: The authentication key
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        :param pulumi.Input[str] recreate_if_invalid: Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        if created_at is not None:
            pulumi.set(__self__, "created_at", created_at)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ephemeral is not None:
            pulumi.set(__self__, "ephemeral", ephemeral)
        if expires_at is not None:
            pulumi.set(__self__, "expires_at", expires_at)
        if expiry is not None:
            pulumi.set(__self__, "expiry", expiry)
        if invalid is not None:
            pulumi.set(__self__, "invalid", invalid)
        if key is not None:
            pulumi.set(__self__, "key", key)
        if preauthorized is not None:
            pulumi.set(__self__, "preauthorized", preauthorized)
        if recreate_if_invalid is not None:
            pulumi.set(__self__, "recreate_if_invalid", recreate_if_invalid)
        if reusable is not None:
            pulumi.set(__self__, "reusable", reusable)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> Optional[pulumi.Input[str]]:
        """
        The creation timestamp of the key in RFC3339 format
        """
        return pulumi.get(self, "created_at")

    @created_at.setter
    def created_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "created_at", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the key consisting of alphanumeric characters. Defaults to `""`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ephemeral(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is ephemeral. Defaults to `false`.
        """
        return pulumi.get(self, "ephemeral")

    @ephemeral.setter
    def ephemeral(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ephemeral", value)

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> Optional[pulumi.Input[str]]:
        """
        The expiry timestamp of the key in RFC3339 format
        """
        return pulumi.get(self, "expires_at")

    @expires_at.setter
    def expires_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expires_at", value)

    @property
    @pulumi.getter
    def expiry(self) -> Optional[pulumi.Input[int]]:
        """
        The expiry of the key in seconds. Defaults to `7776000` (90 days).
        """
        return pulumi.get(self, "expiry")

    @expiry.setter
    def expiry(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiry", value)

    @property
    @pulumi.getter
    def invalid(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
        """
        return pulumi.get(self, "invalid")

    @invalid.setter
    def invalid(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "invalid", value)

    @property
    @pulumi.getter
    def key(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication key
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def preauthorized(self) -> Optional[pulumi.Input[bool]]:
        """
        Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        """
        return pulumi.get(self, "preauthorized")

    @preauthorized.setter
    def preauthorized(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "preauthorized", value)

    @property
    @pulumi.getter(name="recreateIfInvalid")
    def recreate_if_invalid(self) -> Optional[pulumi.Input[str]]:
        """
        Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        """
        return pulumi.get(self, "recreate_if_invalid")

    @recreate_if_invalid.setter
    def recreate_if_invalid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "recreate_if_invalid", value)

    @property
    @pulumi.getter
    def reusable(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the key is reusable or single-use. Defaults to `false`.
        """
        return pulumi.get(self, "reusable")

    @reusable.setter
    def reusable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "reusable", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of tags to apply to the machines authenticated by the key.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


class TailnetKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[int]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 recreate_if_invalid: Optional[pulumi.Input[str]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        The tailnet_key resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_key = tailscale.TailnetKey("sample_key",
            reusable=True,
            ephemeral=False,
            preauthorized=True,
            expiry=3600,
            description="Sample key")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the key consisting of alphanumeric characters. Defaults to `""`.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral. Defaults to `false`.
        :param pulumi.Input[int] expiry: The expiry of the key in seconds. Defaults to `7776000` (90 days).
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        :param pulumi.Input[str] recreate_if_invalid: Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[TailnetKeyArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The tailnet_key resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_key = tailscale.TailnetKey("sample_key",
            reusable=True,
            ephemeral=False,
            preauthorized=True,
            expiry=3600,
            description="Sample key")
        ```

        :param str resource_name: The name of the resource.
        :param TailnetKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TailnetKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ephemeral: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[int]] = None,
                 preauthorized: Optional[pulumi.Input[bool]] = None,
                 recreate_if_invalid: Optional[pulumi.Input[str]] = None,
                 reusable: Optional[pulumi.Input[bool]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TailnetKeyArgs.__new__(TailnetKeyArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ephemeral"] = ephemeral
            __props__.__dict__["expiry"] = expiry
            __props__.__dict__["preauthorized"] = preauthorized
            __props__.__dict__["recreate_if_invalid"] = recreate_if_invalid
            __props__.__dict__["reusable"] = reusable
            __props__.__dict__["tags"] = tags
            __props__.__dict__["created_at"] = None
            __props__.__dict__["expires_at"] = None
            __props__.__dict__["invalid"] = None
            __props__.__dict__["key"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["key"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(TailnetKey, __self__).__init__(
            'tailscale:index/tailnetKey:TailnetKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            created_at: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ephemeral: Optional[pulumi.Input[bool]] = None,
            expires_at: Optional[pulumi.Input[str]] = None,
            expiry: Optional[pulumi.Input[int]] = None,
            invalid: Optional[pulumi.Input[bool]] = None,
            key: Optional[pulumi.Input[str]] = None,
            preauthorized: Optional[pulumi.Input[bool]] = None,
            recreate_if_invalid: Optional[pulumi.Input[str]] = None,
            reusable: Optional[pulumi.Input[bool]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'TailnetKey':
        """
        Get an existing TailnetKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] created_at: The creation timestamp of the key in RFC3339 format
        :param pulumi.Input[str] description: A description of the key consisting of alphanumeric characters. Defaults to `""`.
        :param pulumi.Input[bool] ephemeral: Indicates if the key is ephemeral. Defaults to `false`.
        :param pulumi.Input[str] expires_at: The expiry timestamp of the key in RFC3339 format
        :param pulumi.Input[int] expiry: The expiry of the key in seconds. Defaults to `7776000` (90 days).
        :param pulumi.Input[bool] invalid: Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
        :param pulumi.Input[str] key: The authentication key
        :param pulumi.Input[bool] preauthorized: Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        :param pulumi.Input[str] recreate_if_invalid: Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        :param pulumi.Input[bool] reusable: Indicates if the key is reusable or single-use. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: List of tags to apply to the machines authenticated by the key.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TailnetKeyState.__new__(_TailnetKeyState)

        __props__.__dict__["created_at"] = created_at
        __props__.__dict__["description"] = description
        __props__.__dict__["ephemeral"] = ephemeral
        __props__.__dict__["expires_at"] = expires_at
        __props__.__dict__["expiry"] = expiry
        __props__.__dict__["invalid"] = invalid
        __props__.__dict__["key"] = key
        __props__.__dict__["preauthorized"] = preauthorized
        __props__.__dict__["recreate_if_invalid"] = recreate_if_invalid
        __props__.__dict__["reusable"] = reusable
        __props__.__dict__["tags"] = tags
        return TailnetKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createdAt")
    def created_at(self) -> pulumi.Output[str]:
        """
        The creation timestamp of the key in RFC3339 format
        """
        return pulumi.get(self, "created_at")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the key consisting of alphanumeric characters. Defaults to `""`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ephemeral(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the key is ephemeral. Defaults to `false`.
        """
        return pulumi.get(self, "ephemeral")

    @property
    @pulumi.getter(name="expiresAt")
    def expires_at(self) -> pulumi.Output[str]:
        """
        The expiry timestamp of the key in RFC3339 format
        """
        return pulumi.get(self, "expires_at")

    @property
    @pulumi.getter
    def expiry(self) -> pulumi.Output[Optional[int]]:
        """
        The expiry of the key in seconds. Defaults to `7776000` (90 days).
        """
        return pulumi.get(self, "expiry")

    @property
    @pulumi.getter
    def invalid(self) -> pulumi.Output[bool]:
        """
        Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
        """
        return pulumi.get(self, "invalid")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        The authentication key
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def preauthorized(self) -> pulumi.Output[Optional[bool]]:
        """
        Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        """
        return pulumi.get(self, "preauthorized")

    @property
    @pulumi.getter(name="recreateIfInvalid")
    def recreate_if_invalid(self) -> pulumi.Output[Optional[str]]:
        """
        Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        """
        return pulumi.get(self, "recreate_if_invalid")

    @property
    @pulumi.getter
    def reusable(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the key is reusable or single-use. Defaults to `false`.
        """
        return pulumi.get(self, "reusable")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of tags to apply to the machines authenticated by the key.
        """
        return pulumi.get(self, "tags")

