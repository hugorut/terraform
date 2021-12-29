package terraform

import "github.com/hugorut/terraform/src/dag"

// GraphDot returns the dot formatting of a visual representation of
// the given Terraform graph.
func GraphDot(g *Graph, opts *dag.DotOpts) (string, error) {
	return string(g.Dot(opts)), nil
}
