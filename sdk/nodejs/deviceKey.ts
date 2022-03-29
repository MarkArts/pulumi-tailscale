// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The `deviceKey` resource allows you to modify the property of a device's key.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const exampleDevice = tailscale.getDevice({
 *     name: "device.example.com",
 * });
 * const exampleKey = new tailscale.DeviceKey("exampleKey", {
 *     deviceId: exampleDevice.then(exampleDevice => exampleDevice.id),
 *     preauthorized: true,
 *     keyExpiryDisabled: true,
 * });
 * ```
 */
export class DeviceKey extends pulumi.CustomResource {
    /**
     * Get an existing DeviceKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceKeyState, opts?: pulumi.CustomResourceOptions): DeviceKey {
        return new DeviceKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/deviceKey:DeviceKey';

    /**
     * Returns true if the given object is an instance of DeviceKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceKey.__pulumiType;
    }

    /**
     * The device to change the key properties of.
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * Whether the device's key will ever expire.
     */
    public readonly keyExpiryDisabled!: pulumi.Output<boolean | undefined>;
    /**
     * Whether the device should be authorized for the tailnet by default, works in tailnets 
     * where device authorization is enabled.
     */
    public readonly preauthorized!: pulumi.Output<boolean | undefined>;

    /**
     * Create a DeviceKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceKeyArgs | DeviceKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceKeyState | undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["keyExpiryDisabled"] = state ? state.keyExpiryDisabled : undefined;
            resourceInputs["preauthorized"] = state ? state.preauthorized : undefined;
        } else {
            const args = argsOrState as DeviceKeyArgs | undefined;
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["keyExpiryDisabled"] = args ? args.keyExpiryDisabled : undefined;
            resourceInputs["preauthorized"] = args ? args.preauthorized : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceKey resources.
 */
export interface DeviceKeyState {
    /**
     * The device to change the key properties of.
     */
    deviceId?: pulumi.Input<string>;
    /**
     * Whether the device's key will ever expire.
     */
    keyExpiryDisabled?: pulumi.Input<boolean>;
    /**
     * Whether the device should be authorized for the tailnet by default, works in tailnets 
     * where device authorization is enabled.
     */
    preauthorized?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a DeviceKey resource.
 */
export interface DeviceKeyArgs {
    /**
     * The device to change the key properties of.
     */
    deviceId: pulumi.Input<string>;
    /**
     * Whether the device's key will ever expire.
     */
    keyExpiryDisabled?: pulumi.Input<boolean>;
    /**
     * Whether the device should be authorized for the tailnet by default, works in tailnets 
     * where device authorization is enabled.
     */
    preauthorized?: pulumi.Input<boolean>;
}
