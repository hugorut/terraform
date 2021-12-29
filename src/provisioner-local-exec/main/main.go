package main

import (
	localexec "github.com/hugorut/terraform/src/builtin/provisioners/local-exec"
	"github.com/hugorut/terraform/src/grpcwrap"
	"github.com/hugorut/terraform/src/plugin"
	"github.com/hugorut/terraform/src/tfplugin5"
)

func main() {
	// Provide a binary version of the src terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
