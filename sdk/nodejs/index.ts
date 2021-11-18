// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./acl";
export * from "./deviceAuthorization";
export * from "./deviceSubnetRoutes";
export * from "./dnsNameservers";
export * from "./dnsPreferences";
export * from "./dnsSearchPaths";
export * from "./getDevice";
export * from "./getDevices";
export * from "./provider";

// Export sub-modules:
import * as config from "./config";
import * as types from "./types";

export {
    config,
    types,
};

// Import resources to register:
import { Acl } from "./acl";
import { DeviceAuthorization } from "./deviceAuthorization";
import { DeviceSubnetRoutes } from "./deviceSubnetRoutes";
import { DnsNameservers } from "./dnsNameservers";
import { DnsPreferences } from "./dnsPreferences";
import { DnsSearchPaths } from "./dnsSearchPaths";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "tailscale:index/acl:Acl":
                return new Acl(name, <any>undefined, { urn })
            case "tailscale:index/deviceAuthorization:DeviceAuthorization":
                return new DeviceAuthorization(name, <any>undefined, { urn })
            case "tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes":
                return new DeviceSubnetRoutes(name, <any>undefined, { urn })
            case "tailscale:index/dnsNameservers:DnsNameservers":
                return new DnsNameservers(name, <any>undefined, { urn })
            case "tailscale:index/dnsPreferences:DnsPreferences":
                return new DnsPreferences(name, <any>undefined, { urn })
            case "tailscale:index/dnsSearchPaths:DnsSearchPaths":
                return new DnsSearchPaths(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("tailscale", "index/acl", _module)
pulumi.runtime.registerResourceModule("tailscale", "index/deviceAuthorization", _module)
pulumi.runtime.registerResourceModule("tailscale", "index/deviceSubnetRoutes", _module)
pulumi.runtime.registerResourceModule("tailscale", "index/dnsNameservers", _module)
pulumi.runtime.registerResourceModule("tailscale", "index/dnsPreferences", _module)
pulumi.runtime.registerResourceModule("tailscale", "index/dnsSearchPaths", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("tailscale", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:tailscale") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
