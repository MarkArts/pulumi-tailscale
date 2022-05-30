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
    /// The device_authorization resource is used to approve new devices before they can join the tailnet. See https://tailscale.com/kb/1099/device-authorization/ for more details.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Tailscale = Pulumi.Tailscale;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var sampleDevice = Output.Create(Tailscale.GetDevice.InvokeAsync(new Tailscale.GetDeviceArgs
    ///         {
    ///             Name = "device.example.com",
    ///         }));
    ///         var sampleAuthorization = new Tailscale.DeviceAuthorization("sampleAuthorization", new Tailscale.DeviceAuthorizationArgs
    ///         {
    ///             DeviceId = sampleDevice.Apply(sampleDevice =&gt; sampleDevice.Id),
    ///             Authorized = true,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [TailscaleResourceType("tailscale:index/deviceAuthorization:DeviceAuthorization")]
    public partial class DeviceAuthorization : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether or not the device is authorized
        /// </summary>
        [Output("authorized")]
        public Output<bool> Authorized { get; private set; } = null!;

        /// <summary>
        /// The device to set as authorized
        /// </summary>
        [Output("deviceId")]
        public Output<string> DeviceId { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceAuthorization resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceAuthorization(string name, DeviceAuthorizationArgs args, CustomResourceOptions? options = null)
            : base("tailscale:index/deviceAuthorization:DeviceAuthorization", name, args ?? new DeviceAuthorizationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceAuthorization(string name, Input<string> id, DeviceAuthorizationState? state = null, CustomResourceOptions? options = null)
            : base("tailscale:index/deviceAuthorization:DeviceAuthorization", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DeviceAuthorization resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceAuthorization Get(string name, Input<string> id, DeviceAuthorizationState? state = null, CustomResourceOptions? options = null)
        {
            return new DeviceAuthorization(name, id, state, options);
        }
    }

    public sealed class DeviceAuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the device is authorized
        /// </summary>
        [Input("authorized", required: true)]
        public Input<bool> Authorized { get; set; } = null!;

        /// <summary>
        /// The device to set as authorized
        /// </summary>
        [Input("deviceId", required: true)]
        public Input<string> DeviceId { get; set; } = null!;

        public DeviceAuthorizationArgs()
        {
        }
    }

    public sealed class DeviceAuthorizationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether or not the device is authorized
        /// </summary>
        [Input("authorized")]
        public Input<bool>? Authorized { get; set; }

        /// <summary>
        /// The device to set as authorized
        /// </summary>
        [Input("deviceId")]
        public Input<string>? DeviceId { get; set; }

        public DeviceAuthorizationState()
        {
        }
    }
}
