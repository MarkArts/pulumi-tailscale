// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
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
// 	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := tailscale.NewDnsSearchPaths(ctx, "sampleSearchPaths", &tailscale.DnsSearchPathsArgs{
// 			SearchPaths: pulumi.StringArray{
// 				pulumi.String("example.com"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type DnsSearchPaths struct {
	pulumi.CustomResourceState

	// Devices on your network will use these domain suffixes to resolve DNS names.
	SearchPaths pulumi.StringArrayOutput `pulumi:"searchPaths"`
}

// NewDnsSearchPaths registers a new resource with the given unique name, arguments, and options.
func NewDnsSearchPaths(ctx *pulumi.Context,
	name string, args *DnsSearchPathsArgs, opts ...pulumi.ResourceOption) (*DnsSearchPaths, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SearchPaths == nil {
		return nil, errors.New("invalid value for required argument 'SearchPaths'")
	}
	var resource DnsSearchPaths
	err := ctx.RegisterResource("tailscale:index/dnsSearchPaths:DnsSearchPaths", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDnsSearchPaths gets an existing DnsSearchPaths resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDnsSearchPaths(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DnsSearchPathsState, opts ...pulumi.ResourceOption) (*DnsSearchPaths, error) {
	var resource DnsSearchPaths
	err := ctx.ReadResource("tailscale:index/dnsSearchPaths:DnsSearchPaths", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DnsSearchPaths resources.
type dnsSearchPathsState struct {
	// Devices on your network will use these domain suffixes to resolve DNS names.
	SearchPaths []string `pulumi:"searchPaths"`
}

type DnsSearchPathsState struct {
	// Devices on your network will use these domain suffixes to resolve DNS names.
	SearchPaths pulumi.StringArrayInput
}

func (DnsSearchPathsState) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsSearchPathsState)(nil)).Elem()
}

type dnsSearchPathsArgs struct {
	// Devices on your network will use these domain suffixes to resolve DNS names.
	SearchPaths []string `pulumi:"searchPaths"`
}

// The set of arguments for constructing a DnsSearchPaths resource.
type DnsSearchPathsArgs struct {
	// Devices on your network will use these domain suffixes to resolve DNS names.
	SearchPaths pulumi.StringArrayInput
}

func (DnsSearchPathsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dnsSearchPathsArgs)(nil)).Elem()
}

type DnsSearchPathsInput interface {
	pulumi.Input

	ToDnsSearchPathsOutput() DnsSearchPathsOutput
	ToDnsSearchPathsOutputWithContext(ctx context.Context) DnsSearchPathsOutput
}

func (*DnsSearchPaths) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSearchPaths)(nil)).Elem()
}

func (i *DnsSearchPaths) ToDnsSearchPathsOutput() DnsSearchPathsOutput {
	return i.ToDnsSearchPathsOutputWithContext(context.Background())
}

func (i *DnsSearchPaths) ToDnsSearchPathsOutputWithContext(ctx context.Context) DnsSearchPathsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSearchPathsOutput)
}

// DnsSearchPathsArrayInput is an input type that accepts DnsSearchPathsArray and DnsSearchPathsArrayOutput values.
// You can construct a concrete instance of `DnsSearchPathsArrayInput` via:
//
//          DnsSearchPathsArray{ DnsSearchPathsArgs{...} }
type DnsSearchPathsArrayInput interface {
	pulumi.Input

	ToDnsSearchPathsArrayOutput() DnsSearchPathsArrayOutput
	ToDnsSearchPathsArrayOutputWithContext(context.Context) DnsSearchPathsArrayOutput
}

type DnsSearchPathsArray []DnsSearchPathsInput

func (DnsSearchPathsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsSearchPaths)(nil)).Elem()
}

func (i DnsSearchPathsArray) ToDnsSearchPathsArrayOutput() DnsSearchPathsArrayOutput {
	return i.ToDnsSearchPathsArrayOutputWithContext(context.Background())
}

func (i DnsSearchPathsArray) ToDnsSearchPathsArrayOutputWithContext(ctx context.Context) DnsSearchPathsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSearchPathsArrayOutput)
}

// DnsSearchPathsMapInput is an input type that accepts DnsSearchPathsMap and DnsSearchPathsMapOutput values.
// You can construct a concrete instance of `DnsSearchPathsMapInput` via:
//
//          DnsSearchPathsMap{ "key": DnsSearchPathsArgs{...} }
type DnsSearchPathsMapInput interface {
	pulumi.Input

	ToDnsSearchPathsMapOutput() DnsSearchPathsMapOutput
	ToDnsSearchPathsMapOutputWithContext(context.Context) DnsSearchPathsMapOutput
}

type DnsSearchPathsMap map[string]DnsSearchPathsInput

func (DnsSearchPathsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsSearchPaths)(nil)).Elem()
}

func (i DnsSearchPathsMap) ToDnsSearchPathsMapOutput() DnsSearchPathsMapOutput {
	return i.ToDnsSearchPathsMapOutputWithContext(context.Background())
}

func (i DnsSearchPathsMap) ToDnsSearchPathsMapOutputWithContext(ctx context.Context) DnsSearchPathsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DnsSearchPathsMapOutput)
}

type DnsSearchPathsOutput struct{ *pulumi.OutputState }

func (DnsSearchPathsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DnsSearchPaths)(nil)).Elem()
}

func (o DnsSearchPathsOutput) ToDnsSearchPathsOutput() DnsSearchPathsOutput {
	return o
}

func (o DnsSearchPathsOutput) ToDnsSearchPathsOutputWithContext(ctx context.Context) DnsSearchPathsOutput {
	return o
}

type DnsSearchPathsArrayOutput struct{ *pulumi.OutputState }

func (DnsSearchPathsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DnsSearchPaths)(nil)).Elem()
}

func (o DnsSearchPathsArrayOutput) ToDnsSearchPathsArrayOutput() DnsSearchPathsArrayOutput {
	return o
}

func (o DnsSearchPathsArrayOutput) ToDnsSearchPathsArrayOutputWithContext(ctx context.Context) DnsSearchPathsArrayOutput {
	return o
}

func (o DnsSearchPathsArrayOutput) Index(i pulumi.IntInput) DnsSearchPathsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DnsSearchPaths {
		return vs[0].([]*DnsSearchPaths)[vs[1].(int)]
	}).(DnsSearchPathsOutput)
}

type DnsSearchPathsMapOutput struct{ *pulumi.OutputState }

func (DnsSearchPathsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DnsSearchPaths)(nil)).Elem()
}

func (o DnsSearchPathsMapOutput) ToDnsSearchPathsMapOutput() DnsSearchPathsMapOutput {
	return o
}

func (o DnsSearchPathsMapOutput) ToDnsSearchPathsMapOutputWithContext(ctx context.Context) DnsSearchPathsMapOutput {
	return o
}

func (o DnsSearchPathsMapOutput) MapIndex(k pulumi.StringInput) DnsSearchPathsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DnsSearchPaths {
		return vs[0].(map[string]*DnsSearchPaths)[vs[1].(string)]
	}).(DnsSearchPathsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSearchPathsInput)(nil)).Elem(), &DnsSearchPaths{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSearchPathsArrayInput)(nil)).Elem(), DnsSearchPathsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DnsSearchPathsMapInput)(nil)).Elem(), DnsSearchPathsMap{})
	pulumi.RegisterOutputType(DnsSearchPathsOutput{})
	pulumi.RegisterOutputType(DnsSearchPathsArrayOutput{})
	pulumi.RegisterOutputType(DnsSearchPathsMapOutput{})
}
