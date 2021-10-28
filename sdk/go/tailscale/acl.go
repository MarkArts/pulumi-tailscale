// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The acl resource allows you to configure a Tailscale ACL. See the [Tailscale ACL documentation](https://tailscale.com/kb/1018/acls)
// for more information. You can set the ACL in multiple ways including hujson.
type Acl struct {
	pulumi.CustomResourceState

	// The JSON-based policy that defines which devices and users are allowed to connect in your network.
	// This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
	// local ACL file.
	Acl pulumi.StringOutput `pulumi:"acl"`
}

// NewAcl registers a new resource with the given unique name, arguments, and options.
func NewAcl(ctx *pulumi.Context,
	name string, args *AclArgs, opts ...pulumi.ResourceOption) (*Acl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Acl == nil {
		return nil, errors.New("invalid value for required argument 'Acl'")
	}
	var resource Acl
	err := ctx.RegisterResource("tailscale:index/acl:Acl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAcl gets an existing Acl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AclState, opts ...pulumi.ResourceOption) (*Acl, error) {
	var resource Acl
	err := ctx.ReadResource("tailscale:index/acl:Acl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Acl resources.
type aclState struct {
	// The JSON-based policy that defines which devices and users are allowed to connect in your network.
	// This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
	// local ACL file.
	Acl *string `pulumi:"acl"`
}

type AclState struct {
	// The JSON-based policy that defines which devices and users are allowed to connect in your network.
	// This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
	// local ACL file.
	Acl pulumi.StringPtrInput
}

func (AclState) ElementType() reflect.Type {
	return reflect.TypeOf((*aclState)(nil)).Elem()
}

type aclArgs struct {
	// The JSON-based policy that defines which devices and users are allowed to connect in your network.
	// This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
	// local ACL file.
	Acl string `pulumi:"acl"`
}

// The set of arguments for constructing a Acl resource.
type AclArgs struct {
	// The JSON-based policy that defines which devices and users are allowed to connect in your network.
	// This can be JSON or HuJSON. Comments will not be provided when sent to the Tailscale API, they were only appear in your
	// local ACL file.
	Acl pulumi.StringInput
}

func (AclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*aclArgs)(nil)).Elem()
}

type AclInput interface {
	pulumi.Input

	ToAclOutput() AclOutput
	ToAclOutputWithContext(ctx context.Context) AclOutput
}

func (*Acl) ElementType() reflect.Type {
	return reflect.TypeOf((*Acl)(nil))
}

func (i *Acl) ToAclOutput() AclOutput {
	return i.ToAclOutputWithContext(context.Background())
}

func (i *Acl) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclOutput)
}

func (i *Acl) ToAclPtrOutput() AclPtrOutput {
	return i.ToAclPtrOutputWithContext(context.Background())
}

func (i *Acl) ToAclPtrOutputWithContext(ctx context.Context) AclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclPtrOutput)
}

type AclPtrInput interface {
	pulumi.Input

	ToAclPtrOutput() AclPtrOutput
	ToAclPtrOutputWithContext(ctx context.Context) AclPtrOutput
}

type aclPtrType AclArgs

func (*aclPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil))
}

func (i *aclPtrType) ToAclPtrOutput() AclPtrOutput {
	return i.ToAclPtrOutputWithContext(context.Background())
}

func (i *aclPtrType) ToAclPtrOutputWithContext(ctx context.Context) AclPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclPtrOutput)
}

// AclArrayInput is an input type that accepts AclArray and AclArrayOutput values.
// You can construct a concrete instance of `AclArrayInput` via:
//
//          AclArray{ AclArgs{...} }
type AclArrayInput interface {
	pulumi.Input

	ToAclArrayOutput() AclArrayOutput
	ToAclArrayOutputWithContext(context.Context) AclArrayOutput
}

type AclArray []AclInput

func (AclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Acl)(nil)).Elem()
}

func (i AclArray) ToAclArrayOutput() AclArrayOutput {
	return i.ToAclArrayOutputWithContext(context.Background())
}

func (i AclArray) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclArrayOutput)
}

// AclMapInput is an input type that accepts AclMap and AclMapOutput values.
// You can construct a concrete instance of `AclMapInput` via:
//
//          AclMap{ "key": AclArgs{...} }
type AclMapInput interface {
	pulumi.Input

	ToAclMapOutput() AclMapOutput
	ToAclMapOutputWithContext(context.Context) AclMapOutput
}

type AclMap map[string]AclInput

func (AclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Acl)(nil)).Elem()
}

func (i AclMap) ToAclMapOutput() AclMapOutput {
	return i.ToAclMapOutputWithContext(context.Background())
}

func (i AclMap) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AclMapOutput)
}

type AclOutput struct{ *pulumi.OutputState }

func (AclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Acl)(nil))
}

func (o AclOutput) ToAclOutput() AclOutput {
	return o
}

func (o AclOutput) ToAclOutputWithContext(ctx context.Context) AclOutput {
	return o
}

func (o AclOutput) ToAclPtrOutput() AclPtrOutput {
	return o.ToAclPtrOutputWithContext(context.Background())
}

func (o AclOutput) ToAclPtrOutputWithContext(ctx context.Context) AclPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Acl) *Acl {
		return &v
	}).(AclPtrOutput)
}

type AclPtrOutput struct{ *pulumi.OutputState }

func (AclPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Acl)(nil))
}

func (o AclPtrOutput) ToAclPtrOutput() AclPtrOutput {
	return o
}

func (o AclPtrOutput) ToAclPtrOutputWithContext(ctx context.Context) AclPtrOutput {
	return o
}

func (o AclPtrOutput) Elem() AclOutput {
	return o.ApplyT(func(v *Acl) Acl {
		if v != nil {
			return *v
		}
		var ret Acl
		return ret
	}).(AclOutput)
}

type AclArrayOutput struct{ *pulumi.OutputState }

func (AclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Acl)(nil))
}

func (o AclArrayOutput) ToAclArrayOutput() AclArrayOutput {
	return o
}

func (o AclArrayOutput) ToAclArrayOutputWithContext(ctx context.Context) AclArrayOutput {
	return o
}

func (o AclArrayOutput) Index(i pulumi.IntInput) AclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Acl {
		return vs[0].([]Acl)[vs[1].(int)]
	}).(AclOutput)
}

type AclMapOutput struct{ *pulumi.OutputState }

func (AclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Acl)(nil))
}

func (o AclMapOutput) ToAclMapOutput() AclMapOutput {
	return o
}

func (o AclMapOutput) ToAclMapOutputWithContext(ctx context.Context) AclMapOutput {
	return o
}

func (o AclMapOutput) MapIndex(k pulumi.StringInput) AclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Acl {
		return vs[0].(map[string]Acl)[vs[1].(string)]
	}).(AclOutput)
}

func init() {
	pulumi.RegisterOutputType(AclOutput{})
	pulumi.RegisterOutputType(AclPtrOutput{})
	pulumi.RegisterOutputType(AclArrayOutput{})
	pulumi.RegisterOutputType(AclMapOutput{})
}
