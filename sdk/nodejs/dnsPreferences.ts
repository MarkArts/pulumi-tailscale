// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The dnsPreferences resource allows you to configure DNS preferences for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const samplePreferences = new tailscale.DnsPreferences("sample_preferences", {
 *     magicDns: true,
 * });
 * ```
 */
export class DnsPreferences extends pulumi.CustomResource {
    /**
     * Get an existing DnsPreferences resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DnsPreferencesState, opts?: pulumi.CustomResourceOptions): DnsPreferences {
        return new DnsPreferences(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/dnsPreferences:DnsPreferences';

    /**
     * Returns true if the given object is an instance of DnsPreferences.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsPreferences {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsPreferences.__pulumiType;
    }

    /**
     * Whether or not to enable magic DNS
     */
    public readonly magicDns!: pulumi.Output<boolean>;

    /**
     * Create a DnsPreferences resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsPreferencesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DnsPreferencesArgs | DnsPreferencesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DnsPreferencesState | undefined;
            resourceInputs["magicDns"] = state ? state.magicDns : undefined;
        } else {
            const args = argsOrState as DnsPreferencesArgs | undefined;
            if ((!args || args.magicDns === undefined) && !opts.urn) {
                throw new Error("Missing required property 'magicDns'");
            }
            resourceInputs["magicDns"] = args ? args.magicDns : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsPreferences.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DnsPreferences resources.
 */
export interface DnsPreferencesState {
    /**
     * Whether or not to enable magic DNS
     */
    magicDns?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DnsPreferences resource.
 */
export interface DnsPreferencesArgs {
    /**
     * Whether or not to enable magic DNS
     */
    magicDns: pulumi.Input<boolean>;
}
