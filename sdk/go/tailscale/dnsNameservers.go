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

// The dnsNameservers resource allows you to configure DNS nameservers for your Tailscale network. See https://tailscale.com/kb/1054/dns for more information.
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
//			_, err := tailscale.NewDnsNameservers(ctx, "sampleNameservers", &tailscale.DnsNameserversArgs{
//				Nameservers: pulumi.StringArray{
//					pulumi.String("8.8.8.8"),
//					pulumi.String("8.8.4.4"),
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
type DnsNameservers struct {
	pulumi.CustomResourceState

	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers pulumi.StringArrayOutput `pulumi:"nameservers"`
}

// NewDnsNameservers registers a new resource with the given unique name, arguments, and options.
func NewDnsNameservers(ctx *pulumi.Context,
	name string, args *DnsNameserversArgs, opts ...pulumi.ResourceOption) (*DnsNameservers, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Nameservers == nil {
		return nil, errors.New("invalid value for required argument 'Nameservers'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DnsNameservers
	err := ctx.RegisterResource("tailscale:index/dnsNameservers:DnsNameservers", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsNameservers gets an existing DnsNameservers resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsNameservers(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsNameserversState, opts ...pulumi.ResourceOption) (*DnsNameservers, error) {
	var resource DnsNameservers
	err := ctx.ReadResource("tailscale:index/dnsNameservers:DnsNameservers", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsNameservers resources.
type dnsNameserversState struct {
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers []string `pulumi:"nameservers"`
}

type DnsNameserversState struct {
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers pulumi.StringArrayInput
}

func (DnsNameserversState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsNameserversState)(nil)).Elem()
}

type dnsNameserversArgs struct {
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers []string `pulumi:"nameservers"`
}

// The set of arguments for constructing a DnsNameservers resource.
type DnsNameserversArgs struct {
	// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
	Nameservers pulumi.StringArrayInput
}

func (DnsNameserversArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsNameserversArgs)(nil)).Elem()
}

type DnsNameserversInput interface {
	pulumi.Input

	ToDnsNameserversOutput() DnsNameserversOutput
	ToDnsNameserversOutputWithContext(ctx context.Context) DnsNameserversOutput
}

func (*DnsNameservers) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsNameservers)(nil)).Elem()
}

func (i *DnsNameservers) ToDnsNameserversOutput() DnsNameserversOutput {
	return i.ToDnsNameserversOutputWithContext(context.Background())
}

func (i *DnsNameservers) ToDnsNameserversOutputWithContext(ctx context.Context) DnsNameserversOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsNameserversOutput)
}

// DnsNameserversArrayInput is an input type that accepts DnsNameserversArray and DnsNameserversArrayOutput values.
// You can construct a concrete instance of `DnsNameserversArrayInput` via:
//
//	DnsNameserversArray{ DnsNameserversArgs{...} }
type DnsNameserversArrayInput interface {
	pulumi.Input

	ToDnsNameserversArrayOutput() DnsNameserversArrayOutput
	ToDnsNameserversArrayOutputWithContext(context.Context) DnsNameserversArrayOutput
}

type DnsNameserversArray []DnsNameserversInput

func (DnsNameserversArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsNameservers)(nil)).Elem()
}

func (i DnsNameserversArray) ToDnsNameserversArrayOutput() DnsNameserversArrayOutput {
	return i.ToDnsNameserversArrayOutputWithContext(context.Background())
}

func (i DnsNameserversArray) ToDnsNameserversArrayOutputWithContext(ctx context.Context) DnsNameserversArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsNameserversArrayOutput)
}

// DnsNameserversMapInput is an input type that accepts DnsNameserversMap and DnsNameserversMapOutput values.
// You can construct a concrete instance of `DnsNameserversMapInput` via:
//
//	DnsNameserversMap{ "key": DnsNameserversArgs{...} }
type DnsNameserversMapInput interface {
	pulumi.Input

	ToDnsNameserversMapOutput() DnsNameserversMapOutput
	ToDnsNameserversMapOutputWithContext(context.Context) DnsNameserversMapOutput
}

type DnsNameserversMap map[string]DnsNameserversInput

func (DnsNameserversMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsNameservers)(nil)).Elem()
}

func (i DnsNameserversMap) ToDnsNameserversMapOutput() DnsNameserversMapOutput {
	return i.ToDnsNameserversMapOutputWithContext(context.Background())
}

func (i DnsNameserversMap) ToDnsNameserversMapOutputWithContext(ctx context.Context) DnsNameserversMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsNameserversMapOutput)
}

type DnsNameserversOutput struct{ *pulumi.OutputState }

func (DnsNameserversOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsNameservers)(nil)).Elem()
}

func (o DnsNameserversOutput) ToDnsNameserversOutput() DnsNameserversOutput {
	return o
}

func (o DnsNameserversOutput) ToDnsNameserversOutputWithContext(ctx context.Context) DnsNameserversOutput {
	return o
}

// Devices on your network will use these nameservers to resolve DNS names. IPv4 or IPv6 addresses are accepted.
func (o DnsNameserversOutput) Nameservers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DnsNameservers) pulumi.StringArrayOutput { return v.Nameservers }).(pulumi.StringArrayOutput)
}

type DnsNameserversArrayOutput struct{ *pulumi.OutputState }

func (DnsNameserversArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsNameservers)(nil)).Elem()
}

func (o DnsNameserversArrayOutput) ToDnsNameserversArrayOutput() DnsNameserversArrayOutput {
	return o
}

func (o DnsNameserversArrayOutput) ToDnsNameserversArrayOutputWithContext(ctx context.Context) DnsNameserversArrayOutput {
	return o
}

func (o DnsNameserversArrayOutput) Index(i pulumi.IntInput) DnsNameserversOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsNameservers {
		return vs[0].([]*DnsNameservers)[vs[1].(int)]
	}).(DnsNameserversOutput)
}

type DnsNameserversMapOutput struct{ *pulumi.OutputState }

func (DnsNameserversMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsNameservers)(nil)).Elem()
}

func (o DnsNameserversMapOutput) ToDnsNameserversMapOutput() DnsNameserversMapOutput {
	return o
}

func (o DnsNameserversMapOutput) ToDnsNameserversMapOutputWithContext(ctx context.Context) DnsNameserversMapOutput {
	return o
}

func (o DnsNameserversMapOutput) MapIndex(k pulumi.StringInput) DnsNameserversOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsNameservers {
		return vs[0].(map[string]*DnsNameservers)[vs[1].(string)]
	}).(DnsNameserversOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsNameserversInput)(nil)).Elem(), &DnsNameservers{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsNameserversArrayInput)(nil)).Elem(), DnsNameserversArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsNameserversMapInput)(nil)).Elem(), DnsNameserversMap{})
	pulumi.RegisterOutputType(DnsNameserversOutput{})
	pulumi.RegisterOutputType(DnsNameserversArrayOutput{})
	pulumi.RegisterOutputType(DnsNameserversMapOutput{})
}
