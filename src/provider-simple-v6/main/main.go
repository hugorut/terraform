package main

import (
	"github.com/hugorut/terraform/src/grpcwrap"
	plugin "github.com/hugorut/terraform/src/plugin6"
	simple "github.com/hugorut/terraform/src/provider-simple-v6"
	"github.com/hugorut/terraform/src/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
