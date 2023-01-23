// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale
{
    /// <summary>
    /// The acl resource allows you to configure a Tailscale ACL. See https://tailscale.com/kb/1018/acls for more information. Note that this resource will completely overwrite existing ACL contents for a given tailnet.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Tailscale = Pulumi.Tailscale;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sampleAcl = new Tailscale.Acl("sampleAcl", new()
    ///     {
    ///         AclJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["acls"] = new[]
    ///             {
    ///                 new Dictionary&lt;string, object?&gt;
    ///                 {
    ///                     ["action"] = "accept",
    ///                     ["users"] = new[]
    ///                     {
    ///                         "*",
    ///                     },
    ///                     ["ports"] = new[]
    ///                     {
    ///                         "*:*",
    ///                     },
    ///                 },
    ///             },
    ///         }),
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/acl:Acl")]
    public partial class Acl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The JSON-based policy that defines which devices and users are allowed to connect in your network
        /// </summary>
        [Output("acl")]
        public Output<string> AclJson { get; private set; } = null!;


        /// <summary>
        /// Create a Acl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Acl(string name, AclArgs args, CustomResourceOptions? options = null)
            : base("tailscale:index/acl:Acl", name, args ?? new AclArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Acl(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/acl:Acl", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Acl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Acl Get(string name, Input<string> id, AclState? state = null, CustomResourceOptions? options = null)
        {
            return new Acl(name, id, state, options);
        }
    }

    public sealed class AclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The JSON-based policy that defines which devices and users are allowed to connect in your network
        /// </summary>
        [Input("acl", required: true)]
        public Input<string> AclJson { get; set; } = null!;

        public AclArgs()
        {
        }
        public static new AclArgs Empty => new AclArgs();
    }

    public sealed class AclState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The JSON-based policy that defines which devices and users are allowed to connect in your network
        /// </summary>
        [Input("acl")]
        public Input<string>? AclJson { get; set; }

        public AclState()
        {
        }
        public static new AclState Empty => new AclState();
    }
}
