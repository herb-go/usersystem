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

	sessions, err := ExecGetActiveSessions(s, "test")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 1 {
		t.Fatal(sessions)
	}
	uid, err := sessions[0].Session.UID()
	if uid != "active" || err != nil {
		t.Fatal(uid, err)
	}
	sessions, err = ExecGetActiveSessions(s, "notexists")
	if err != nil {
		t.Fatal(err)
	}
	if len(sessions) != 0 {
		t.Fatal(sessions)
	}

}
