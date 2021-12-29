package etcdv2

import (
	"testing"

	"github.com/hugorut/terraform/src/backend"
)

func TestBackend_impl(t *testing.T) {
	var _ backend.Backend = new(Backend)
}
