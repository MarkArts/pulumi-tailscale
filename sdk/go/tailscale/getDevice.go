// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The device data source describes a single device in a tailnet.
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
// 		_, err := tailscale.GetDevice(ctx, &GetDeviceArgs{
// 			Name:    "user1-device.example.com",
// 			WaitFor: pulumi.StringRef("60s"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func GetDevice(ctx *pulumi.Context, args *GetDeviceArgs, opts ...pulumi.InvokeOption) (*GetDeviceResult, error) {
	var rv GetDeviceResult
	err := ctx.Invoke("tailscale:index/getDevice:getDevice", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDevice.
type GetDeviceArgs struct {
	// The name of the tailnet device to obtain the attributes of.
	Name string `pulumi:"name"`
	// If specified, the provider will retry obtaining the device data every second until the specified duration has been reached. Must be a value greater than 1 second
	WaitFor *string `pulumi:"waitFor"`
}

// A collection of values returned by getDevice.
type GetDeviceResult struct {
	// Tailscale IPs for the device
	Addresses []string `pulumi:"addresses"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Tags applied to the device
	Tags []string `pulumi:"tags"`
	// The user associated with the device
	User    string  `pulumi:"user"`
	WaitFor *string `pulumi:"waitFor"`
}

func GetDeviceOutput(ctx *pulumi.Context, args GetDeviceOutputArgs, opts ...pulumi.InvokeOption) GetDeviceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDeviceResult, error) {
			args := v.(GetDeviceArgs)
			r, err := GetDevice(ctx, &args, opts...)
			var s GetDeviceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDeviceResultOutput)
}

// A collection of arguments for invoking getDevice.
type GetDeviceOutputArgs struct {
	// The name of the tailnet device to obtain the attributes of.
	Name pulumi.StringInput `pulumi:"name"`
	// If specified, the provider will retry obtaining the device data every second until the specified duration has been reached. Must be a value greater than 1 second
	WaitFor pulumi.StringPtrInput `pulumi:"waitFor"`
}

func (GetDeviceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceArgs)(nil)).Elem()
}

// A collection of values returned by getDevice.
type GetDeviceResultOutput struct{ *pulumi.OutputState }

func (GetDeviceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDeviceResult)(nil)).Elem()
}

func (o GetDeviceResultOutput) ToGetDeviceResultOutput() GetDeviceResultOutput {
	return o
}

func (o GetDeviceResultOutput) ToGetDeviceResultOutputWithContext(ctx context.Context) GetDeviceResultOutput {
	return o
}

// Tailscale IPs for the device
func (o GetDeviceResultOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeviceResult) []string { return v.Addresses }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetDeviceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Tags applied to the device
func (o GetDeviceResultOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDeviceResult) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

// The user associated with the device
func (o GetDeviceResultOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v GetDeviceResult) string { return v.User }).(pulumi.StringOutput)
}

func (o GetDeviceResultOutput) WaitFor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetDeviceResult) *string { return v.WaitFor }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDeviceResultOutput{})
}
