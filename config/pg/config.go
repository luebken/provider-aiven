package repository

import "github.com/upbound/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("aiven_pg", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		// this resource
		r.ShortGroup = "pg"
	})
}
