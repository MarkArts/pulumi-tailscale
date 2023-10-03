// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Tailscale.Outputs
{

    [OutputType]
    public sealed class GetDevicesDeviceResult
    {
        public readonly ImmutableArray<string> Addresses;
        public readonly string Id;
        public readonly string Name;
        public readonly ImmutableArray<string> Tags;
        public readonly string User;

        [OutputConstructor]
        private GetDevicesDeviceResult(
            ImmutableArray<string> addresses,

            string id,

            string name,

            ImmutableArray<string> tags,

            string user)
        {
            Addresses = addresses;
            Id = id;
            Name = name;
            Tags = tags;
            User = user;
        }
    }
}
