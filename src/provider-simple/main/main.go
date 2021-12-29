package main

import (
	"github.com/hashicorp/terraform/src/grpcwrap"
	"github.com/hashicorp/terraform/src/plugin"
	simple "github.com/hashicorp/terraform/src/provider-simple"
	"github.com/hashicorp/terraform/src/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
