package main

import (
	localexec "github.com/hashicorp/terraform/src/builtin/provisioners/local-exec"
	"github.com/hashicorp/terraform/src/grpcwrap"
	"github.com/hashicorp/terraform/src/plugin"
	"github.com/hashicorp/terraform/src/tfplugin5"
)

func main() {
	// Provide a binary version of the src terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
