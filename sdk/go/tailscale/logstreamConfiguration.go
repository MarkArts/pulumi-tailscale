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

// The logstreamConfiguration resource allows you to configure streaming configuration or network flow logs to a supported security information and event management (SIEM) system. See https://tailscale.com/kb/1255/log-streaming for more information.
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
//			_, err := tailscale.NewLogstreamConfiguration(ctx, "sample_logstream_configuration", &tailscale.LogstreamConfigurationArgs{
//				LogType:         pulumi.String("configuration"),
//				DestinationType: pulumi.String("panther"),
//				Url:             pulumi.String("https://example.com"),
//				Token:           pulumi.String("some-token"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Logstream configuration can be imported using the logstream configuration id, e.g.,
//
// ```sh
// $ pulumi import tailscale:index/logstreamConfiguration:LogstreamConfiguration sample_logstream_configuration 123456789
// ```
type LogstreamConfiguration struct {
	pulumi.CustomResourceState

	// The type of system to which logs are being streamed.
	DestinationType pulumi.StringOutput `pulumi:"destinationType"`
	// The type of log that is streamed to this endpoint.
	LogType pulumi.StringOutput `pulumi:"logType"`
	// The token/password with which log streams to this endpoint should be authenticated.
	Token pulumi.StringOutput `pulumi:"token"`
	// The URL to which log streams are being posted.
	Url pulumi.StringOutput `pulumi:"url"`
	// The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
	User pulumi.StringPtrOutput `pulumi:"user"`
}

// NewLogstreamConfiguration registers a new resource with the given unique name, arguments, and options.
func NewLogstreamConfiguration(ctx *pulumi.Context,
	name string, args *LogstreamConfigurationArgs, opts ...pulumi.ResourceOption) (*LogstreamConfiguration, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DestinationType == nil {
		return nil, errors.New("invalid value for required argument 'DestinationType'")
	}
	if args.LogType == nil {
		return nil, errors.New("invalid value for required argument 'LogType'")
	}
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogstreamConfiguration
	err := ctx.RegisterResource("tailscale:index/logstreamConfiguration:LogstreamConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogstreamConfiguration gets an existing LogstreamConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogstreamConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogstreamConfigurationState, opts ...pulumi.ResourceOption) (*LogstreamConfiguration, error) {
	var resource LogstreamConfiguration
	err := ctx.ReadResource("tailscale:index/logstreamConfiguration:LogstreamConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogstreamConfiguration resources.
type logstreamConfigurationState struct {
	// The type of system to which logs are being streamed.
	DestinationType *string `pulumi:"destinationType"`
	// The type of log that is streamed to this endpoint.
	LogType *string `pulumi:"logType"`
	// The token/password with which log streams to this endpoint should be authenticated.
	Token *string `pulumi:"token"`
	// The URL to which log streams are being posted.
	Url *string `pulumi:"url"`
	// The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
	User *string `pulumi:"user"`
}

type LogstreamConfigurationState struct {
	// The type of system to which logs are being streamed.
	DestinationType pulumi.StringPtrInput
	// The type of log that is streamed to this endpoint.
	LogType pulumi.StringPtrInput
	// The token/password with which log streams to this endpoint should be authenticated.
	Token pulumi.StringPtrInput
	// The URL to which log streams are being posted.
	Url pulumi.StringPtrInput
	// The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
	User pulumi.StringPtrInput
}

func (LogstreamConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*logstreamConfigurationState)(nil)).Elem()
}

type logstreamConfigurationArgs struct {
	// The type of system to which logs are being streamed.
	DestinationType string `pulumi:"destinationType"`
	// The type of log that is streamed to this endpoint.
	LogType string `pulumi:"logType"`
	// The token/password with which log streams to this endpoint should be authenticated.
	Token string `pulumi:"token"`
	// The URL to which log streams are being posted.
	Url string `pulumi:"url"`
	// The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
	User *string `pulumi:"user"`
}

// The set of arguments for constructing a LogstreamConfiguration resource.
type LogstreamConfigurationArgs struct {
	// The type of system to which logs are being streamed.
	DestinationType pulumi.StringInput
	// The type of log that is streamed to this endpoint.
	LogType pulumi.StringInput
	// The token/password with which log streams to this endpoint should be authenticated.
	Token pulumi.StringInput
	// The URL to which log streams are being posted.
	Url pulumi.StringInput
	// The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
	User pulumi.StringPtrInput
}

func (LogstreamConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logstreamConfigurationArgs)(nil)).Elem()
}

type LogstreamConfigurationInput interface {
	pulumi.Input

	ToLogstreamConfigurationOutput() LogstreamConfigurationOutput
	ToLogstreamConfigurationOutputWithContext(ctx context.Context) LogstreamConfigurationOutput
}

func (*LogstreamConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((**LogstreamConfiguration)(nil)).Elem()
}

func (i *LogstreamConfiguration) ToLogstreamConfigurationOutput() LogstreamConfigurationOutput {
	return i.ToLogstreamConfigurationOutputWithContext(context.Background())
}

func (i *LogstreamConfiguration) ToLogstreamConfigurationOutputWithContext(ctx context.Context) LogstreamConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogstreamConfigurationOutput)
}

// LogstreamConfigurationArrayInput is an input type that accepts LogstreamConfigurationArray and LogstreamConfigurationArrayOutput values.
// You can construct a concrete instance of `LogstreamConfigurationArrayInput` via:
//
//	LogstreamConfigurationArray{ LogstreamConfigurationArgs{...} }
type LogstreamConfigurationArrayInput interface {
	pulumi.Input

	ToLogstreamConfigurationArrayOutput() LogstreamConfigurationArrayOutput
	ToLogstreamConfigurationArrayOutputWithContext(context.Context) LogstreamConfigurationArrayOutput
}

type LogstreamConfigurationArray []LogstreamConfigurationInput

func (LogstreamConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogstreamConfiguration)(nil)).Elem()
}

func (i LogstreamConfigurationArray) ToLogstreamConfigurationArrayOutput() LogstreamConfigurationArrayOutput {
	return i.ToLogstreamConfigurationArrayOutputWithContext(context.Background())
}

func (i LogstreamConfigurationArray) ToLogstreamConfigurationArrayOutputWithContext(ctx context.Context) LogstreamConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogstreamConfigurationArrayOutput)
}

// LogstreamConfigurationMapInput is an input type that accepts LogstreamConfigurationMap and LogstreamConfigurationMapOutput values.
// You can construct a concrete instance of `LogstreamConfigurationMapInput` via:
//
//	LogstreamConfigurationMap{ "key": LogstreamConfigurationArgs{...} }
type LogstreamConfigurationMapInput interface {
	pulumi.Input

	ToLogstreamConfigurationMapOutput() LogstreamConfigurationMapOutput
	ToLogstreamConfigurationMapOutputWithContext(context.Context) LogstreamConfigurationMapOutput
}

type LogstreamConfigurationMap map[string]LogstreamConfigurationInput

func (LogstreamConfigurationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogstreamConfiguration)(nil)).Elem()
}

func (i LogstreamConfigurationMap) ToLogstreamConfigurationMapOutput() LogstreamConfigurationMapOutput {
	return i.ToLogstreamConfigurationMapOutputWithContext(context.Background())
}

func (i LogstreamConfigurationMap) ToLogstreamConfigurationMapOutputWithContext(ctx context.Context) LogstreamConfigurationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogstreamConfigurationMapOutput)
}

type LogstreamConfigurationOutput struct{ *pulumi.OutputState }

func (LogstreamConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogstreamConfiguration)(nil)).Elem()
}

func (o LogstreamConfigurationOutput) ToLogstreamConfigurationOutput() LogstreamConfigurationOutput {
	return o
}

func (o LogstreamConfigurationOutput) ToLogstreamConfigurationOutputWithContext(ctx context.Context) LogstreamConfigurationOutput {
	return o
}

// The type of system to which logs are being streamed.
func (o LogstreamConfigurationOutput) DestinationType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogstreamConfiguration) pulumi.StringOutput { return v.DestinationType }).(pulumi.StringOutput)
}

// The type of log that is streamed to this endpoint.
func (o LogstreamConfigurationOutput) LogType() pulumi.StringOutput {
	return o.ApplyT(func(v *LogstreamConfiguration) pulumi.StringOutput { return v.LogType }).(pulumi.StringOutput)
}

// The token/password with which log streams to this endpoint should be authenticated.
func (o LogstreamConfigurationOutput) Token() pulumi.StringOutput {
	return o.ApplyT(func(v *LogstreamConfiguration) pulumi.StringOutput { return v.Token }).(pulumi.StringOutput)
}

// The URL to which log streams are being posted.
func (o LogstreamConfigurationOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *LogstreamConfiguration) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The username with which log streams to this endpoint are authenticated. Only required if destinationType is 'elastic', defaults to 'user' if not set.
func (o LogstreamConfigurationOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogstreamConfiguration) pulumi.StringPtrOutput { return v.User }).(pulumi.StringPtrOutput)
}

type LogstreamConfigurationArrayOutput struct{ *pulumi.OutputState }

func (LogstreamConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogstreamConfiguration)(nil)).Elem()
}

func (o LogstreamConfigurationArrayOutput) ToLogstreamConfigurationArrayOutput() LogstreamConfigurationArrayOutput {
	return o
}

func (o LogstreamConfigurationArrayOutput) ToLogstreamConfigurationArrayOutputWithContext(ctx context.Context) LogstreamConfigurationArrayOutput {
	return o
}

func (o LogstreamConfigurationArrayOutput) Index(i pulumi.IntInput) LogstreamConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogstreamConfiguration {
		return vs[0].([]*LogstreamConfiguration)[vs[1].(int)]
	}).(LogstreamConfigurationOutput)
}

type LogstreamConfigurationMapOutput struct{ *pulumi.OutputState }

func (LogstreamConfigurationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogstreamConfiguration)(nil)).Elem()
}

func (o LogstreamConfigurationMapOutput) ToLogstreamConfigurationMapOutput() LogstreamConfigurationMapOutput {
	return o
}

func (o LogstreamConfigurationMapOutput) ToLogstreamConfigurationMapOutputWithContext(ctx context.Context) LogstreamConfigurationMapOutput {
	return o
}

func (o LogstreamConfigurationMapOutput) MapIndex(k pulumi.StringInput) LogstreamConfigurationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogstreamConfiguration {
		return vs[0].(map[string]*LogstreamConfiguration)[vs[1].(string)]
	}).(LogstreamConfigurationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogstreamConfigurationInput)(nil)).Elem(), &LogstreamConfiguration{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogstreamConfigurationArrayInput)(nil)).Elem(), LogstreamConfigurationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogstreamConfigurationMapInput)(nil)).Elem(), LogstreamConfigurationMap{})
	pulumi.RegisterOutputType(LogstreamConfigurationOutput{})
	pulumi.RegisterOutputType(LogstreamConfigurationArrayOutput{})
	pulumi.RegisterOutputType(LogstreamConfigurationMapOutput{})
}
