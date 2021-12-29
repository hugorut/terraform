package main

import (
	"github.com/hugorut/terraform/src/builtin/providers/terraform"
	"github.com/hugorut/terraform/src/grpcwrap"
	"github.com/hugorut/terraform/src/plugin"
	"github.com/hugorut/terraform/src/tfplugin5"
)

func main() {
	// Provide a binary version of the src terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
