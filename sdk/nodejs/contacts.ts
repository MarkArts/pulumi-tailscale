// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as tailscale from "@pulumi/tailscale";
 *
 * const sampleContacts = new tailscale.Contacts("sample_contacts", {
 *     account: {
 *         email: "account@example.com",
 *     },
 *     support: {
 *         email: "support@example.com",
 *     },
 *     security: {
 *         email: "security@example.com",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * ID doesn't matter.
 *
 * ```sh
 * $ pulumi import tailscale:index/contacts:Contacts sample_contacts contacts
 * ```
 */
export class Contacts extends pulumi.CustomResource {
    /**
     * Get an existing Contacts resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ContactsState, opts?: pulumi.CustomResourceOptions): Contacts {
        return new Contacts(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'tailscale:index/contacts:Contacts';

    /**
     * Returns true if the given object is an instance of Contacts.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Contacts {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Contacts.__pulumiType;
    }

    /**
     * Configuration for communications about important changes to your tailnet
     */
    public readonly account!: pulumi.Output<outputs.ContactsAccount>;
    /**
     * Configuration for communications about security issues affecting your tailnet
     */
    public readonly security!: pulumi.Output<outputs.ContactsSecurity>;
    /**
     * Configuration for communications about misconfigurations in your tailnet
     */
    public readonly support!: pulumi.Output<outputs.ContactsSupport>;

    /**
     * Create a Contacts resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ContactsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ContactsArgs | ContactsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ContactsState | undefined;
            resourceInputs["account"] = state ? state.account : undefined;
            resourceInputs["security"] = state ? state.security : undefined;
            resourceInputs["support"] = state ? state.support : undefined;
        } else {
            const args = argsOrState as ContactsArgs | undefined;
            if ((!args || args.account === undefined) && !opts.urn) {
                throw new Error("Missing required property 'account'");
            }
            if ((!args || args.security === undefined) && !opts.urn) {
                throw new Error("Missing required property 'security'");
            }
            if ((!args || args.support === undefined) && !opts.urn) {
                throw new Error("Missing required property 'support'");
            }
            resourceInputs["account"] = args ? args.account : undefined;
            resourceInputs["security"] = args ? args.security : undefined;
            resourceInputs["support"] = args ? args.support : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Contacts.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Contacts resources.
 */
export interface ContactsState {
    /**
     * Configuration for communications about important changes to your tailnet
     */
    account?: pulumi.Input<inputs.ContactsAccount>;
    /**
     * Configuration for communications about security issues affecting your tailnet
     */
    security?: pulumi.Input<inputs.ContactsSecurity>;
    /**
     * Configuration for communications about misconfigurations in your tailnet
     */
    support?: pulumi.Input<inputs.ContactsSupport>;
}

/**
 * The set of arguments for constructing a Contacts resource.
 */
export interface ContactsArgs {
    /**
     * Configuration for communications about important changes to your tailnet
     */
    account: pulumi.Input<inputs.ContactsAccount>;
    /**
     * Configuration for communications about security issues affecting your tailnet
     */
    security: pulumi.Input<inputs.ContactsSecurity>;
    /**
     * Configuration for communications about misconfigurations in your tailnet
     */
    support: pulumi.Input<inputs.ContactsSupport>;
}
