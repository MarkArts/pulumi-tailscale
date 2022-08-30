// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GetDevicesDevice struct {
	Addresses []string `pulumi:"addresses"`
	// The ID of this resource.
	Id   string   `pulumi:"id"`
	Name string   `pulumi:"name"`
	Tags []string `pulumi:"tags"`
	User string   `pulumi:"user"`
}

// GetDevicesDeviceInput is an input type that accepts GetDevicesDeviceArgs and GetDevicesDeviceOutput values.
// You can construct a concrete instance of `GetDevicesDeviceInput` via:
//
//	GetDevicesDeviceArgs{...}
type GetDevicesDeviceInput interface {
	pulumi.Input

	ToGetDevicesDeviceOutput() GetDevicesDeviceOutput
	ToGetDevicesDeviceOutputWithContext(context.Context) GetDevicesDeviceOutput
}

type GetDevicesDeviceArgs struct {
	Addresses pulumi.StringArrayInput `pulumi:"addresses"`
	// The ID of this resource.
	Id   pulumi.StringInput      `pulumi:"id"`
	Name pulumi.StringInput      `pulumi:"name"`
	Tags pulumi.StringArrayInput `pulumi:"tags"`
	User pulumi.StringInput      `pulumi:"user"`
}

func (GetDevicesDeviceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDevicesDevice)(nil)).Elem()
}

func (i GetDevicesDeviceArgs) ToGetDevicesDeviceOutput() GetDevicesDeviceOutput {
	return i.ToGetDevicesDeviceOutputWithContext(context.Background())
}

func (i GetDevicesDeviceArgs) ToGetDevicesDeviceOutputWithContext(ctx context.Context) GetDevicesDeviceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDevicesDeviceOutput)
}

// GetDevicesDeviceArrayInput is an input type that accepts GetDevicesDeviceArray and GetDevicesDeviceArrayOutput values.
// You can construct a concrete instance of `GetDevicesDeviceArrayInput` via:
//
//	GetDevicesDeviceArray{ GetDevicesDeviceArgs{...} }
type GetDevicesDeviceArrayInput interface {
	pulumi.Input

	ToGetDevicesDeviceArrayOutput() GetDevicesDeviceArrayOutput
	ToGetDevicesDeviceArrayOutputWithContext(context.Context) GetDevicesDeviceArrayOutput
}

type GetDevicesDeviceArray []GetDevicesDeviceInput

func (GetDevicesDeviceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDevicesDevice)(nil)).Elem()
}

func (i GetDevicesDeviceArray) ToGetDevicesDeviceArrayOutput() GetDevicesDeviceArrayOutput {
	return i.ToGetDevicesDeviceArrayOutputWithContext(context.Background())
}

func (i GetDevicesDeviceArray) ToGetDevicesDeviceArrayOutputWithContext(ctx context.Context) GetDevicesDeviceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetDevicesDeviceArrayOutput)
}

type GetDevicesDeviceOutput struct{ *pulumi.OutputState }

func (GetDevicesDeviceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDevicesDevice)(nil)).Elem()
}

func (o GetDevicesDeviceOutput) ToGetDevicesDeviceOutput() GetDevicesDeviceOutput {
	return o
}

func (o GetDevicesDeviceOutput) ToGetDevicesDeviceOutputWithContext(ctx context.Context) GetDevicesDeviceOutput {
	return o
}

func (o GetDevicesDeviceOutput) Addresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDevicesDevice) []string { return v.Addresses }).(pulumi.StringArrayOutput)
}

// The ID of this resource.
func (o GetDevicesDeviceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDevicesDevice) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetDevicesDeviceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDevicesDevice) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDevicesDeviceOutput) Tags() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetDevicesDevice) []string { return v.Tags }).(pulumi.StringArrayOutput)
}

func (o GetDevicesDeviceOutput) User() pulumi.StringOutput {
	return o.ApplyT(func(v GetDevicesDevice) string { return v.User }).(pulumi.StringOutput)
}

type GetDevicesDeviceArrayOutput struct{ *pulumi.OutputState }

func (GetDevicesDeviceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetDevicesDevice)(nil)).Elem()
}

func (o GetDevicesDeviceArrayOutput) ToGetDevicesDeviceArrayOutput() GetDevicesDeviceArrayOutput {
	return o
}

func (o GetDevicesDeviceArrayOutput) ToGetDevicesDeviceArrayOutputWithContext(ctx context.Context) GetDevicesDeviceArrayOutput {
	return o
}

func (o GetDevicesDeviceArrayOutput) Index(i pulumi.IntInput) GetDevicesDeviceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetDevicesDevice {
		return vs[0].([]GetDevicesDevice)[vs[1].(int)]
	}).(GetDevicesDeviceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetDevicesDeviceInput)(nil)).Elem(), GetDevicesDeviceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetDevicesDeviceArrayInput)(nil)).Elem(), GetDevicesDeviceArray{})
	pulumi.RegisterOutputType(GetDevicesDeviceOutput{})
	pulumi.RegisterOutputType(GetDevicesDeviceArrayOutput{})
}
