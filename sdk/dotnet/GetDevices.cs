// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale
{
    public static class GetDevices
    {
        /// <summary>
        /// The devices data source describes a list of devices in a tailnet.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Tailscale = Pulumi.Tailscale;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var sampleDevices = Output.Create(Tailscale.GetDevices.InvokeAsync(new Tailscale.GetDevicesArgs
        ///         {
        ///             NamePrefix = "example-",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDevicesResult> InvokeAsync(GetDevicesArgs? args = null, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDevicesResult>("tailscale:index/getDevices:getDevices", args ?? new GetDevicesArgs(), options.WithVersion());
    }


    public sealed class GetDevicesArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filters the returned list of devices to those whose name have this prefix.
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        public GetDevicesArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDevicesResult
    {
        /// <summary>
        /// The list of devices returned from the Tailscale API. Each element contains the following:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDevicesDeviceResult> Devices;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string? NamePrefix;

        [OutputConstructor]
        private GetDevicesResult(
            ImmutableArray<Outputs.GetDevicesDeviceResult> devices,

            string id,

            string? namePrefix)
        {
            Devices = devices;
            Id = id;
            NamePrefix = namePrefix;
        }
    }
}
