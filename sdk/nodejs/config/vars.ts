// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("tailscale");

export declare const apiKey: string | undefined;
Object.defineProperty(exports, "apiKey", {
    get() {
        return __config.get("apiKey");
    },
    enumerable: true,
});

export declare const baseUrl: string | undefined;
Object.defineProperty(exports, "baseUrl", {
    get() {
        return __config.get("baseUrl");
    },
    enumerable: true,
});

export declare const tailnet: string | undefined;
Object.defineProperty(exports, "tailnet", {
    get() {
        return __config.get("tailnet");
    },
    enumerable: true,
});

