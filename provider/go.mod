module github.com/pulumi/pulumi-tailscale/provider

go 1.16

replace (
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211019193655-b23eb83463bb
)

require (
	github.com/davidsbond/terraform-provider-tailscale v0.9.0
	github.com/hashicorp/terraform-plugin-sdk v1.9.1 // indirect
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.20.0
	github.com/pulumi/pulumi/sdk/v3 v3.27.0
)
