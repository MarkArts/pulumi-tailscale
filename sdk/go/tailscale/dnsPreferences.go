// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The dnsPreferences resource allows you to configure DNS preferences for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
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
//			_, err := tailscale.NewDnsPreferences(ctx, "samplePreferences", &tailscale.DnsPreferencesArgs{
//				MagicDns: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type DnsPreferences struct {
	pulumi.CustomResourceState

	// Whether or not to enable magic DNS
	MagicDns pulumi.BoolOutput `pulumi:"magicDns"`
}

// NewDnsPreferences registers a new resource with the given unique name, arguments, and options.
func NewDnsPreferences(ctx *pulumi.Context,
	name string, args *DnsPreferencesArgs, opts ...pulumi.ResourceOption) (*DnsPreferences, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MagicDns == nil {
		return nil, errors.New("invalid value for required argument 'MagicDns'")
	}
	var resource DnsPreferences
	err := ctx.RegisterResource("tailscale:index/dnsPreferences:DnsPreferences", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsPreferences gets an existing DnsPreferences resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsPreferences(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsPreferencesState, opts ...pulumi.ResourceOption) (*DnsPreferences, error) {
	var resource DnsPreferences
	err := ctx.ReadResource("tailscale:index/dnsPreferences:DnsPreferences", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsPreferences resources.
type dnsPreferencesState struct {
	// Whether or not to enable magic DNS
	MagicDns *bool `pulumi:"magicDns"`
}

type DnsPreferencesState struct {
	// Whether or not to enable magic DNS
	MagicDns pulumi.BoolPtrInput
}

func (DnsPreferencesState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsPreferencesState)(nil)).Elem()
}

type dnsPreferencesArgs struct {
	// Whether or not to enable magic DNS
	MagicDns bool `pulumi:"magicDns"`
}

// The set of arguments for constructing a DnsPreferences resource.
type DnsPreferencesArgs struct {
	// Whether or not to enable magic DNS
	MagicDns pulumi.BoolInput
}

func (DnsPreferencesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsPreferencesArgs)(nil)).Elem()
}

type DnsPreferencesInput interface {
	pulumi.Input

	ToDnsPreferencesOutput() DnsPreferencesOutput
	ToDnsPreferencesOutputWithContext(ctx context.Context) DnsPreferencesOutput
}

func (*DnsPreferences) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsPreferences)(nil)).Elem()
}

func (i *DnsPreferences) ToDnsPreferencesOutput() DnsPreferencesOutput {
	return i.ToDnsPreferencesOutputWithContext(context.Background())
}

func (i *DnsPreferences) ToDnsPreferencesOutputWithContext(ctx context.Context) DnsPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPreferencesOutput)
}

// DnsPreferencesArrayInput is an input type that accepts DnsPreferencesArray and DnsPreferencesArrayOutput values.
// You can construct a concrete instance of `DnsPreferencesArrayInput` via:
//
//	DnsPreferencesArray{ DnsPreferencesArgs{...} }
type DnsPreferencesArrayInput interface {
	pulumi.Input

	ToDnsPreferencesArrayOutput() DnsPreferencesArrayOutput
	ToDnsPreferencesArrayOutputWithContext(context.Context) DnsPreferencesArrayOutput
}

type DnsPreferencesArray []DnsPreferencesInput

func (DnsPreferencesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsPreferences)(nil)).Elem()
}

func (i DnsPreferencesArray) ToDnsPreferencesArrayOutput() DnsPreferencesArrayOutput {
	return i.ToDnsPreferencesArrayOutputWithContext(context.Background())
}

func (i DnsPreferencesArray) ToDnsPreferencesArrayOutputWithContext(ctx context.Context) DnsPreferencesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPreferencesArrayOutput)
}

// DnsPreferencesMapInput is an input type that accepts DnsPreferencesMap and DnsPreferencesMapOutput values.
// You can construct a concrete instance of `DnsPreferencesMapInput` via:
//
//	DnsPreferencesMap{ "key": DnsPreferencesArgs{...} }
type DnsPreferencesMapInput interface {
	pulumi.Input

	ToDnsPreferencesMapOutput() DnsPreferencesMapOutput
	ToDnsPreferencesMapOutputWithContext(context.Context) DnsPreferencesMapOutput
}

type DnsPreferencesMap map[string]DnsPreferencesInput

func (DnsPreferencesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsPreferences)(nil)).Elem()
}

func (i DnsPreferencesMap) ToDnsPreferencesMapOutput() DnsPreferencesMapOutput {
	return i.ToDnsPreferencesMapOutputWithContext(context.Background())
}

func (i DnsPreferencesMap) ToDnsPreferencesMapOutputWithContext(ctx context.Context) DnsPreferencesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPreferencesMapOutput)
}

type DnsPreferencesOutput struct{ *pulumi.OutputState }

func (DnsPreferencesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsPreferences)(nil)).Elem()
}

func (o DnsPreferencesOutput) ToDnsPreferencesOutput() DnsPreferencesOutput {
	return o
}

func (o DnsPreferencesOutput) ToDnsPreferencesOutputWithContext(ctx context.Context) DnsPreferencesOutput {
	return o
}

// Whether or not to enable magic DNS
func (o DnsPreferencesOutput) MagicDns() pulumi.BoolOutput {
	return o.ApplyT(func(v *DnsPreferences) pulumi.BoolOutput { return v.MagicDns }).(pulumi.BoolOutput)
}

type DnsPreferencesArrayOutput struct{ *pulumi.OutputState }

func (DnsPreferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsPreferences)(nil)).Elem()
}

func (o DnsPreferencesArrayOutput) ToDnsPreferencesArrayOutput() DnsPreferencesArrayOutput {
	return o
}

func (o DnsPreferencesArrayOutput) ToDnsPreferencesArrayOutputWithContext(ctx context.Context) DnsPreferencesArrayOutput {
	return o
}

func (o DnsPreferencesArrayOutput) Index(i pulumi.IntInput) DnsPreferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsPreferences {
		return vs[0].([]*DnsPreferences)[vs[1].(int)]
	}).(DnsPreferencesOutput)
}

type DnsPreferencesMapOutput struct{ *pulumi.OutputState }

func (DnsPreferencesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsPreferences)(nil)).Elem()
}

func (o DnsPreferencesMapOutput) ToDnsPreferencesMapOutput() DnsPreferencesMapOutput {
	return o
}

func (o DnsPreferencesMapOutput) ToDnsPreferencesMapOutputWithContext(ctx context.Context) DnsPreferencesMapOutput {
	return o
}

func (o DnsPreferencesMapOutput) MapIndex(k pulumi.StringInput) DnsPreferencesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsPreferences {
		return vs[0].(map[string]*DnsPreferences)[vs[1].(string)]
	}).(DnsPreferencesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesInput)(nil)).Elem(), &DnsPreferences{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesArrayInput)(nil)).Elem(), DnsPreferencesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesMapInput)(nil)).Elem(), DnsPreferencesMap{})
	pulumi.RegisterOutputType(DnsPreferencesOutput{})
	pulumi.RegisterOutputType(DnsPreferencesArrayOutput{})
	pulumi.RegisterOutputType(DnsPreferencesMapOutput{})
}
