package usersession

import (
	"testing"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

func TestOnSessionActive(t *testing.T) {
	lastactive = ""
	s := usersystem.New()
	s.MustRegisterSystemModule(&testModule{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)

	if lastactive != "" {
		t.Fatal()
	}
	MustExecOnSessionActive(s, testSession("exists"))
	if lastactive != "exists" {
		t.Fatal()
	}
}
