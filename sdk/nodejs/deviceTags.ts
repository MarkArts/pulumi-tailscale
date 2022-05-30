// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The deviceTags resource is used to apply tags to Tailscale devices. See https://tailscale.com/kb/1068/acl-tags/ for more details.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleDevice = tailscale.getDevice({
 *     name: "device.example.com",
 * });
 * const sampleTags = new tailscale.DeviceTags("sampleTags", {
 *     deviceId: sampleDevice.then(sampleDevice => sampleDevice.id),
 *     tags: ["room:bedroom"],
 * });
 * ```
 */
export class DeviceTags extends pulumi.CustomResource {
    /**
     * Get an existing DeviceTags resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceTagsState, opts?: pulumi.CustomResourceOptions): DeviceTags {
        return new DeviceTags(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/deviceTags:DeviceTags';

    /**
     * Returns true if the given object is an instance of DeviceTags.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceTags {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceTags.__pulumiType;
    }

    /**
     * The device to set tags for
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * The tags to apply to the device
     */
    public readonly tags!: pulumi.Output<string[]>;

    /**
     * Create a DeviceTags resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceTagsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceTagsArgs | DeviceTagsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceTagsState | undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
        } else {
            const args = argsOrState as DeviceTagsArgs | undefined;
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            if ((!args || args.tags === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tags'");
            }
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceTags.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceTags resources.
 */
export interface DeviceTagsState {
    /**
     * The device to set tags for
     */
    deviceId?: pulumi.Input<string>;
    /**
     * The tags to apply to the device
     */
    tags?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DeviceTags resource.
 */
export interface DeviceTagsArgs {
    /**
     * The device to set tags for
     */
    deviceId: pulumi.Input<string>;
    /**
     * The tags to apply to the device
     */
    tags: pulumi.Input<pulumi.Input<string>[]>;
}
