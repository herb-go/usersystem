package usersession

import (
	"testing"

	"github.com/herb-go/usersystem"
)

func TestRevokeSession(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()

	ok, err := ExecRevokeSession(s, "test", "")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal(ok)
	}
	ok, err = ExecRevokeSession(s, "test", "notexist")
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal(ok)
	}
	ok, err = ExecRevokeSession(s, "test", "revokecode")
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal(ok)
	}

}
