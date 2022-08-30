// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The 4via6 data source is calculates an IPv6 prefix for a given site ID and IPv4 CIDR. See Tailscale documentation for [4via6 subnets](https://tailscale.com/kb/1201/4via6-subnets/) for more details.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const example = pulumi.output(tailscale.get4Via6({
 *     cidr: "10.1.1.0/24",
 *     site: 7,
 * }));
 * ```
 */
export function get4Via6(args: Get4Via6Args, opts?: pulumi.InvokeOptions): Promise<Get4Via6Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("tailscale:index/get4Via6:get4Via6", {
        "cidr": args.cidr,
        "site": args.site,
    }, opts);
}

/**
 * A collection of arguments for invoking get4Via6.
 */
export interface Get4Via6Args {
    /**
     * The IPv4 CIDR to map
     */
    cidr: string;
    /**
     * Site ID (between 0 and 255)
     */
    site: number;
}

/**
 * A collection of values returned by get4Via6.
 */
export interface Get4Via6Result {
    /**
     * The IPv4 CIDR to map
     */
    readonly cidr: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The 4via6 mapped address
     */
    readonly ipv6: string;
    /**
     * Site ID (between 0 and 255)
     */
    readonly site: number;
}

export function get4Via6Output(args: Get4Via6OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<Get4Via6Result> {
    return pulumi.output(args).apply(a => get4Via6(a, opts))
}

/**
 * A collection of arguments for invoking get4Via6.
 */
export interface Get4Via6OutputArgs {
    /**
     * The IPv4 CIDR to map
     */
    cidr: pulumi.Input<string>;
    /**
     * Site ID (between 0 and 255)
     */
    site: pulumi.Input<number>;
}
