// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

/**
 * The devices data source describes a list of devices in a tailnet
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleDevices = pulumi.output(tailscale.getDevices({
 *     namePrefix: "example-",
 * }));
 * ```
 */
export function getDevices(args?: GetDevicesArgs, opts?: pulumi.InvokeOptions): Promise<GetDevicesResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tailscale:index/getDevices:getDevices", {
        "namePrefix": args.namePrefix,
    }, opts);
}

/**
 * A collection of arguments for invoking getDevices.
 */
export interface GetDevicesArgs {
    namePrefix?: string;
}

/**
 * A collection of values returned by getDevices.
 */
export interface GetDevicesResult {
    readonly devices: outputs.GetDevicesDevice[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly namePrefix?: string;
}

export function getDevicesOutput(args?: GetDevicesOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDevicesResult> {
    return pulumi.output(args).apply(a => getDevices(a, opts))
}

/**
 * A collection of arguments for invoking getDevices.
 */
export interface GetDevicesOutputArgs {
    namePrefix?: pulumi.Input<string>;
}
