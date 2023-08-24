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
    /// The dns_nameservers resource allows you to configure DNS nameservers for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
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
    ///     var sampleSearchPaths = new Tailscale.DnsSearchPaths("sampleSearchPaths", new()
    ///     {
    ///         SearchPaths = new[]
    ///         {
    ///             "example.com",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/dnsSearchPaths:DnsSearchPaths")]
    public partial class DnsSearchPaths : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Devices on your network will use these domain suffixes to resolve DNS names.
        /// </summary>
        [Output("searchPaths")]
        public Output<ImmutableArray<string>> SearchPaths { get; private set; } = null!;


        /// <summary>
        /// Create a DnsSearchPaths resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DnsSearchPaths(string name, DnsSearchPathsArgs args, CustomResourceOptions? options = null)
            : base("tailscale:index/dnsSearchPaths:DnsSearchPaths", name, args ?? new DnsSearchPathsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DnsSearchPaths(string name, Input<string> id, DnsSearchPathsState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/dnsSearchPaths:DnsSearchPaths", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DnsSearchPaths resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DnsSearchPaths Get(string name, Input<string> id, DnsSearchPathsState? state = null, CustomResourceOptions? options = null)
        {
            return new DnsSearchPaths(name, id, state, options);
        }
    }

    public sealed class DnsSearchPathsArgs : global::Pulumi.ResourceArgs
    {
        [Input("searchPaths", required: true)]
        private InputList<string>? _searchPaths;

        /// <summary>
        /// Devices on your network will use these domain suffixes to resolve DNS names.
        /// </summary>
        public InputList<string> SearchPaths
        {
            get => _searchPaths ?? (_searchPaths = new InputList<string>());
            set => _searchPaths = value;
        }

        public DnsSearchPathsArgs()
        {
        }
        public static new DnsSearchPathsArgs Empty => new DnsSearchPathsArgs();
    }

    public sealed class DnsSearchPathsState : global::Pulumi.ResourceArgs
    {
        [Input("searchPaths")]
        private InputList<string>? _searchPaths;

        /// <summary>
        /// Devices on your network will use these domain suffixes to resolve DNS names.
        /// </summary>
        public InputList<string> SearchPaths
        {
            get => _searchPaths ?? (_searchPaths = new InputList<string>());
            set => _searchPaths = value;
        }

        public DnsSearchPathsState()
        {
        }
        public static new DnsSearchPathsState Empty => new DnsSearchPathsState();
    }
}
