package main

import (
	"github.com/hugorut/terraform/src/grpcwrap"
	"github.com/hugorut/terraform/src/plugin"
	simple "github.com/hugorut/terraform/src/provider-simple"
	"github.com/hugorut/terraform/src/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
