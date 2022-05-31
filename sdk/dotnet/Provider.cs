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
    /// The provider type for the tailscale package. By default, resources use package-wide configuration
    /// settings, however an explicit `Provider` instance may be created and passed during resource
    /// construction to achieve fine-grained programmatic control over provider settings. See the
    /// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
    /// </summary>
    [TailscaleResourceType("pulumi:providers:tailscale")]
    public partial class Provider : Pulumi.ProviderResource
    {
        /// <summary>
        /// The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
        /// </summary>
        [Output("apiKey")]
        public Output<string> ApiKey { get; private set; } = null!;

        /// <summary>
        /// The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
        /// environment variable.
        /// </summary>
        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// The Tailnet to perform actions in. Can be set via the TAILSCALE_TAILNET environment variable.
        /// </summary>
        [Output("tailnet")]
        public Output<string> Tailnet { get; private set; } = null!;


        /// <summary>
        /// Create a Provider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Provider(string name, ProviderArgs args, CustomResourceOptions? options = null)
            : base("tailscale", name, args ?? new ProviderArgs(), MakeResourceOptions(options, ""))
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
    }

    public sealed class ProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The API key to use for authenticating requests to the API. Can be set via the TAILSCALE_API_KEY environment variable.
        /// </summary>
        [Input("apiKey", required: true)]
        public Input<string> ApiKey { get; set; } = null!;

        /// <summary>
        /// The base URL of the Tailscale API. Defaults to https://api.tailscale.com. Can be set via the TAILSCALE_BASE_URL
        /// environment variable.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// The Tailnet to perform actions in. Can be set via the TAILSCALE_TAILNET environment variable.
        /// </summary>
        [Input("tailnet", required: true)]
        public Input<string> Tailnet { get; set; } = null!;

        public ProviderArgs()
        {
        }
    }
}
