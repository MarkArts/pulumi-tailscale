// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The deviceAuthorization resource is used to approve new devices before they can join the tailnet.
 * See the [Tailscale device authorization documentation](https://tailscale.com/kb/1099/device-authorization) for more
 * information.
 *
 * The Tailscale API currently only supports authorizing devices, but not rejecting/removing them. Once a device is
 * authorized by this provider it cannot be modified again afterwards. Modifying or deleting the resource
 * will not affect the device's authorization within the tailnet.
 */
export class DeviceAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing DeviceAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceAuthorizationState, opts?: pulumi.CustomResourceOptions): DeviceAuthorization {
        return new DeviceAuthorization(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/deviceAuthorization:DeviceAuthorization';

    /**
     * Returns true if the given object is an instance of DeviceAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceAuthorization.__pulumiType;
    }

    /**
     * Indicates if the device is allowed to join the tailnet.
     */
    public readonly authorized!: pulumi.Output<boolean>;
    /**
     * The device to authorize.
     */
    public readonly deviceId!: pulumi.Output<string>;

    /**
     * Create a DeviceAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceAuthorizationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceAuthorizationArgs | DeviceAuthorizationState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceAuthorizationState | undefined;
            inputs["authorized"] = state ? state.authorized : undefined;
            inputs["deviceId"] = state ? state.deviceId : undefined;
        } else {
            const args = argsOrState as DeviceAuthorizationArgs | undefined;
            if ((!args || args.authorized === undefined) && !opts.urn) {
                throw new Error("Missing required property 'authorized'");
            }
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            inputs["authorized"] = args ? args.authorized : undefined;
            inputs["deviceId"] = args ? args.deviceId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(DeviceAuthorization.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceAuthorization resources.
 */
export interface DeviceAuthorizationState {
    /**
     * Indicates if the device is allowed to join the tailnet.
     */
    authorized?: pulumi.Input<boolean>;
    /**
     * The device to authorize.
     */
    deviceId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DeviceAuthorization resource.
 */
export interface DeviceAuthorizationArgs {
    /**
     * Indicates if the device is allowed to join the tailnet.
     */
    authorized: pulumi.Input<boolean>;
    /**
     * The device to authorize.
     */
    deviceId: pulumi.Input<string>;
}