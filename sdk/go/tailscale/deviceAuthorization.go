// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The deviceAuthorization resource is used to approve new devices before they can join the tailnet. See https://tailscale.com/kb/1099/device-authorization/ for more details.
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
//			sampleDevice, err := tailscale.GetDevice(ctx, &GetDeviceArgs{
//				Name: "device.example.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = tailscale.NewDeviceAuthorization(ctx, "sampleAuthorization", &tailscale.DeviceAuthorizationArgs{
//				DeviceId:   pulumi.String(sampleDevice.Id),
//				Authorized: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DeviceAuthorization struct {
	pulumi.CustomResourceState

	// Whether or not the device is authorized
	Authorized pulumi.BoolOutput `pulumi:"authorized"`
	// The device to set as authorized
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
}

// NewDeviceAuthorization registers a new resource with the given unique name, arguments, and options.
func NewDeviceAuthorization(ctx *pulumi.Context,
	name string, args *DeviceAuthorizationArgs, opts ...pulumi.ResourceOption) (*DeviceAuthorization, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Authorized == nil {
		return nil, errors.New("invalid value for required argument 'Authorized'")
	}
	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	var resource DeviceAuthorization
	err := ctx.RegisterResource("tailscale:index/deviceAuthorization:DeviceAuthorization", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceAuthorization gets an existing DeviceAuthorization resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceAuthorization(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceAuthorizationState, opts ...pulumi.ResourceOption) (*DeviceAuthorization, error) {
	var resource DeviceAuthorization
	err := ctx.ReadResource("tailscale:index/deviceAuthorization:DeviceAuthorization", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceAuthorization resources.
type deviceAuthorizationState struct {
	// Whether or not the device is authorized
	Authorized *bool `pulumi:"authorized"`
	// The device to set as authorized
	DeviceId *string `pulumi:"deviceId"`
}

type DeviceAuthorizationState struct {
	// Whether or not the device is authorized
	Authorized pulumi.BoolPtrInput
	// The device to set as authorized
	DeviceId pulumi.StringPtrInput
}

func (DeviceAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceAuthorizationState)(nil)).Elem()
}

type deviceAuthorizationArgs struct {
	// Whether or not the device is authorized
	Authorized bool `pulumi:"authorized"`
	// The device to set as authorized
	DeviceId string `pulumi:"deviceId"`
}

// The set of arguments for constructing a DeviceAuthorization resource.
type DeviceAuthorizationArgs struct {
	// Whether or not the device is authorized
	Authorized pulumi.BoolInput
	// The device to set as authorized
	DeviceId pulumi.StringInput
}

func (DeviceAuthorizationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceAuthorizationArgs)(nil)).Elem()
}

type DeviceAuthorizationInput interface {
	pulumi.Input

	ToDeviceAuthorizationOutput() DeviceAuthorizationOutput
	ToDeviceAuthorizationOutputWithContext(ctx context.Context) DeviceAuthorizationOutput
}

func (*DeviceAuthorization) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceAuthorization)(nil)).Elem()
}

func (i *DeviceAuthorization) ToDeviceAuthorizationOutput() DeviceAuthorizationOutput {
	return i.ToDeviceAuthorizationOutputWithContext(context.Background())
}

func (i *DeviceAuthorization) ToDeviceAuthorizationOutputWithContext(ctx context.Context) DeviceAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceAuthorizationOutput)
}

// DeviceAuthorizationArrayInput is an input type that accepts DeviceAuthorizationArray and DeviceAuthorizationArrayOutput values.
// You can construct a concrete instance of `DeviceAuthorizationArrayInput` via:
//
//	DeviceAuthorizationArray{ DeviceAuthorizationArgs{...} }
type DeviceAuthorizationArrayInput interface {
	pulumi.Input

	ToDeviceAuthorizationArrayOutput() DeviceAuthorizationArrayOutput
	ToDeviceAuthorizationArrayOutputWithContext(context.Context) DeviceAuthorizationArrayOutput
}

type DeviceAuthorizationArray []DeviceAuthorizationInput

func (DeviceAuthorizationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceAuthorization)(nil)).Elem()
}

func (i DeviceAuthorizationArray) ToDeviceAuthorizationArrayOutput() DeviceAuthorizationArrayOutput {
	return i.ToDeviceAuthorizationArrayOutputWithContext(context.Background())
}

func (i DeviceAuthorizationArray) ToDeviceAuthorizationArrayOutputWithContext(ctx context.Context) DeviceAuthorizationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceAuthorizationArrayOutput)
}

// DeviceAuthorizationMapInput is an input type that accepts DeviceAuthorizationMap and DeviceAuthorizationMapOutput values.
// You can construct a concrete instance of `DeviceAuthorizationMapInput` via:
//
//	DeviceAuthorizationMap{ "key": DeviceAuthorizationArgs{...} }
type DeviceAuthorizationMapInput interface {
	pulumi.Input

	ToDeviceAuthorizationMapOutput() DeviceAuthorizationMapOutput
	ToDeviceAuthorizationMapOutputWithContext(context.Context) DeviceAuthorizationMapOutput
}

type DeviceAuthorizationMap map[string]DeviceAuthorizationInput

func (DeviceAuthorizationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceAuthorization)(nil)).Elem()
}

func (i DeviceAuthorizationMap) ToDeviceAuthorizationMapOutput() DeviceAuthorizationMapOutput {
	return i.ToDeviceAuthorizationMapOutputWithContext(context.Background())
}

func (i DeviceAuthorizationMap) ToDeviceAuthorizationMapOutputWithContext(ctx context.Context) DeviceAuthorizationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceAuthorizationMapOutput)
}

type DeviceAuthorizationOutput struct{ *pulumi.OutputState }

func (DeviceAuthorizationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceAuthorization)(nil)).Elem()
}

func (o DeviceAuthorizationOutput) ToDeviceAuthorizationOutput() DeviceAuthorizationOutput {
	return o
}

func (o DeviceAuthorizationOutput) ToDeviceAuthorizationOutputWithContext(ctx context.Context) DeviceAuthorizationOutput {
	return o
}

// Whether or not the device is authorized
func (o DeviceAuthorizationOutput) Authorized() pulumi.BoolOutput {
	return o.ApplyT(func(v *DeviceAuthorization) pulumi.BoolOutput { return v.Authorized }).(pulumi.BoolOutput)
}

// The device to set as authorized
func (o DeviceAuthorizationOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceAuthorization) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

type DeviceAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (DeviceAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceAuthorization)(nil)).Elem()
}

func (o DeviceAuthorizationArrayOutput) ToDeviceAuthorizationArrayOutput() DeviceAuthorizationArrayOutput {
	return o
}

func (o DeviceAuthorizationArrayOutput) ToDeviceAuthorizationArrayOutputWithContext(ctx context.Context) DeviceAuthorizationArrayOutput {
	return o
}

func (o DeviceAuthorizationArrayOutput) Index(i pulumi.IntInput) DeviceAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceAuthorization {
		return vs[0].([]*DeviceAuthorization)[vs[1].(int)]
	}).(DeviceAuthorizationOutput)
}

type DeviceAuthorizationMapOutput struct{ *pulumi.OutputState }

func (DeviceAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceAuthorization)(nil)).Elem()
}

func (o DeviceAuthorizationMapOutput) ToDeviceAuthorizationMapOutput() DeviceAuthorizationMapOutput {
	return o
}

func (o DeviceAuthorizationMapOutput) ToDeviceAuthorizationMapOutputWithContext(ctx context.Context) DeviceAuthorizationMapOutput {
	return o
}

func (o DeviceAuthorizationMapOutput) MapIndex(k pulumi.StringInput) DeviceAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceAuthorization {
		return vs[0].(map[string]*DeviceAuthorization)[vs[1].(string)]
	}).(DeviceAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationInput)(nil)).Elem(), &DeviceAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationArrayInput)(nil)).Elem(), DeviceAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationMapInput)(nil)).Elem(), DeviceAuthorizationMap{})
	pulumi.RegisterOutputType(DeviceAuthorizationOutput{})
	pulumi.RegisterOutputType(DeviceAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(DeviceAuthorizationMapOutput{})
}
