// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The device data source describes a single device in a tailnet
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleDevice = tailscale.getDevice({
 *     name: "device1.example.ts.net",
 *     waitFor: "60s",
 * });
 * const sampleDevice2 = tailscale.getDevice({
 *     hostname: "device2",
 *     waitFor: "60s",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDevice(args?: GetDeviceArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("tailscale:index/getDevice:getDevice", {
        "hostname": args.hostname,
        "name": args.name,
        "waitFor": args.waitFor,
    }, opts);
}

/**
 * A collection of arguments for invoking getDevice.
 */
export interface GetDeviceArgs {
    /**
     * The short hostname of the device
     */
    hostname?: string;
    /**
     * The full name of the device (e.g. `hostname.domain.ts.net`)
     */
    name?: string;
    /**
     * If specified, the provider will make multiple attempts to obtain the data source until the waitFor duration is reached. Retries are made every second so this value should be greater than 1s
     */
    waitFor?: string;
}

/**
 * A collection of values returned by getDevice.
 */
export interface GetDeviceResult {
    /**
     * The list of device's IPs
     */
    readonly addresses: string[];
    /**
     * The short hostname of the device
     */
    readonly hostname?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The full name of the device (e.g. `hostname.domain.ts.net`)
     */
    readonly name?: string;
    /**
     * The tags applied to the device
     */
    readonly tags: string[];
    /**
     * The user associated with the device
     */
    readonly user: string;
    /**
     * If specified, the provider will make multiple attempts to obtain the data source until the waitFor duration is reached. Retries are made every second so this value should be greater than 1s
     */
    readonly waitFor?: string;
}
/**
 * The device data source describes a single device in a tailnet
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleDevice = tailscale.getDevice({
 *     name: "device1.example.ts.net",
 *     waitFor: "60s",
 * });
 * const sampleDevice2 = tailscale.getDevice({
 *     hostname: "device2",
 *     waitFor: "60s",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getDeviceOutput(args?: GetDeviceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceResult> {
    return pulumi.output(args).apply((a: any) => getDevice(a, opts))
}

/**
 * A collection of arguments for invoking getDevice.
 */
export interface GetDeviceOutputArgs {
    /**
     * The short hostname of the device
     */
    hostname?: pulumi.Input<string>;
    /**
     * The full name of the device (e.g. `hostname.domain.ts.net`)
     */
    name?: pulumi.Input<string>;
    /**
     * If specified, the provider will make multiple attempts to obtain the data source until the waitFor duration is reached. Retries are made every second so this value should be greater than 1s
     */
    waitFor?: pulumi.Input<string>;
}
