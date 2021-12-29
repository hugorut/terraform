package terraform

import (
	"testing"

	"github.com/hugorut/terraform/src/addrs"
	"github.com/hugorut/terraform/src/configs"
	"github.com/zclconf/go-cty/cty"
)

func TestNodeRootVariableExecute(t *testing.T) {
	ctx := new(MockEvalContext)

	n := &NodeRootVariable{
		Addr: addrs.InputVariable{Name: "foo"},
		Config: &configs.Variable{
			Name:           "foo",
			Type:           cty.String,
			ConstraintType: cty.String,
		},
	}

	diags := n.Execute(ctx, walkApply)
	if diags.HasErrors() {
		t.Fatalf("unexpected error: %s", diags.Err())
	}

}
