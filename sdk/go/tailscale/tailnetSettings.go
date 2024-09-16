// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The tailnetSettings resource allows you to configure settings for your tailnet. See https://tailscale.com/api#tag/tailnetsettings for more information.
type TailnetSettings struct {
	pulumi.CustomResourceState

	// Whether device approval is enabled for the tailnet
	DevicesApprovalOn pulumi.BoolPtrOutput `pulumi:"devicesApprovalOn"`
	// Whether auto updates are enabled for devices that belong to this tailnet
	DevicesAutoUpdatesOn pulumi.BoolPtrOutput `pulumi:"devicesAutoUpdatesOn"`
	// The key expiry duration for devices on this tailnet
	DevicesKeyDurationDays pulumi.IntPtrOutput `pulumi:"devicesKeyDurationDays"`
	// Whether network flog logs are enabled for the tailnet
	NetworkFlowLoggingOn pulumi.BoolPtrOutput `pulumi:"networkFlowLoggingOn"`
	// Whether identity collection is enabled for device posture integrations for the tailnet
	PostureIdentityCollectionOn pulumi.BoolPtrOutput `pulumi:"postureIdentityCollectionOn"`
	// Whether regional routing is enabled for the tailnet
	RegionalRoutingOn pulumi.BoolPtrOutput `pulumi:"regionalRoutingOn"`
	// Whether user approval is enabled for this tailnet
	UsersApprovalOn pulumi.BoolPtrOutput `pulumi:"usersApprovalOn"`
	// Which user roles are allowed to join external tailnets
	UsersRoleAllowedToJoinExternalTailnet pulumi.StringPtrOutput `pulumi:"usersRoleAllowedToJoinExternalTailnet"`
}

// NewTailnetSettings registers a new resource with the given unique name, arguments, and options.
func NewTailnetSettings(ctx *pulumi.Context,
	name string, args *TailnetSettingsArgs, opts ...pulumi.ResourceOption) (*TailnetSettings, error) {
	if args == nil {
		args = &TailnetSettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource TailnetSettings
	err := ctx.RegisterResource("tailscale:index/tailnetSettings:TailnetSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTailnetSettings gets an existing TailnetSettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTailnetSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TailnetSettingsState, opts ...pulumi.ResourceOption) (*TailnetSettings, error) {
	var resource TailnetSettings
	err := ctx.ReadResource("tailscale:index/tailnetSettings:TailnetSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering TailnetSettings resources.
type tailnetSettingsState struct {
	// Whether device approval is enabled for the tailnet
	DevicesApprovalOn *bool `pulumi:"devicesApprovalOn"`
	// Whether auto updates are enabled for devices that belong to this tailnet
	DevicesAutoUpdatesOn *bool `pulumi:"devicesAutoUpdatesOn"`
	// The key expiry duration for devices on this tailnet
	DevicesKeyDurationDays *int `pulumi:"devicesKeyDurationDays"`
	// Whether network flog logs are enabled for the tailnet
	NetworkFlowLoggingOn *bool `pulumi:"networkFlowLoggingOn"`
	// Whether identity collection is enabled for device posture integrations for the tailnet
	PostureIdentityCollectionOn *bool `pulumi:"postureIdentityCollectionOn"`
	// Whether regional routing is enabled for the tailnet
	RegionalRoutingOn *bool `pulumi:"regionalRoutingOn"`
	// Whether user approval is enabled for this tailnet
	UsersApprovalOn *bool `pulumi:"usersApprovalOn"`
	// Which user roles are allowed to join external tailnets
	UsersRoleAllowedToJoinExternalTailnet *string `pulumi:"usersRoleAllowedToJoinExternalTailnet"`
}

type TailnetSettingsState struct {
	// Whether device approval is enabled for the tailnet
	DevicesApprovalOn pulumi.BoolPtrInput
	// Whether auto updates are enabled for devices that belong to this tailnet
	DevicesAutoUpdatesOn pulumi.BoolPtrInput
	// The key expiry duration for devices on this tailnet
	DevicesKeyDurationDays pulumi.IntPtrInput
	// Whether network flog logs are enabled for the tailnet
	NetworkFlowLoggingOn pulumi.BoolPtrInput
	// Whether identity collection is enabled for device posture integrations for the tailnet
	PostureIdentityCollectionOn pulumi.BoolPtrInput
	// Whether regional routing is enabled for the tailnet
	RegionalRoutingOn pulumi.BoolPtrInput
	// Whether user approval is enabled for this tailnet
	UsersApprovalOn pulumi.BoolPtrInput
	// Which user roles are allowed to join external tailnets
	UsersRoleAllowedToJoinExternalTailnet pulumi.StringPtrInput
}

func (TailnetSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*tailnetSettingsState)(nil)).Elem()
}

type tailnetSettingsArgs struct {
	// Whether device approval is enabled for the tailnet
	DevicesApprovalOn *bool `pulumi:"devicesApprovalOn"`
	// Whether auto updates are enabled for devices that belong to this tailnet
	DevicesAutoUpdatesOn *bool `pulumi:"devicesAutoUpdatesOn"`
	// The key expiry duration for devices on this tailnet
	DevicesKeyDurationDays *int `pulumi:"devicesKeyDurationDays"`
	// Whether network flog logs are enabled for the tailnet
	NetworkFlowLoggingOn *bool `pulumi:"networkFlowLoggingOn"`
	// Whether identity collection is enabled for device posture integrations for the tailnet
	PostureIdentityCollectionOn *bool `pulumi:"postureIdentityCollectionOn"`
	// Whether regional routing is enabled for the tailnet
	RegionalRoutingOn *bool `pulumi:"regionalRoutingOn"`
	// Whether user approval is enabled for this tailnet
	UsersApprovalOn *bool `pulumi:"usersApprovalOn"`
	// Which user roles are allowed to join external tailnets
	UsersRoleAllowedToJoinExternalTailnet *string `pulumi:"usersRoleAllowedToJoinExternalTailnet"`
}

// The set of arguments for constructing a TailnetSettings resource.
type TailnetSettingsArgs struct {
	// Whether device approval is enabled for the tailnet
	DevicesApprovalOn pulumi.BoolPtrInput
	// Whether auto updates are enabled for devices that belong to this tailnet
	DevicesAutoUpdatesOn pulumi.BoolPtrInput
	// The key expiry duration for devices on this tailnet
	DevicesKeyDurationDays pulumi.IntPtrInput
	// Whether network flog logs are enabled for the tailnet
	NetworkFlowLoggingOn pulumi.BoolPtrInput
	// Whether identity collection is enabled for device posture integrations for the tailnet
	PostureIdentityCollectionOn pulumi.BoolPtrInput
	// Whether regional routing is enabled for the tailnet
	RegionalRoutingOn pulumi.BoolPtrInput
	// Whether user approval is enabled for this tailnet
	UsersApprovalOn pulumi.BoolPtrInput
	// Which user roles are allowed to join external tailnets
	UsersRoleAllowedToJoinExternalTailnet pulumi.StringPtrInput
}

func (TailnetSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tailnetSettingsArgs)(nil)).Elem()
}

type TailnetSettingsInput interface {
	pulumi.Input

	ToTailnetSettingsOutput() TailnetSettingsOutput
	ToTailnetSettingsOutputWithContext(ctx context.Context) TailnetSettingsOutput
}

func (*TailnetSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**TailnetSettings)(nil)).Elem()
}

func (i *TailnetSettings) ToTailnetSettingsOutput() TailnetSettingsOutput {
	return i.ToTailnetSettingsOutputWithContext(context.Background())
}

func (i *TailnetSettings) ToTailnetSettingsOutputWithContext(ctx context.Context) TailnetSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TailnetSettingsOutput)
}

// TailnetSettingsArrayInput is an input type that accepts TailnetSettingsArray and TailnetSettingsArrayOutput values.
// You can construct a concrete instance of `TailnetSettingsArrayInput` via:
//
//	TailnetSettingsArray{ TailnetSettingsArgs{...} }
type TailnetSettingsArrayInput interface {
	pulumi.Input

	ToTailnetSettingsArrayOutput() TailnetSettingsArrayOutput
	ToTailnetSettingsArrayOutputWithContext(context.Context) TailnetSettingsArrayOutput
}

type TailnetSettingsArray []TailnetSettingsInput

func (TailnetSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TailnetSettings)(nil)).Elem()
}

func (i TailnetSettingsArray) ToTailnetSettingsArrayOutput() TailnetSettingsArrayOutput {
	return i.ToTailnetSettingsArrayOutputWithContext(context.Background())
}

func (i TailnetSettingsArray) ToTailnetSettingsArrayOutputWithContext(ctx context.Context) TailnetSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TailnetSettingsArrayOutput)
}

// TailnetSettingsMapInput is an input type that accepts TailnetSettingsMap and TailnetSettingsMapOutput values.
// You can construct a concrete instance of `TailnetSettingsMapInput` via:
//
//	TailnetSettingsMap{ "key": TailnetSettingsArgs{...} }
type TailnetSettingsMapInput interface {
	pulumi.Input

	ToTailnetSettingsMapOutput() TailnetSettingsMapOutput
	ToTailnetSettingsMapOutputWithContext(context.Context) TailnetSettingsMapOutput
}

type TailnetSettingsMap map[string]TailnetSettingsInput

func (TailnetSettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TailnetSettings)(nil)).Elem()
}

func (i TailnetSettingsMap) ToTailnetSettingsMapOutput() TailnetSettingsMapOutput {
	return i.ToTailnetSettingsMapOutputWithContext(context.Background())
}

func (i TailnetSettingsMap) ToTailnetSettingsMapOutputWithContext(ctx context.Context) TailnetSettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TailnetSettingsMapOutput)
}

type TailnetSettingsOutput struct{ *pulumi.OutputState }

func (TailnetSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TailnetSettings)(nil)).Elem()
}

func (o TailnetSettingsOutput) ToTailnetSettingsOutput() TailnetSettingsOutput {
	return o
}

func (o TailnetSettingsOutput) ToTailnetSettingsOutputWithContext(ctx context.Context) TailnetSettingsOutput {
	return o
}

// Whether device approval is enabled for the tailnet
func (o TailnetSettingsOutput) DevicesApprovalOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.BoolPtrOutput { return v.DevicesApprovalOn }).(pulumi.BoolPtrOutput)
}

// Whether auto updates are enabled for devices that belong to this tailnet
func (o TailnetSettingsOutput) DevicesAutoUpdatesOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.BoolPtrOutput { return v.DevicesAutoUpdatesOn }).(pulumi.BoolPtrOutput)
}

// The key expiry duration for devices on this tailnet
func (o TailnetSettingsOutput) DevicesKeyDurationDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.IntPtrOutput { return v.DevicesKeyDurationDays }).(pulumi.IntPtrOutput)
}

// Whether network flog logs are enabled for the tailnet
func (o TailnetSettingsOutput) NetworkFlowLoggingOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.BoolPtrOutput { return v.NetworkFlowLoggingOn }).(pulumi.BoolPtrOutput)
}

// Whether identity collection is enabled for device posture integrations for the tailnet
func (o TailnetSettingsOutput) PostureIdentityCollectionOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.BoolPtrOutput { return v.PostureIdentityCollectionOn }).(pulumi.BoolPtrOutput)
}

// Whether regional routing is enabled for the tailnet
func (o TailnetSettingsOutput) RegionalRoutingOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.BoolPtrOutput { return v.RegionalRoutingOn }).(pulumi.BoolPtrOutput)
}

// Whether user approval is enabled for this tailnet
func (o TailnetSettingsOutput) UsersApprovalOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.BoolPtrOutput { return v.UsersApprovalOn }).(pulumi.BoolPtrOutput)
}

// Which user roles are allowed to join external tailnets
func (o TailnetSettingsOutput) UsersRoleAllowedToJoinExternalTailnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TailnetSettings) pulumi.StringPtrOutput { return v.UsersRoleAllowedToJoinExternalTailnet }).(pulumi.StringPtrOutput)
}

type TailnetSettingsArrayOutput struct{ *pulumi.OutputState }

func (TailnetSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*TailnetSettings)(nil)).Elem()
}

func (o TailnetSettingsArrayOutput) ToTailnetSettingsArrayOutput() TailnetSettingsArrayOutput {
	return o
}

func (o TailnetSettingsArrayOutput) ToTailnetSettingsArrayOutputWithContext(ctx context.Context) TailnetSettingsArrayOutput {
	return o
}

func (o TailnetSettingsArrayOutput) Index(i pulumi.IntInput) TailnetSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *TailnetSettings {
		return vs[0].([]*TailnetSettings)[vs[1].(int)]
	}).(TailnetSettingsOutput)
}

type TailnetSettingsMapOutput struct{ *pulumi.OutputState }

func (TailnetSettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*TailnetSettings)(nil)).Elem()
}

func (o TailnetSettingsMapOutput) ToTailnetSettingsMapOutput() TailnetSettingsMapOutput {
	return o
}

func (o TailnetSettingsMapOutput) ToTailnetSettingsMapOutputWithContext(ctx context.Context) TailnetSettingsMapOutput {
	return o
}

func (o TailnetSettingsMapOutput) MapIndex(k pulumi.StringInput) TailnetSettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *TailnetSettings {
		return vs[0].(map[string]*TailnetSettings)[vs[1].(string)]
	}).(TailnetSettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TailnetSettingsInput)(nil)).Elem(), &TailnetSettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*TailnetSettingsArrayInput)(nil)).Elem(), TailnetSettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TailnetSettingsMapInput)(nil)).Elem(), TailnetSettingsMap{})
	pulumi.RegisterOutputType(TailnetSettingsOutput{})
	pulumi.RegisterOutputType(TailnetSettingsArrayOutput{})
	pulumi.RegisterOutputType(TailnetSettingsMapOutput{})
}
