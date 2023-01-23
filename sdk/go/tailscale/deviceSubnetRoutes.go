// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The deviceSubnetRoutes resource allows you to configure subnet routes for your Tailscale devices. See https://tailscale.com/kb/1019/subnets for more information.
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
//			sampleDevice, err := tailscale.GetDevice(ctx, &tailscale.GetDeviceArgs{
//				Name: "device.example.com",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = tailscale.NewDeviceSubnetRoutes(ctx, "sampleRoutes", &tailscale.DeviceSubnetRoutesArgs{
//				DeviceId: *pulumi.String(sampleDevice.Id),
//				Routes: pulumi.StringArray{
//					pulumi.String("10.0.1.0/24"),
//					pulumi.String("1.2.0.0/16"),
//					pulumi.String("2.0.0.0/24"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DeviceSubnetRoutes struct {
	pulumi.CustomResourceState

	// The device to set subnet routes for
	DeviceId pulumi.StringOutput `pulumi:"deviceId"`
	// The subnet routes that are enabled to be routed by a device
	Routes pulumi.StringArrayOutput `pulumi:"routes"`
}

// NewDeviceSubnetRoutes registers a new resource with the given unique name, arguments, and options.
func NewDeviceSubnetRoutes(ctx *pulumi.Context,
	name string, args *DeviceSubnetRoutesArgs, opts ...pulumi.ResourceOption) (*DeviceSubnetRoutes, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeviceId == nil {
		return nil, errors.New("invalid value for required argument 'DeviceId'")
	}
	if args.Routes == nil {
		return nil, errors.New("invalid value for required argument 'Routes'")
	}
	var resource DeviceSubnetRoutes
	err := ctx.RegisterResource("tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDeviceSubnetRoutes gets an existing DeviceSubnetRoutes resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDeviceSubnetRoutes(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DeviceSubnetRoutesState, opts ...pulumi.ResourceOption) (*DeviceSubnetRoutes, error) {
	var resource DeviceSubnetRoutes
	err := ctx.ReadResource("tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DeviceSubnetRoutes resources.
type deviceSubnetRoutesState struct {
	// The device to set subnet routes for
	DeviceId *string `pulumi:"deviceId"`
	// The subnet routes that are enabled to be routed by a device
	Routes []string `pulumi:"routes"`
}

type DeviceSubnetRoutesState struct {
	// The device to set subnet routes for
	DeviceId pulumi.StringPtrInput
	// The subnet routes that are enabled to be routed by a device
	Routes pulumi.StringArrayInput
}

func (DeviceSubnetRoutesState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceSubnetRoutesState)(nil)).Elem()
}

type deviceSubnetRoutesArgs struct {
	// The device to set subnet routes for
	DeviceId string `pulumi:"deviceId"`
	// The subnet routes that are enabled to be routed by a device
	Routes []string `pulumi:"routes"`
}

// The set of arguments for constructing a DeviceSubnetRoutes resource.
type DeviceSubnetRoutesArgs struct {
	// The device to set subnet routes for
	DeviceId pulumi.StringInput
	// The subnet routes that are enabled to be routed by a device
	Routes pulumi.StringArrayInput
}

func (DeviceSubnetRoutesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceSubnetRoutesArgs)(nil)).Elem()
}

type DeviceSubnetRoutesInput interface {
	pulumi.Input

	ToDeviceSubnetRoutesOutput() DeviceSubnetRoutesOutput
	ToDeviceSubnetRoutesOutputWithContext(ctx context.Context) DeviceSubnetRoutesOutput
}

func (*DeviceSubnetRoutes) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceSubnetRoutes)(nil)).Elem()
}

func (i *DeviceSubnetRoutes) ToDeviceSubnetRoutesOutput() DeviceSubnetRoutesOutput {
	return i.ToDeviceSubnetRoutesOutputWithContext(context.Background())
}

func (i *DeviceSubnetRoutes) ToDeviceSubnetRoutesOutputWithContext(ctx context.Context) DeviceSubnetRoutesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceSubnetRoutesOutput)
}

// DeviceSubnetRoutesArrayInput is an input type that accepts DeviceSubnetRoutesArray and DeviceSubnetRoutesArrayOutput values.
// You can construct a concrete instance of `DeviceSubnetRoutesArrayInput` via:
//
//	DeviceSubnetRoutesArray{ DeviceSubnetRoutesArgs{...} }
type DeviceSubnetRoutesArrayInput interface {
	pulumi.Input

	ToDeviceSubnetRoutesArrayOutput() DeviceSubnetRoutesArrayOutput
	ToDeviceSubnetRoutesArrayOutputWithContext(context.Context) DeviceSubnetRoutesArrayOutput
}

type DeviceSubnetRoutesArray []DeviceSubnetRoutesInput

func (DeviceSubnetRoutesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceSubnetRoutes)(nil)).Elem()
}

func (i DeviceSubnetRoutesArray) ToDeviceSubnetRoutesArrayOutput() DeviceSubnetRoutesArrayOutput {
	return i.ToDeviceSubnetRoutesArrayOutputWithContext(context.Background())
}

func (i DeviceSubnetRoutesArray) ToDeviceSubnetRoutesArrayOutputWithContext(ctx context.Context) DeviceSubnetRoutesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceSubnetRoutesArrayOutput)
}

// DeviceSubnetRoutesMapInput is an input type that accepts DeviceSubnetRoutesMap and DeviceSubnetRoutesMapOutput values.
// You can construct a concrete instance of `DeviceSubnetRoutesMapInput` via:
//
//	DeviceSubnetRoutesMap{ "key": DeviceSubnetRoutesArgs{...} }
type DeviceSubnetRoutesMapInput interface {
	pulumi.Input

	ToDeviceSubnetRoutesMapOutput() DeviceSubnetRoutesMapOutput
	ToDeviceSubnetRoutesMapOutputWithContext(context.Context) DeviceSubnetRoutesMapOutput
}

type DeviceSubnetRoutesMap map[string]DeviceSubnetRoutesInput

func (DeviceSubnetRoutesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceSubnetRoutes)(nil)).Elem()
}

func (i DeviceSubnetRoutesMap) ToDeviceSubnetRoutesMapOutput() DeviceSubnetRoutesMapOutput {
	return i.ToDeviceSubnetRoutesMapOutputWithContext(context.Background())
}

func (i DeviceSubnetRoutesMap) ToDeviceSubnetRoutesMapOutputWithContext(ctx context.Context) DeviceSubnetRoutesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceSubnetRoutesMapOutput)
}

type DeviceSubnetRoutesOutput struct{ *pulumi.OutputState }

func (DeviceSubnetRoutesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceSubnetRoutes)(nil)).Elem()
}

func (o DeviceSubnetRoutesOutput) ToDeviceSubnetRoutesOutput() DeviceSubnetRoutesOutput {
	return o
}

func (o DeviceSubnetRoutesOutput) ToDeviceSubnetRoutesOutputWithContext(ctx context.Context) DeviceSubnetRoutesOutput {
	return o
}

// The device to set subnet routes for
func (o DeviceSubnetRoutesOutput) DeviceId() pulumi.StringOutput {
	return o.ApplyT(func(v *DeviceSubnetRoutes) pulumi.StringOutput { return v.DeviceId }).(pulumi.StringOutput)
}

// The subnet routes that are enabled to be routed by a device
func (o DeviceSubnetRoutesOutput) Routes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DeviceSubnetRoutes) pulumi.StringArrayOutput { return v.Routes }).(pulumi.StringArrayOutput)
}

type DeviceSubnetRoutesArrayOutput struct{ *pulumi.OutputState }

func (DeviceSubnetRoutesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DeviceSubnetRoutes)(nil)).Elem()
}

func (o DeviceSubnetRoutesArrayOutput) ToDeviceSubnetRoutesArrayOutput() DeviceSubnetRoutesArrayOutput {
	return o
}

func (o DeviceSubnetRoutesArrayOutput) ToDeviceSubnetRoutesArrayOutputWithContext(ctx context.Context) DeviceSubnetRoutesArrayOutput {
	return o
}

func (o DeviceSubnetRoutesArrayOutput) Index(i pulumi.IntInput) DeviceSubnetRoutesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DeviceSubnetRoutes {
		return vs[0].([]*DeviceSubnetRoutes)[vs[1].(int)]
	}).(DeviceSubnetRoutesOutput)
}

type DeviceSubnetRoutesMapOutput struct{ *pulumi.OutputState }

func (DeviceSubnetRoutesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DeviceSubnetRoutes)(nil)).Elem()
}

func (o DeviceSubnetRoutesMapOutput) ToDeviceSubnetRoutesMapOutput() DeviceSubnetRoutesMapOutput {
	return o
}

func (o DeviceSubnetRoutesMapOutput) ToDeviceSubnetRoutesMapOutputWithContext(ctx context.Context) DeviceSubnetRoutesMapOutput {
	return o
}

func (o DeviceSubnetRoutesMapOutput) MapIndex(k pulumi.StringInput) DeviceSubnetRoutesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DeviceSubnetRoutes {
		return vs[0].(map[string]*DeviceSubnetRoutes)[vs[1].(string)]
	}).(DeviceSubnetRoutesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceSubnetRoutesInput)(nil)).Elem(), &DeviceSubnetRoutes{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceSubnetRoutesArrayInput)(nil)).Elem(), DeviceSubnetRoutesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceSubnetRoutesMapInput)(nil)).Elem(), DeviceSubnetRoutesMap{})
	pulumi.RegisterOutputType(DeviceSubnetRoutesOutput{})
	pulumi.RegisterOutputType(DeviceSubnetRoutesArrayOutput{})
	pulumi.RegisterOutputType(DeviceSubnetRoutesMapOutput{})
}
