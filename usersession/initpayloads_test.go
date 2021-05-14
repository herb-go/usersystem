package usersession

import (
	"testing"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

func TestInitPayloads(t *testing.T) {
	s := usersystem.New()
	s.MustRegisterSystemModule(&testModule{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)

	payloads := MustExecInitPayloads(s, s.SystemContext(), "test", "exists")
	if payloads.LoadString(usersystem.PayloadUID) != "exists" {
		t.Fatal()
	}
	if payloads.LoadString("test") != "testvalue" {
		t.Fatal()
	}
}
