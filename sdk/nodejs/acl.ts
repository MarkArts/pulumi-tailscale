// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The acl resource allows you to configure a Tailscale ACL. See the [Tailscale ACL documentation](https://tailscale.com/kb/1018/acls)
 * for more information. You can set the ACL in multiple ways including hujson.
 */
export class Acl extends pulumi.CustomResource {
    /**
     * Get an existing Acl resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AclState, opts?: pulumi.CustomResourceOptions): Acl {
        return new Acl(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/acl:Acl';

    /**
     * Returns true if the given object is an instance of Acl.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Acl {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Acl.__pulumiType;
    }

    /**
     * The JSON-based policy that defines which devices and users are allowed to connect in your network.
     * This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
     * local ACL file.
     */
    public readonly acl!: pulumi.Output<string>;

    /**
     * Create a Acl resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AclArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AclArgs | AclState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AclState | undefined;
            inputs["acl"] = state ? state.acl : undefined;
        } else {
            const args = argsOrState as AclArgs | undefined;
            if ((!args || args.acl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'acl'");
            }
            inputs["acl"] = args ? args.acl : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Acl.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Acl resources.
 */
export interface AclState {
    /**
     * The JSON-based policy that defines which devices and users are allowed to connect in your network.
     * This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
     * local ACL file.
     */
    acl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Acl resource.
 */
export interface AclArgs {
    /**
     * The JSON-based policy that defines which devices and users are allowed to connect in your network.
     * This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
     * local ACL file.
     */
    acl: pulumi.Input<string>;
}
