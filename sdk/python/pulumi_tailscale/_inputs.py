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
    'ContactsAccountArgs',
    'ContactsSecurityArgs',
    'ContactsSupportArgs',
]

@pulumi.input_type
class ContactsAccountArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str]):
        """
        :param pulumi.Input[str] email: Email address to send communications to
        """
        pulumi.set(__self__, "email", email)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email address to send communications to
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)


@pulumi.input_type
class ContactsSecurityArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str]):
        """
        :param pulumi.Input[str] email: Email address to send communications to
        """
        pulumi.set(__self__, "email", email)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email address to send communications to
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)


@pulumi.input_type
class ContactsSupportArgs:
    def __init__(__self__, *,
                 email: pulumi.Input[str]):
        """
        :param pulumi.Input[str] email: Email address to send communications to
        """
        pulumi.set(__self__, "email", email)

    @property
    @pulumi.getter
    def email(self) -> pulumi.Input[str]:
        """
        Email address to send communications to
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: pulumi.Input[str]):
        pulumi.set(self, "email", value)


