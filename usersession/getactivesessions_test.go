package usersession

import (
	"testing"

	"github.com/herb-go/usersystem"
)

func TestGetActiveSessions(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()

	sessions, err := ExecGetActiveSessions(s, "test", "test")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 {
		t.Fatal(sessions)
	}
	if sessions[0].SessionID != "active" {
		t.Fatal(s)
	}
	sessions, err = ExecGetActiveSessions(s, "test", "notexists")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 0 {
		t.Fatal(sessions)
	}

}
