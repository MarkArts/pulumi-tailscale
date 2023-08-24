# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

import types

__config__ = pulumi.Config('tailscale')


class _ExportableConfig(types.ModuleType):
    @property
    def api_key(self) -> Optional[str]:
        """
        The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
        Conflicts with 'oauth_client_id' and 'oauth_client_secret'.
        """
        return __config__.get('apiKey')

    @property
    def base_url(self) -> Optional[str]:
        """
        The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
        environment variable.
        """
        return __config__.get('baseUrl')

    @property
    def oauth_client_id(self) -> Optional[str]:
        """
        The OAuth application's ID when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_ID environment
        variable. Both 'oauth_client_id' and 'oauth_client_secret' must be set. Conflicts with 'api_key'.
        """
        return __config__.get('oauthClientId')

    @property
    def oauth_client_secret(self) -> Optional[str]:
        """
        The OAuth application's secret when using OAuth client credentials. Can be set via the TAILSCALE_OAUTH_CLIENT_SECRET
        environment variable. Both 'oauth_client_id' and 'oauth_client_secret' must be set. Conflicts with 'api_key'.
        """
        return __config__.get('oauthClientSecret')

    @property
    def scopes(self) -> Optional[str]:
        """
        The OAuth 2.0 scopes to request when for the access token generated using the supplied OAuth client credentials. See
        https://tailscale.com/kb/1215/oauth-clients/#scopes for available scopes. Only valid when both 'oauth_client_id' and
        'oauth_client_secret' are set.
        """
        return __config__.get('scopes')

    @property
    def tailnet(self) -> Optional[str]:
        """
        The organization name of the Tailnet in which to perform actions. Can be set via the TAILSCALE_TAILNET environment
        variable. Default is the tailnet that owns API credentials passed to the provider.
        """
        return __config__.get('tailnet')

    @property
    def user_agent(self) -> Optional[str]:
        """
        User-Agent header for API requests.
        """
        return __config__.get('userAgent')

