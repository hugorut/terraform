package clistate

import (
	"testing"

	"github.com/hugorut/terraform/src/command/arguments"
	"github.com/hugorut/terraform/src/command/views"
	"github.com/hugorut/terraform/src/states/statemgr"
	"github.com/hugorut/terraform/src/terminal"
)

func TestUnlock(t *testing.T) {
	streams, _ := terminal.StreamsForTesting(t)
	view := views.NewView(streams)

	l := NewLocker(0, views.NewStateLocker(arguments.ViewHuman, view))
	l.Lock(statemgr.NewUnlockErrorFull(nil, nil), "test-lock")

	diags := l.Unlock()
	if diags.HasErrors() {
		t.Log(diags.Err().Error())
	} else {
		t.Error("expected error")
	}
}
