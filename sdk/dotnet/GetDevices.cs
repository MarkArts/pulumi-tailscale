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
        /// The devices data source describes a list of devices in a tailnet
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tailscale = Pulumi.Tailscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sampleDevices = Tailscale.GetDevices.Invoke(new()
        ///     {
        ///         NamePrefix = "example-",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetDevicesResult> InvokeAsync(GetDevicesArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDevicesResult>("tailscale:index/getDevices:getDevices", args ?? new GetDevicesArgs(), options.WithDefaults());

        /// <summary>
        /// The devices data source describes a list of devices in a tailnet
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Tailscale = Pulumi.Tailscale;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var sampleDevices = Tailscale.GetDevices.Invoke(new()
        ///     {
        ///         NamePrefix = "example-",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetDevicesResult> Invoke(GetDevicesInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDevicesResult>("tailscale:index/getDevices:getDevices", args ?? new GetDevicesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDevicesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filters the device list to elements whose name has the provided prefix
        /// </summary>
        [Input("namePrefix")]
        public string? NamePrefix { get; set; }

        public GetDevicesArgs()
        {
        }
        public static new GetDevicesArgs Empty => new GetDevicesArgs();
    }

    public sealed class GetDevicesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Filters the device list to elements whose name has the provided prefix
        /// </summary>
        [Input("namePrefix")]
        public Input<string>? NamePrefix { get; set; }

        public GetDevicesInvokeArgs()
        {
        }
        public static new GetDevicesInvokeArgs Empty => new GetDevicesInvokeArgs();
    }


    [OutputType]
    public sealed class GetDevicesResult
    {
        /// <summary>
        /// The list of devices in the tailnet
        /// </summary>
        public readonly ImmutableArray<Outputs.GetDevicesDeviceResult> Devices;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Filters the device list to elements whose name has the provided prefix
        /// </summary>
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
