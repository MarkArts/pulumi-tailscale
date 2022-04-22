// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package tailscale

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "tailscale:index/acl:Acl":
		r = &Acl{}
	case "tailscale:index/deviceAuthorization:DeviceAuthorization":
		r = &DeviceAuthorization{}
	case "tailscale:index/deviceKey:DeviceKey":
		r = &DeviceKey{}
	case "tailscale:index/deviceSubnetRoutes:DeviceSubnetRoutes":
		r = &DeviceSubnetRoutes{}
	case "tailscale:index/deviceTags:DeviceTags":
		r = &DeviceTags{}
	case "tailscale:index/dnsNameservers:DnsNameservers":
		r = &DnsNameservers{}
	case "tailscale:index/dnsPreferences:DnsPreferences":
		r = &DnsPreferences{}
	case "tailscale:index/dnsSearchPaths:DnsSearchPaths":
		r = &DnsSearchPaths{}
	case "tailscale:index/tailnetKey:TailnetKey":
		r = &TailnetKey{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:tailscale" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/acl",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceAuthorization",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceKey",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceSubnetRoutes",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/deviceTags",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsNameservers",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsPreferences",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/dnsSearchPaths",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"tailscale",
		"index/tailnetKey",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"tailscale",
		&pkg{version},
	)
}
