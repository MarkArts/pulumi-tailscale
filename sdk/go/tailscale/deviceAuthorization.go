// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The deviceAuthorization resource is used to approve new devices before they can join the tailnet.
// See the [Tailscale device authorization documentation](https://tailscale.com/kb/1099/device-authorization) for more
// information.
//
// The Tailscale API currently only supports authorizing devices, but not rejecting/removing them. Once a device is
// authorized by this provider it cannot be modified again afterwards. Modifying or deleting the resource
// will not affect the device's authorization within the tailnet.
type DeviceAuthorization struct {
	pulumi.CustomResourceState

	// Indicates if the device is allowed to join the tailnet.
	Authorized pulumi.BoolOutput `pulumi:"authorized"`
	// The device to authorize.
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
	// Indicates if the device is allowed to join the tailnet.
	Authorized *bool `pulumi:"authorized"`
	// The device to authorize.
	DeviceId *string `pulumi:"deviceId"`
}

type DeviceAuthorizationState struct {
	// Indicates if the device is allowed to join the tailnet.
	Authorized pulumi.BoolPtrInput
	// The device to authorize.
	DeviceId pulumi.StringPtrInput
}

func (DeviceAuthorizationState) ElementType() reflect.Type {
	return reflect.TypeOf((*deviceAuthorizationState)(nil)).Elem()
}

type deviceAuthorizationArgs struct {
	// Indicates if the device is allowed to join the tailnet.
	Authorized bool `pulumi:"authorized"`
	// The device to authorize.
	DeviceId string `pulumi:"deviceId"`
}

// The set of arguments for constructing a DeviceAuthorization resource.
type DeviceAuthorizationArgs struct {
	// Indicates if the device is allowed to join the tailnet.
	Authorized pulumi.BoolInput
	// The device to authorize.
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
	return reflect.TypeOf((*DeviceAuthorization)(nil))
}

func (i *DeviceAuthorization) ToDeviceAuthorizationOutput() DeviceAuthorizationOutput {
	return i.ToDeviceAuthorizationOutputWithContext(context.Background())
}

func (i *DeviceAuthorization) ToDeviceAuthorizationOutputWithContext(ctx context.Context) DeviceAuthorizationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceAuthorizationOutput)
}

func (i *DeviceAuthorization) ToDeviceAuthorizationPtrOutput() DeviceAuthorizationPtrOutput {
	return i.ToDeviceAuthorizationPtrOutputWithContext(context.Background())
}

func (i *DeviceAuthorization) ToDeviceAuthorizationPtrOutputWithContext(ctx context.Context) DeviceAuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceAuthorizationPtrOutput)
}

type DeviceAuthorizationPtrInput interface {
	pulumi.Input

	ToDeviceAuthorizationPtrOutput() DeviceAuthorizationPtrOutput
	ToDeviceAuthorizationPtrOutputWithContext(ctx context.Context) DeviceAuthorizationPtrOutput
}

type deviceAuthorizationPtrType DeviceAuthorizationArgs

func (*deviceAuthorizationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceAuthorization)(nil))
}

func (i *deviceAuthorizationPtrType) ToDeviceAuthorizationPtrOutput() DeviceAuthorizationPtrOutput {
	return i.ToDeviceAuthorizationPtrOutputWithContext(context.Background())
}

func (i *deviceAuthorizationPtrType) ToDeviceAuthorizationPtrOutputWithContext(ctx context.Context) DeviceAuthorizationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeviceAuthorizationPtrOutput)
}

// DeviceAuthorizationArrayInput is an input type that accepts DeviceAuthorizationArray and DeviceAuthorizationArrayOutput values.
// You can construct a concrete instance of `DeviceAuthorizationArrayInput` via:
//
//          DeviceAuthorizationArray{ DeviceAuthorizationArgs{...} }
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
//          DeviceAuthorizationMap{ "key": DeviceAuthorizationArgs{...} }
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
	return reflect.TypeOf((*DeviceAuthorization)(nil))
}

func (o DeviceAuthorizationOutput) ToDeviceAuthorizationOutput() DeviceAuthorizationOutput {
	return o
}

func (o DeviceAuthorizationOutput) ToDeviceAuthorizationOutputWithContext(ctx context.Context) DeviceAuthorizationOutput {
	return o
}

func (o DeviceAuthorizationOutput) ToDeviceAuthorizationPtrOutput() DeviceAuthorizationPtrOutput {
	return o.ToDeviceAuthorizationPtrOutputWithContext(context.Background())
}

func (o DeviceAuthorizationOutput) ToDeviceAuthorizationPtrOutputWithContext(ctx context.Context) DeviceAuthorizationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeviceAuthorization) *DeviceAuthorization {
		return &v
	}).(DeviceAuthorizationPtrOutput)
}

type DeviceAuthorizationPtrOutput struct{ *pulumi.OutputState }

func (DeviceAuthorizationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeviceAuthorization)(nil))
}

func (o DeviceAuthorizationPtrOutput) ToDeviceAuthorizationPtrOutput() DeviceAuthorizationPtrOutput {
	return o
}

func (o DeviceAuthorizationPtrOutput) ToDeviceAuthorizationPtrOutputWithContext(ctx context.Context) DeviceAuthorizationPtrOutput {
	return o
}

func (o DeviceAuthorizationPtrOutput) Elem() DeviceAuthorizationOutput {
	return o.ApplyT(func(v *DeviceAuthorization) DeviceAuthorization {
		if v != nil {
			return *v
		}
		var ret DeviceAuthorization
		return ret
	}).(DeviceAuthorizationOutput)
}

type DeviceAuthorizationArrayOutput struct{ *pulumi.OutputState }

func (DeviceAuthorizationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DeviceAuthorization)(nil))
}

func (o DeviceAuthorizationArrayOutput) ToDeviceAuthorizationArrayOutput() DeviceAuthorizationArrayOutput {
	return o
}

func (o DeviceAuthorizationArrayOutput) ToDeviceAuthorizationArrayOutputWithContext(ctx context.Context) DeviceAuthorizationArrayOutput {
	return o
}

func (o DeviceAuthorizationArrayOutput) Index(i pulumi.IntInput) DeviceAuthorizationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DeviceAuthorization {
		return vs[0].([]DeviceAuthorization)[vs[1].(int)]
	}).(DeviceAuthorizationOutput)
}

type DeviceAuthorizationMapOutput struct{ *pulumi.OutputState }

func (DeviceAuthorizationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DeviceAuthorization)(nil))
}

func (o DeviceAuthorizationMapOutput) ToDeviceAuthorizationMapOutput() DeviceAuthorizationMapOutput {
	return o
}

func (o DeviceAuthorizationMapOutput) ToDeviceAuthorizationMapOutputWithContext(ctx context.Context) DeviceAuthorizationMapOutput {
	return o
}

func (o DeviceAuthorizationMapOutput) MapIndex(k pulumi.StringInput) DeviceAuthorizationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DeviceAuthorization {
		return vs[0].(map[string]DeviceAuthorization)[vs[1].(string)]
	}).(DeviceAuthorizationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationInput)(nil)).Elem(), &DeviceAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationPtrInput)(nil)).Elem(), &DeviceAuthorization{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationArrayInput)(nil)).Elem(), DeviceAuthorizationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DeviceAuthorizationMapInput)(nil)).Elem(), DeviceAuthorizationMap{})
	pulumi.RegisterOutputType(DeviceAuthorizationOutput{})
	pulumi.RegisterOutputType(DeviceAuthorizationPtrOutput{})
	pulumi.RegisterOutputType(DeviceAuthorizationArrayOutput{})
	pulumi.RegisterOutputType(DeviceAuthorizationMapOutput{})
}
