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
    /// The tailnet_key resource allows you to create pre-authentication keys that can register new nodes without needing to sign in via a web browser. See https://tailscale.com/kb/1085/auth-keys for more information
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Tailscale = Pulumi.Tailscale;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var sampleKey = new Tailscale.TailnetKey("sampleKey", new()
    ///     {
    ///         Description = "Sample key",
    ///         Ephemeral = false,
    ///         Expiry = 3600,
    ///         Preauthorized = true,
    ///         Reusable = true,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/tailnetKey:TailnetKey")]
    public partial class TailnetKey : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The creation timestamp of the key in RFC3339 format
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// A description of the key consisting of alphanumeric characters. Defaults to `""`.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Indicates if the key is ephemeral. Defaults to `false`.
        /// </summary>
        [Output("ephemeral")]
        public Output<bool?> Ephemeral { get; private set; } = null!;

        /// <summary>
        /// The expiry timestamp of the key in RFC3339 format
        /// </summary>
        [Output("expiresAt")]
        public Output<string> ExpiresAt { get; private set; } = null!;

        /// <summary>
        /// The expiry of the key in seconds. Defaults to `7776000` (90 days).
        /// </summary>
        [Output("expiry")]
        public Output<int?> Expiry { get; private set; } = null!;

        /// <summary>
        /// Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
        /// </summary>
        [Output("invalid")]
        public Output<bool> Invalid { get; private set; } = null!;

        /// <summary>
        /// The authentication key
        /// </summary>
        [Output("key")]
        public Output<string> Key { get; private set; } = null!;

        /// <summary>
        /// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        /// </summary>
        [Output("preauthorized")]
        public Output<bool?> Preauthorized { get; private set; } = null!;

        /// <summary>
        /// Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        /// </summary>
        [Output("recreateIfInvalid")]
        public Output<string?> RecreateIfInvalid { get; private set; } = null!;

        /// <summary>
        /// Indicates if the key is reusable or single-use. Defaults to `false`.
        /// </summary>
        [Output("reusable")]
        public Output<bool?> Reusable { get; private set; } = null!;

        /// <summary>
        /// List of tags to apply to the machines authenticated by the key.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a TailnetKey resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public TailnetKey(string name, TailnetKeyArgs? args = null, CustomResourceOptions? options = null)
            : base("tailscale:index/tailnetKey:TailnetKey", name, args ?? new TailnetKeyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private TailnetKey(string name, Input<string> id, TailnetKeyState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/tailnetKey:TailnetKey", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "key",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing TailnetKey resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static TailnetKey Get(string name, Input<string> id, TailnetKeyState? state = null, CustomResourceOptions? options = null)
        {
            return new TailnetKey(name, id, state, options);
        }
    }

    public sealed class TailnetKeyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A description of the key consisting of alphanumeric characters. Defaults to `""`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates if the key is ephemeral. Defaults to `false`.
        /// </summary>
        [Input("ephemeral")]
        public Input<bool>? Ephemeral { get; set; }

        /// <summary>
        /// The expiry of the key in seconds. Defaults to `7776000` (90 days).
        /// </summary>
        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        /// <summary>
        /// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        /// </summary>
        [Input("preauthorized")]
        public Input<bool>? Preauthorized { get; set; }

        /// <summary>
        /// Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        /// </summary>
        [Input("recreateIfInvalid")]
        public Input<string>? RecreateIfInvalid { get; set; }

        /// <summary>
        /// Indicates if the key is reusable or single-use. Defaults to `false`.
        /// </summary>
        [Input("reusable")]
        public Input<bool>? Reusable { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags to apply to the machines authenticated by the key.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public TailnetKeyArgs()
        {
        }
        public static new TailnetKeyArgs Empty => new TailnetKeyArgs();
    }

    public sealed class TailnetKeyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The creation timestamp of the key in RFC3339 format
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// A description of the key consisting of alphanumeric characters. Defaults to `""`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Indicates if the key is ephemeral. Defaults to `false`.
        /// </summary>
        [Input("ephemeral")]
        public Input<bool>? Ephemeral { get; set; }

        /// <summary>
        /// The expiry timestamp of the key in RFC3339 format
        /// </summary>
        [Input("expiresAt")]
        public Input<string>? ExpiresAt { get; set; }

        /// <summary>
        /// The expiry of the key in seconds. Defaults to `7776000` (90 days).
        /// </summary>
        [Input("expiry")]
        public Input<int>? Expiry { get; set; }

        /// <summary>
        /// Indicates whether the key is invalid (e.g. expired, revoked or has been deleted).
        /// </summary>
        [Input("invalid")]
        public Input<bool>? Invalid { get; set; }

        [Input("key")]
        private Input<string>? _key;

        /// <summary>
        /// The authentication key
        /// </summary>
        public Input<string>? Key
        {
            get => _key;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _key = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Determines whether or not the machines authenticated by the key will be authorized for the tailnet by default. Defaults to `false`.
        /// </summary>
        [Input("preauthorized")]
        public Input<bool>? Preauthorized { get; set; }

        /// <summary>
        /// Determines whether the key should be created again if it becomes invalid. By default, reusable keys will be recreated, but single-use keys will not. Possible values: 'always', 'never'.
        /// </summary>
        [Input("recreateIfInvalid")]
        public Input<string>? RecreateIfInvalid { get; set; }

        /// <summary>
        /// Indicates if the key is reusable or single-use. Defaults to `false`.
        /// </summary>
        [Input("reusable")]
        public Input<bool>? Reusable { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// List of tags to apply to the machines authenticated by the key.
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public TailnetKeyState()
        {
        }
        public static new TailnetKeyState Empty => new TailnetKeyState();
    }
}
