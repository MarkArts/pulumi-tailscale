// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The dnsPreferences resource allows you to configure DNS preferences for your Tailscale network. See the
// [Tailscale DNS documentation](https://tailscale.com/kb/1054/dns) for more information.
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
// 		_, err := tailscale.NewDnsPreferences(ctx, "samplePreferences", &tailscale.DnsPreferencesArgs{
// 			MagicDns: pulumi.Bool(true),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DnsPreferences struct {
	pulumi.CustomResourceState

	// Enables or disables MagicDNS, automatically registers DNS names for devices on your network.
	// At least one DNS server must be set before enabling Magic DNS.
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
	// Enables or disables MagicDNS, automatically registers DNS names for devices on your network.
	// At least one DNS server must be set before enabling Magic DNS.
	MagicDns *bool `pulumi:"magicDns"`
}

type DnsPreferencesState struct {
	// Enables or disables MagicDNS, automatically registers DNS names for devices on your network.
	// At least one DNS server must be set before enabling Magic DNS.
	MagicDns pulumi.BoolPtrInput
}

func (DnsPreferencesState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsPreferencesState)(nil)).Elem()
}

type dnsPreferencesArgs struct {
	// Enables or disables MagicDNS, automatically registers DNS names for devices on your network.
	// At least one DNS server must be set before enabling Magic DNS.
	MagicDns bool `pulumi:"magicDns"`
}

// The set of arguments for constructing a DnsPreferences resource.
type DnsPreferencesArgs struct {
	// Enables or disables MagicDNS, automatically registers DNS names for devices on your network.
	// At least one DNS server must be set before enabling Magic DNS.
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
	return reflect.TypeOf((*DnsPreferences)(nil))
}

func (i *DnsPreferences) ToDnsPreferencesOutput() DnsPreferencesOutput {
	return i.ToDnsPreferencesOutputWithContext(context.Background())
}

func (i *DnsPreferences) ToDnsPreferencesOutputWithContext(ctx context.Context) DnsPreferencesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPreferencesOutput)
}

func (i *DnsPreferences) ToDnsPreferencesPtrOutput() DnsPreferencesPtrOutput {
	return i.ToDnsPreferencesPtrOutputWithContext(context.Background())
}

func (i *DnsPreferences) ToDnsPreferencesPtrOutputWithContext(ctx context.Context) DnsPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPreferencesPtrOutput)
}

type DnsPreferencesPtrInput interface {
	pulumi.Input

	ToDnsPreferencesPtrOutput() DnsPreferencesPtrOutput
	ToDnsPreferencesPtrOutputWithContext(ctx context.Context) DnsPreferencesPtrOutput
}

type dnsPreferencesPtrType DnsPreferencesArgs

func (*dnsPreferencesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsPreferences)(nil))
}

func (i *dnsPreferencesPtrType) ToDnsPreferencesPtrOutput() DnsPreferencesPtrOutput {
	return i.ToDnsPreferencesPtrOutputWithContext(context.Background())
}

func (i *dnsPreferencesPtrType) ToDnsPreferencesPtrOutputWithContext(ctx context.Context) DnsPreferencesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsPreferencesPtrOutput)
}

// DnsPreferencesArrayInput is an input type that accepts DnsPreferencesArray and DnsPreferencesArrayOutput values.
// You can construct a concrete instance of `DnsPreferencesArrayInput` via:
//
//          DnsPreferencesArray{ DnsPreferencesArgs{...} }
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
//          DnsPreferencesMap{ "key": DnsPreferencesArgs{...} }
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
	return reflect.TypeOf((*DnsPreferences)(nil))
}

func (o DnsPreferencesOutput) ToDnsPreferencesOutput() DnsPreferencesOutput {
	return o
}

func (o DnsPreferencesOutput) ToDnsPreferencesOutputWithContext(ctx context.Context) DnsPreferencesOutput {
	return o
}

func (o DnsPreferencesOutput) ToDnsPreferencesPtrOutput() DnsPreferencesPtrOutput {
	return o.ToDnsPreferencesPtrOutputWithContext(context.Background())
}

func (o DnsPreferencesOutput) ToDnsPreferencesPtrOutputWithContext(ctx context.Context) DnsPreferencesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DnsPreferences) *DnsPreferences {
		return &v
	}).(DnsPreferencesPtrOutput)
}

type DnsPreferencesPtrOutput struct{ *pulumi.OutputState }

func (DnsPreferencesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsPreferences)(nil))
}

func (o DnsPreferencesPtrOutput) ToDnsPreferencesPtrOutput() DnsPreferencesPtrOutput {
	return o
}

func (o DnsPreferencesPtrOutput) ToDnsPreferencesPtrOutputWithContext(ctx context.Context) DnsPreferencesPtrOutput {
	return o
}

func (o DnsPreferencesPtrOutput) Elem() DnsPreferencesOutput {
	return o.ApplyT(func(v *DnsPreferences) DnsPreferences {
		if v != nil {
			return *v
		}
		var ret DnsPreferences
		return ret
	}).(DnsPreferencesOutput)
}

type DnsPreferencesArrayOutput struct{ *pulumi.OutputState }

func (DnsPreferencesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DnsPreferences)(nil))
}

func (o DnsPreferencesArrayOutput) ToDnsPreferencesArrayOutput() DnsPreferencesArrayOutput {
	return o
}

func (o DnsPreferencesArrayOutput) ToDnsPreferencesArrayOutputWithContext(ctx context.Context) DnsPreferencesArrayOutput {
	return o
}

func (o DnsPreferencesArrayOutput) Index(i pulumi.IntInput) DnsPreferencesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DnsPreferences {
		return vs[0].([]DnsPreferences)[vs[1].(int)]
	}).(DnsPreferencesOutput)
}

type DnsPreferencesMapOutput struct{ *pulumi.OutputState }

func (DnsPreferencesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DnsPreferences)(nil))
}

func (o DnsPreferencesMapOutput) ToDnsPreferencesMapOutput() DnsPreferencesMapOutput {
	return o
}

func (o DnsPreferencesMapOutput) ToDnsPreferencesMapOutputWithContext(ctx context.Context) DnsPreferencesMapOutput {
	return o
}

func (o DnsPreferencesMapOutput) MapIndex(k pulumi.StringInput) DnsPreferencesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DnsPreferences {
		return vs[0].(map[string]DnsPreferences)[vs[1].(string)]
	}).(DnsPreferencesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesInput)(nil)).Elem(), &DnsPreferences{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesPtrInput)(nil)).Elem(), &DnsPreferences{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesArrayInput)(nil)).Elem(), DnsPreferencesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsPreferencesMapInput)(nil)).Elem(), DnsPreferencesMap{})
	pulumi.RegisterOutputType(DnsPreferencesOutput{})
	pulumi.RegisterOutputType(DnsPreferencesPtrOutput{})
	pulumi.RegisterOutputType(DnsPreferencesArrayOutput{})
	pulumi.RegisterOutputType(DnsPreferencesMapOutput{})
}
