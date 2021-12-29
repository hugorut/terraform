package terraform

import (
	backendInit "github.com/hugorut/terraform/src/backend/init"
)

func init() {
	// Initialize the backends
	backendInit.Init(nil)
}
