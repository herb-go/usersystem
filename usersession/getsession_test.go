package usersession

import (
	"testing"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

func TestGetSession(t *testing.T) {
	s := usersystem.New()
	s.MustRegisterSystemModule(&testModule{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)

	session := MustExecGetSession(s, "test", "ttt")
	if session.UID() != "got" {
		t.Fatal(session)
	}
	session = MustExecGetSession(s, "notexists", "ttt")
	if session != nil {
		t.Fatal(session)
	}
}
