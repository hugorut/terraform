package local

import (
	"flag"
	"os"
	"testing"

	_ "github.com/hugorut/terraform/src/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
