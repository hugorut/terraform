package main

import (
	"github.com/hashicorp/terraform/src/grpcwrap"
	plugin "github.com/hashicorp/terraform/src/plugin6"
	simple "github.com/hashicorp/terraform/src/provider-simple-v6"
	"github.com/hashicorp/terraform/src/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
