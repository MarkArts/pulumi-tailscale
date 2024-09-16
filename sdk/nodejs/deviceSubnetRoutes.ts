// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleDevice = tailscale.getDevice({
 *     name: "device.example.com",
 * });
 * const sampleRoutes = new tailscale.DeviceSubnetRoutes("sample_routes", {
 *     deviceId: sampleDevice.then(sampleDevice => sampleDevice.id),
 *     routes: [
 *         "10.0.1.0/24",
 *         "1.2.0.0/16",
 *         "2.0.0.0/24",
 *     ],
 * });
 * const sampleExitNode = new tailscale.DeviceSubnetRoutes("sample_exit_node", {
 *     deviceId: sampleDevice.then(sampleDevice => sampleDevice.id),
 *     routes: [
 *         "0.0.0.0/0",
 *         "::/0",
 *     ],
 * });
 * ```
 */
export class DeviceSubnetRoutes extends pulumi.CustomResource {
    /**
     * Get an existing DeviceSubnetRoutes resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DeviceSubnetRoutesState, opts?: pulumi.CustomResourceOptions): DeviceSubnetRoutes {
        return new DeviceSubnetRoutes(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes';

    /**
     * Returns true if the given object is an instance of DeviceSubnetRoutes.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DeviceSubnetRoutes {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DeviceSubnetRoutes.__pulumiType;
    }

    /**
     * The device to set subnet routes for
     */
    public readonly deviceId!: pulumi.Output<string>;
    /**
     * The subnet routes that are enabled to be routed by a device
     */
    public readonly routes!: pulumi.Output<string[]>;

    /**
     * Create a DeviceSubnetRoutes resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DeviceSubnetRoutesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DeviceSubnetRoutesArgs | DeviceSubnetRoutesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DeviceSubnetRoutesState | undefined;
            resourceInputs["deviceId"] = state ? state.deviceId : undefined;
            resourceInputs["routes"] = state ? state.routes : undefined;
        } else {
            const args = argsOrState as DeviceSubnetRoutesArgs | undefined;
            if ((!args || args.deviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'deviceId'");
            }
            if ((!args || args.routes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'routes'");
            }
            resourceInputs["deviceId"] = args ? args.deviceId : undefined;
            resourceInputs["routes"] = args ? args.routes : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DeviceSubnetRoutes.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DeviceSubnetRoutes resources.
 */
export interface DeviceSubnetRoutesState {
    /**
     * The device to set subnet routes for
     */
    deviceId?: pulumi.Input<string>;
    /**
     * The subnet routes that are enabled to be routed by a device
     */
    routes?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a DeviceSubnetRoutes resource.
 */
export interface DeviceSubnetRoutesArgs {
    /**
     * The device to set subnet routes for
     */
    deviceId: pulumi.Input<string>;
    /**
     * The subnet routes that are enabled to be routed by a device
     */
    routes: pulumi.Input<pulumi.Input<string>[]>;
}
