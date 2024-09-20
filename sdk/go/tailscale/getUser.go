// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-tailscale/sdk/go/tailscale/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The user data source describes a single user in a tailnet
func GetUser(ctx *pulumi.Context, args *GetUserArgs, opts ...pulumi.InvokeOption) (*GetUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetUserResult
	err := ctx.Invoke("tailscale:index/getUser:getUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getUser.
type GetUserArgs struct {
	// The unique identifier for the user.
	Id *string `pulumi:"id"`
	// The emailish login name of the user.
	LoginName *string `pulumi:"loginName"`
}

// A collection of values returned by getUser.
type GetUserResult struct {
	// The time the user joined their tailnet.
	Created string `pulumi:"created"`
	// true when the user has a node currently connected to the control server.
	CurrentlyConnected bool `pulumi:"currentlyConnected"`
	// Number of devices the user owns.
	DeviceCount int `pulumi:"deviceCount"`
	// The name of the user.
	DisplayName string `pulumi:"displayName"`
	// The unique identifier for the user.
	Id *string `pulumi:"id"`
	// The later of either: a) The last time any of the user's nodes were connected to the network or b) The last time the user authenticated to any tailscale service, including the admin panel.
	LastSeen string `pulumi:"lastSeen"`
	// The emailish login name of the user.
	LoginName *string `pulumi:"loginName"`
	// The profile pic URL for the user.
	ProfilePicUrl string `pulumi:"profilePicUrl"`
	// The role of the user.
	Role string `pulumi:"role"`
	// The status of the user.
	Status string `pulumi:"status"`
	// The tailnet that owns the user.
	TailnetId string `pulumi:"tailnetId"`
	// The type of relation this user has to the tailnet associated with the request.
	Type string `pulumi:"type"`
}

func GetUserOutput(ctx *pulumi.Context, args GetUserOutputArgs, opts ...pulumi.InvokeOption) GetUserResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetUserResultOutput, error) {
			args := v.(GetUserArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv GetUserResult
			secret, err := ctx.InvokePackageRaw("tailscale:index/getUser:getUser", args, &rv, "", opts...)
			if err != nil {
				return GetUserResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(GetUserResultOutput)
			if secret {
				return pulumi.ToSecret(output).(GetUserResultOutput), nil
			}
			return output, nil
		}).(GetUserResultOutput)
}

// A collection of arguments for invoking getUser.
type GetUserOutputArgs struct {
	// The unique identifier for the user.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The emailish login name of the user.
	LoginName pulumi.StringPtrInput `pulumi:"loginName"`
}

func (GetUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserArgs)(nil)).Elem()
}

// A collection of values returned by getUser.
type GetUserResultOutput struct{ *pulumi.OutputState }

func (GetUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetUserResult)(nil)).Elem()
}

func (o GetUserResultOutput) ToGetUserResultOutput() GetUserResultOutput {
	return o
}

func (o GetUserResultOutput) ToGetUserResultOutputWithContext(ctx context.Context) GetUserResultOutput {
	return o
}

// The time the user joined their tailnet.
func (o GetUserResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Created }).(pulumi.StringOutput)
}

// true when the user has a node currently connected to the control server.
func (o GetUserResultOutput) CurrentlyConnected() pulumi.BoolOutput {
	return o.ApplyT(func(v GetUserResult) bool { return v.CurrentlyConnected }).(pulumi.BoolOutput)
}

// Number of devices the user owns.
func (o GetUserResultOutput) DeviceCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetUserResult) int { return v.DeviceCount }).(pulumi.IntOutput)
}

// The name of the user.
func (o GetUserResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The unique identifier for the user.
func (o GetUserResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The later of either: a) The last time any of the user's nodes were connected to the network or b) The last time the user authenticated to any tailscale service, including the admin panel.
func (o GetUserResultOutput) LastSeen() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.LastSeen }).(pulumi.StringOutput)
}

// The emailish login name of the user.
func (o GetUserResultOutput) LoginName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetUserResult) *string { return v.LoginName }).(pulumi.StringPtrOutput)
}

// The profile pic URL for the user.
func (o GetUserResultOutput) ProfilePicUrl() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.ProfilePicUrl }).(pulumi.StringOutput)
}

// The role of the user.
func (o GetUserResultOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Role }).(pulumi.StringOutput)
}

// The status of the user.
func (o GetUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Status }).(pulumi.StringOutput)
}

// The tailnet that owns the user.
func (o GetUserResultOutput) TailnetId() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.TailnetId }).(pulumi.StringOutput)
}

// The type of relation this user has to the tailnet associated with the request.
func (o GetUserResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetUserResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetUserResultOutput{})
}
