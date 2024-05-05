/*
Copyright 2021 Upbound Inc.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	"github.com/viletay/provider-vkcs/config/compute"
	"github.com/viletay/provider-vkcs/config/disk"
	"github.com/viletay/provider-vkcs/config/dns"
	"github.com/viletay/provider-vkcs/config/kubernetes"
	"github.com/viletay/provider-vkcs/config/loadbalancer"
	"github.com/viletay/provider-vkcs/config/network"
)

const (
	resourcePrefix = "vkcs"
	modulePath     = "github.com/viletay/provider-vkcs"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("vkcs.upbound.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		dns.Configure,
		disk.Configure,
		network.Configure,
		compute.Configure,
		kubernetes.Configure,
		loadbalancer.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
