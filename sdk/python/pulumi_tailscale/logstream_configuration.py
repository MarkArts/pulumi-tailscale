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

__all__ = ['LogstreamConfigurationArgs', 'LogstreamConfiguration']

@pulumi.input_type
class LogstreamConfigurationArgs:
    def __init__(__self__, *,
                 destination_type: pulumi.Input[str],
                 log_type: pulumi.Input[str],
                 token: pulumi.Input[str],
                 url: pulumi.Input[str],
                 user: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a LogstreamConfiguration resource.
        :param pulumi.Input[str] destination_type: The type of system to which logs are being streamed.
        :param pulumi.Input[str] log_type: The type of log that is streamed to this endpoint.
        :param pulumi.Input[str] token: The token/password with which log streams to this endpoint should be authenticated.
        :param pulumi.Input[str] url: The URL to which log streams are being posted.
        :param pulumi.Input[str] user: The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        pulumi.set(__self__, "destination_type", destination_type)
        pulumi.set(__self__, "log_type", log_type)
        pulumi.set(__self__, "token", token)
        pulumi.set(__self__, "url", url)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> pulumi.Input[str]:
        """
        The type of system to which logs are being streamed.
        """
        return pulumi.get(self, "destination_type")

    @destination_type.setter
    def destination_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "destination_type", value)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Input[str]:
        """
        The type of log that is streamed to this endpoint.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "log_type", value)

    @property
    @pulumi.getter
    def token(self) -> pulumi.Input[str]:
        """
        The token/password with which log streams to this endpoint should be authenticated.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: pulumi.Input[str]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The URL to which log streams are being posted.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


@pulumi.input_type
class _LogstreamConfigurationState:
    def __init__(__self__, *,
                 destination_type: Optional[pulumi.Input[str]] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering LogstreamConfiguration resources.
        :param pulumi.Input[str] destination_type: The type of system to which logs are being streamed.
        :param pulumi.Input[str] log_type: The type of log that is streamed to this endpoint.
        :param pulumi.Input[str] token: The token/password with which log streams to this endpoint should be authenticated.
        :param pulumi.Input[str] url: The URL to which log streams are being posted.
        :param pulumi.Input[str] user: The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        if destination_type is not None:
            pulumi.set(__self__, "destination_type", destination_type)
        if log_type is not None:
            pulumi.set(__self__, "log_type", log_type)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if url is not None:
            pulumi.set(__self__, "url", url)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of system to which logs are being streamed.
        """
        return pulumi.get(self, "destination_type")

    @destination_type.setter
    def destination_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "destination_type", value)

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of log that is streamed to this endpoint.
        """
        return pulumi.get(self, "log_type")

    @log_type.setter
    def log_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "log_type", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The token/password with which log streams to this endpoint should be authenticated.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter
    def url(self) -> Optional[pulumi.Input[str]]:
        """
        The URL to which log streams are being posted.
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class LogstreamConfiguration(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_type: Optional[pulumi.Input[str]] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The logstream_configuration resource allows you to configure streaming configuration or network flow logs to a supported security information and event management (SIEM) system. See https://tailscale.com/kb/1255/log-streaming for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_logstream_configuration = tailscale.LogstreamConfiguration("sample_logstream_configuration",
            log_type="configuration",
            destination_type="panther",
            url="https://example.com",
            token="some-token")
        ```

        ## Import

        Logstream configuration can be imported using the logstream configuration id, e.g.,

        ```sh
        $ pulumi import tailscale:index/logstreamConfiguration:LogstreamConfiguration sample_logstream_configuration 123456789
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_type: The type of system to which logs are being streamed.
        :param pulumi.Input[str] log_type: The type of log that is streamed to this endpoint.
        :param pulumi.Input[str] token: The token/password with which log streams to this endpoint should be authenticated.
        :param pulumi.Input[str] url: The URL to which log streams are being posted.
        :param pulumi.Input[str] user: The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: LogstreamConfigurationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The logstream_configuration resource allows you to configure streaming configuration or network flow logs to a supported security information and event management (SIEM) system. See https://tailscale.com/kb/1255/log-streaming for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_tailscale as tailscale

        sample_logstream_configuration = tailscale.LogstreamConfiguration("sample_logstream_configuration",
            log_type="configuration",
            destination_type="panther",
            url="https://example.com",
            token="some-token")
        ```

        ## Import

        Logstream configuration can be imported using the logstream configuration id, e.g.,

        ```sh
        $ pulumi import tailscale:index/logstreamConfiguration:LogstreamConfiguration sample_logstream_configuration 123456789
        ```

        :param str resource_name: The name of the resource.
        :param LogstreamConfigurationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(LogstreamConfigurationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 destination_type: Optional[pulumi.Input[str]] = None,
                 log_type: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = LogstreamConfigurationArgs.__new__(LogstreamConfigurationArgs)

            if destination_type is None and not opts.urn:
                raise TypeError("Missing required property 'destination_type'")
            __props__.__dict__["destination_type"] = destination_type
            if log_type is None and not opts.urn:
                raise TypeError("Missing required property 'log_type'")
            __props__.__dict__["log_type"] = log_type
            if token is None and not opts.urn:
                raise TypeError("Missing required property 'token'")
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["user"] = user
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(LogstreamConfiguration, __self__).__init__(
            'tailscale:index/logstreamConfiguration:LogstreamConfiguration',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            destination_type: Optional[pulumi.Input[str]] = None,
            log_type: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            url: Optional[pulumi.Input[str]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'LogstreamConfiguration':
        """
        Get an existing LogstreamConfiguration resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] destination_type: The type of system to which logs are being streamed.
        :param pulumi.Input[str] log_type: The type of log that is streamed to this endpoint.
        :param pulumi.Input[str] token: The token/password with which log streams to this endpoint should be authenticated.
        :param pulumi.Input[str] url: The URL to which log streams are being posted.
        :param pulumi.Input[str] user: The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _LogstreamConfigurationState.__new__(_LogstreamConfigurationState)

        __props__.__dict__["destination_type"] = destination_type
        __props__.__dict__["log_type"] = log_type
        __props__.__dict__["token"] = token
        __props__.__dict__["url"] = url
        __props__.__dict__["user"] = user
        return LogstreamConfiguration(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="destinationType")
    def destination_type(self) -> pulumi.Output[str]:
        """
        The type of system to which logs are being streamed.
        """
        return pulumi.get(self, "destination_type")

    @property
    @pulumi.getter(name="logType")
    def log_type(self) -> pulumi.Output[str]:
        """
        The type of log that is streamed to this endpoint.
        """
        return pulumi.get(self, "log_type")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The token/password with which log streams to this endpoint should be authenticated.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL to which log streams are being posted.
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[Optional[str]]:
        """
        The username with which log streams to this endpoint are authenticated. Only required if destination_type is 'elastic', defaults to 'user' if not set.
        """
        return pulumi.get(self, "user")

