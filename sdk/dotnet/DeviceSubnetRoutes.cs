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
    /// The device_subnet_routes resource allows you to configure subnet routes for your Tailscale devices. See https://tailscale.com/kb/1019/subnets for more information.
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
    ///     var sampleDevice = Tailscale.GetDevice.Invoke(new()
    ///     {
    ///         Name = "device.example.com",
    ///     });
    /// 
    ///     var sampleRoutes = new Tailscale.DeviceSubnetRoutes("sampleRoutes", new()
    ///     {
    ///         DeviceId = sampleDevice.Apply(getDeviceResult =&gt; getDeviceResult.Id),
    ///         Routes = new[]
    ///         {
    ///             "10.0.1.0/24",
    ///             "1.2.0.0/16",
    ///             "2.0.0.0/24",
    ///         },
    ///     });
    /// 
    ///     var sampleExitNode = new Tailscale.DeviceSubnetRoutes("sampleExitNode", new()
    ///     {
    ///         DeviceId = sampleDevice.Apply(getDeviceResult =&gt; getDeviceResult.Id),
    ///         Routes = new[]
    ///         {
    ///             "0.0.0.0/0",
    ///             "::/0",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes")]
    public partial class DeviceSubnetRoutes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The device to set subnet routes for
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// The subnet routes that are enabled to be routed by a device
        /// </summary>
        [Output("routes")]
        public Output<ImmutableArray<string>> Routes { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceSubnetRoutes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceSubnetRoutes(string name, DeviceSubnetRoutesArgs args, CustomResourceOptions? options = null)
            : base("tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes", name, args ?? new DeviceSubnetRoutesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceSubnetRoutes(string name, Input<string> id, DeviceSubnetRoutesState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceSubnetRoutes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceSubnetRoutes Get(string name, Input<string> id, DeviceSubnetRoutesState? state = null, CustomResourceOptions? options = null)
        {
            return new DeviceSubnetRoutes(name, id, state, options);
        }
    }

    public sealed class DeviceSubnetRoutesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device to set subnet routes for
        /// </summary>
        [Input("deviceId", required: true)]
        public Input<string> DeviceId { get; set; } = null!;

        [Input("routes", required: true)]
        private InputList<string>? _routes;

        /// <summary>
        /// The subnet routes that are enabled to be routed by a device
        /// </summary>
        public InputList<string> Routes
        {
            get => _routes ?? (_routes = new InputList<string>());
            set => _routes = value;
        }

        public DeviceSubnetRoutesArgs()
        {
        }
        public static new DeviceSubnetRoutesArgs Empty => new DeviceSubnetRoutesArgs();
    }

    public sealed class DeviceSubnetRoutesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device to set subnet routes for
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        [Input("routes")]
        private InputList<string>? _routes;

        /// <summary>
        /// The subnet routes that are enabled to be routed by a device
        /// </summary>
        public InputList<string> Routes
        {
            get => _routes ?? (_routes = new InputList<string>());
            set => _routes = value;
        }

        public DeviceSubnetRoutesState()
        {
        }
        public static new DeviceSubnetRoutesState Empty => new DeviceSubnetRoutesState();
    }
}
