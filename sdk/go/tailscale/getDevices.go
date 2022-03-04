// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The devices data source describes a list of devices in a tailnet.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tailscale.GetDevices(ctx, &GetDevicesArgs{
// 			NamePrefix: pulumi.StringRef("example-"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDevices(ctx *pulumi.Context, args *GetDevicesArgs, opts ...pulumi.InvokeOption) (*GetDevicesResult, error) {
	var rv GetDevicesResult
	err := ctx.Invoke("tailscale:index/getDevices:getDevices", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDevices.
type GetDevicesArgs struct {
	// Filters the returned list of devices to those whose name have this prefix.
	NamePrefix *string `pulumi:"namePrefix"`
}

// A collection of values returned by getDevices.
type GetDevicesResult struct {
	// The list of devices returned from the Tailscale API. Each element contains the following:
	Devices []GetDevicesDevice `pulumi:"devices"`
	// The provider-assigned unique ID for this managed resource.
	Id         string  `pulumi:"id"`
	NamePrefix *string `pulumi:"namePrefix"`
}

func GetDevicesOutput(ctx *pulumi.Context, args GetDevicesOutputArgs, opts ...pulumi.InvokeOption) GetDevicesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDevicesResult, error) {
			args := v.(GetDevicesArgs)
			r, err := GetDevices(ctx, &args, opts...)
			return *r, err
		}).(GetDevicesResultOutput)
}

// A collection of arguments for invoking getDevices.
type GetDevicesOutputArgs struct {
	// Filters the returned list of devices to those whose name have this prefix.
	NamePrefix pulumi.StringPtrInput `pulumi:"namePrefix"`
}

func (GetDevicesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDevicesArgs)(nil)).Elem()
}

// A collection of values returned by getDevices.
type GetDevicesResultOutput struct{ *pulumi.OutputState }

func (GetDevicesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDevicesResult)(nil)).Elem()
}

func (o GetDevicesResultOutput) ToGetDevicesResultOutput() GetDevicesResultOutput {
	return o
}

func (o GetDevicesResultOutput) ToGetDevicesResultOutputWithContext(ctx context.Context) GetDevicesResultOutput {
	return o
}

// The list of devices returned from the Tailscale API. Each element contains the following:
func (o GetDevicesResultOutput) Devices() GetDevicesDeviceArrayOutput {
	return o.ApplyT(func(v GetDevicesResult) []GetDevicesDevice { return v.Devices }).(GetDevicesDeviceArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDevicesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDevicesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDevicesResultOutput) NamePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDevicesResult) *string { return v.NamePrefix }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDevicesResultOutput{})
}
