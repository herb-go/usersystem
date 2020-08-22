package usersession

import (
	"testing"

	"github.com/herb-go/usersystem"
)

func TestGetSession(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()

	session, err := ExecGetSession(s, "test", "ttt")
	if err != nil {
		t.Fatal(err)
	}
	if session.UID() != "got" {
		t.Fatal(session)
	}
	session, err = ExecGetSession(s, "notexists", "ttt")
	if err != nil {
		t.Fatal(err)
	}
	if session != nil {
		t.Fatal(session)
	}
}
