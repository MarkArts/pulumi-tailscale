// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The deviceKey resource allows you to update the properties of a device's key
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			exampleDevice, err := tailscale.GetDevice(ctx, &tailscale.GetDeviceArgs{
//				Name: pulumi.StringRef("device.example.com"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = tailscale.NewDeviceKey(ctx, "example_key", &tailscale.DeviceKeyArgs{
//				DeviceId:          pulumi.String(exampleDevice.Id),
//				KeyExpiryDisabled: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DeviceKey struct {
	pulumi.CustomResourceState

	// The device to update the key properties of
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// Determines whether or not the device's key will expire. Defaults to `false`.
	KeyExpiryDisabled pulumi.BoolPtrOutput `pulumi:"keyExpiryDisabled"`
}

// NewDeviceKey registers a new resource with the given unique name, arguments, and options.
func NewDeviceKey(ctx *pulumi.Context,
	name string, args *DeviceKeyArgs, opts ...pulumi.ResourceOption) (*DeviceKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DeviceKey
	err := ctx.RegisterResource("tailscale:index/deviceKey:DeviceKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceKey gets an existing DeviceKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceKeyState, opts ...pulumi.ResourceOption) (*DeviceKey, error) {
	var resource DeviceKey
	err := ctx.ReadResource("tailscale:index/deviceKey:DeviceKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceKey resources.
type deviceKeyState struct {
	// The device to update the key properties of
	DeviceId *string `pulumi:"deviceId"`
	// Determines whether or not the device's key will expire. Defaults to `false`.
	KeyExpiryDisabled *bool `pulumi:"keyExpiryDisabled"`
}

type DeviceKeyState struct {
	// The device to update the key properties of
	DeviceId pulumi.StringPtrInput
	// Determines whether or not the device's key will expire. Defaults to `false`.
	KeyExpiryDisabled pulumi.BoolPtrInput
}

func (DeviceKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceKeyState)(nil)).Elem()
}

type deviceKeyArgs struct {
	// The device to update the key properties of
	DeviceId string `pulumi:"deviceId"`
	// Determines whether or not the device's key will expire. Defaults to `false`.
	KeyExpiryDisabled *bool `pulumi:"keyExpiryDisabled"`
}

// The set of arguments for constructing a DeviceKey resource.
type DeviceKeyArgs struct {
	// The device to update the key properties of
	DeviceId pulumi.StringInput
	// Determines whether or not the device's key will expire. Defaults to `false`.
	KeyExpiryDisabled pulumi.BoolPtrInput
}

func (DeviceKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceKeyArgs)(nil)).Elem()
}

type DeviceKeyInput interface {
	pulumi.Input

	ToDeviceKeyOutput() DeviceKeyOutput
	ToDeviceKeyOutputWithContext(ctx context.Context) DeviceKeyOutput
}

func (*DeviceKey) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceKey)(nil)).Elem()
}

func (i *DeviceKey) ToDeviceKeyOutput() DeviceKeyOutput {
	return i.ToDeviceKeyOutputWithContext(context.Background())
}

func (i *DeviceKey) ToDeviceKeyOutputWithContext(ctx context.Context) DeviceKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceKeyOutput)
}

// DeviceKeyArrayInput is an input type that accepts DeviceKeyArray and DeviceKeyArrayOutput values.
// You can construct a concrete instance of `DeviceKeyArrayInput` via:
//
//	DeviceKeyArray{ DeviceKeyArgs{...} }
type DeviceKeyArrayInput interface {
	pulumi.Input

	ToDeviceKeyArrayOutput() DeviceKeyArrayOutput
	ToDeviceKeyArrayOutputWithContext(context.Context) DeviceKeyArrayOutput
}

type DeviceKeyArray []DeviceKeyInput

func (DeviceKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceKey)(nil)).Elem()
}

func (i DeviceKeyArray) ToDeviceKeyArrayOutput() DeviceKeyArrayOutput {
	return i.ToDeviceKeyArrayOutputWithContext(context.Background())
}

func (i DeviceKeyArray) ToDeviceKeyArrayOutputWithContext(ctx context.Context) DeviceKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceKeyArrayOutput)
}

// DeviceKeyMapInput is an input type that accepts DeviceKeyMap and DeviceKeyMapOutput values.
// You can construct a concrete instance of `DeviceKeyMapInput` via:
//
//	DeviceKeyMap{ "key": DeviceKeyArgs{...} }
type DeviceKeyMapInput interface {
	pulumi.Input

	ToDeviceKeyMapOutput() DeviceKeyMapOutput
	ToDeviceKeyMapOutputWithContext(context.Context) DeviceKeyMapOutput
}

type DeviceKeyMap map[string]DeviceKeyInput

func (DeviceKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceKey)(nil)).Elem()
}

func (i DeviceKeyMap) ToDeviceKeyMapOutput() DeviceKeyMapOutput {
	return i.ToDeviceKeyMapOutputWithContext(context.Background())
}

func (i DeviceKeyMap) ToDeviceKeyMapOutputWithContext(ctx context.Context) DeviceKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceKeyMapOutput)
}

type DeviceKeyOutput struct{ *pulumi.OutputState }

func (DeviceKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceKey)(nil)).Elem()
}

func (o DeviceKeyOutput) ToDeviceKeyOutput() DeviceKeyOutput {
	return o
}

func (o DeviceKeyOutput) ToDeviceKeyOutputWithContext(ctx context.Context) DeviceKeyOutput {
	return o
}

// The device to update the key properties of
func (o DeviceKeyOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceKey) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// Determines whether or not the device's key will expire. Defaults to `false`.
func (o DeviceKeyOutput) KeyExpiryDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DeviceKey) pulumi.BoolPtrOutput { return v.KeyExpiryDisabled }).(pulumi.BoolPtrOutput)
}

type DeviceKeyArrayOutput struct{ *pulumi.OutputState }

func (DeviceKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceKey)(nil)).Elem()
}

func (o DeviceKeyArrayOutput) ToDeviceKeyArrayOutput() DeviceKeyArrayOutput {
	return o
}

func (o DeviceKeyArrayOutput) ToDeviceKeyArrayOutputWithContext(ctx context.Context) DeviceKeyArrayOutput {
	return o
}

func (o DeviceKeyArrayOutput) Index(i pulumi.IntInput) DeviceKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceKey {
		return vs[0].([]*DeviceKey)[vs[1].(int)]
	}).(DeviceKeyOutput)
}

type DeviceKeyMapOutput struct{ *pulumi.OutputState }

func (DeviceKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceKey)(nil)).Elem()
}

func (o DeviceKeyMapOutput) ToDeviceKeyMapOutput() DeviceKeyMapOutput {
	return o
}

func (o DeviceKeyMapOutput) ToDeviceKeyMapOutputWithContext(ctx context.Context) DeviceKeyMapOutput {
	return o
}

func (o DeviceKeyMapOutput) MapIndex(k pulumi.StringInput) DeviceKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceKey {
		return vs[0].(map[string]*DeviceKey)[vs[1].(string)]
	}).(DeviceKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceKeyInput)(nil)).Elem(), &DeviceKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceKeyArrayInput)(nil)).Elem(), DeviceKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceKeyMapInput)(nil)).Elem(), DeviceKeyMap{})
	pulumi.RegisterOutputType(DeviceKeyOutput{})
	pulumi.RegisterOutputType(DeviceKeyArrayOutput{})
	pulumi.RegisterOutputType(DeviceKeyMapOutput{})
}
