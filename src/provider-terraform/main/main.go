package main

import (
	"github.com/hashicorp/terraform/src/builtin/providers/terraform"
	"github.com/hashicorp/terraform/src/grpcwrap"
	"github.com/hashicorp/terraform/src/plugin"
	"github.com/hashicorp/terraform/src/tfplugin5"
)

func main() {
	// Provide a binary version of the src terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
