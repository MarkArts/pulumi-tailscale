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
    /// The posture_integration resource allows you to manage integrations with device posture data providers. See https://tailscale.com/kb/1288/device-posture for more information.
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
    ///     var samplePostureIntegration = new Tailscale.PostureIntegration("sample_posture_integration", new()
    ///     {
    ///         PostureProvider = "falcon",
    ///         CloudId = "us-1",
    ///         ClientId = "clientid1",
    ///         ClientSecret = "test-secret1",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/postureIntegration:PostureIntegration")]
    public partial class PostureIntegration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Unique identifier for your client.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The secret (auth key, token, etc.) used to authenticate with the provider.
        /// </summary>
        [Output("clientSecret")]
        public Output<string> ClientSecret { get; private set; } = null!;

        /// <summary>
        /// Identifies which of the provider's clouds to integrate with.
        /// </summary>
        [Output("cloudId")]
        public Output<string?> CloudId { get; private set; } = null!;

        /// <summary>
        /// The type of posture integration data provider.
        /// </summary>
        [Output("postureProvider")]
        public Output<string> PostureProvider { get; private set; } = null!;

        /// <summary>
        /// The Microsoft Intune directory (tenant) ID. For other providers, this is left blank.
        /// </summary>
        [Output("tenantId")]
        public Output<string?> TenantId { get; private set; } = null!;


        /// <summary>
        /// Create a PostureIntegration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PostureIntegration(string name, PostureIntegrationArgs args, CustomResourceOptions? options = null)
            : base("tailscale:index/postureIntegration:PostureIntegration", name, args ?? new PostureIntegrationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PostureIntegration(string name, Input<string> id, PostureIntegrationState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/postureIntegration:PostureIntegration", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PostureIntegration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PostureIntegration Get(string name, Input<string> id, PostureIntegrationState? state = null, CustomResourceOptions? options = null)
        {
            return new PostureIntegration(name, id, state, options);
        }
    }

    public sealed class PostureIntegrationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for your client.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The secret (auth key, token, etc.) used to authenticate with the provider.
        /// </summary>
        [Input("clientSecret", required: true)]
        public Input<string> ClientSecret { get; set; } = null!;

        /// <summary>
        /// Identifies which of the provider's clouds to integrate with.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        /// <summary>
        /// The type of posture integration data provider.
        /// </summary>
        [Input("postureProvider", required: true)]
        public Input<string> PostureProvider { get; set; } = null!;

        /// <summary>
        /// The Microsoft Intune directory (tenant) ID. For other providers, this is left blank.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public PostureIntegrationArgs()
        {
        }
        public static new PostureIntegrationArgs Empty => new PostureIntegrationArgs();
    }

    public sealed class PostureIntegrationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique identifier for your client.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The secret (auth key, token, etc.) used to authenticate with the provider.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Identifies which of the provider's clouds to integrate with.
        /// </summary>
        [Input("cloudId")]
        public Input<string>? CloudId { get; set; }

        /// <summary>
        /// The type of posture integration data provider.
        /// </summary>
        [Input("postureProvider")]
        public Input<string>? PostureProvider { get; set; }

        /// <summary>
        /// The Microsoft Intune directory (tenant) ID. For other providers, this is left blank.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        public PostureIntegrationState()
        {
        }
        public static new PostureIntegrationState Empty => new PostureIntegrationState();
    }
}
