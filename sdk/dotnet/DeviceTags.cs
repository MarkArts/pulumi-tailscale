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
    /// The device_tags resource is used to apply tags to Tailscale devices. See https://tailscale.com/kb/1068/acl-tags/ for more details.
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
    ///     var sampleTags = new Tailscale.DeviceTags("sampleTags", new()
    ///     {
    ///         DeviceId = sampleDevice.Apply(getDeviceResult =&gt; getDeviceResult.Id),
    ///         Tags = new[]
    ///         {
    ///             "room:bedroom",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/deviceTags:DeviceTags")]
    public partial class DeviceTags : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The device to set tags for
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;

        /// <summary>
        /// The tags to apply to the device
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceTags resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceTags(string name, DeviceTagsArgs args, CustomResourceOptions? options = null)
            : base("tailscale:index/deviceTags:DeviceTags", name, args ?? new DeviceTagsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceTags(string name, Input<string> id, DeviceTagsState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/deviceTags:DeviceTags", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceTags resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceTags Get(string name, Input<string> id, DeviceTagsState? state = null, CustomResourceOptions? options = null)
        {
            return new DeviceTags(name, id, state, options);
        }
    }

    public sealed class DeviceTagsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device to set tags for
        /// </summary>
        [Input("deviceId", required: true)]
        public Input<string> DeviceId { get; set; } = null!;

        [Input("tags", required: true)]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags to apply to the device
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DeviceTagsArgs()
        {
        }
        public static new DeviceTagsArgs Empty => new DeviceTagsArgs();
    }

    public sealed class DeviceTagsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The device to set tags for
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;

        /// <summary>
        /// The tags to apply to the device
        /// </summary>
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        public DeviceTagsState()
        {
        }
        public static new DeviceTagsState Empty => new DeviceTagsState();
    }
}
