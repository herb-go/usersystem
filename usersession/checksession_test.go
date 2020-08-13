package usersession

import (
	"testing"

	"github.com/herb-go/usersystem"
)

func TestCheckSession(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	ok, err := ExecCheckSession(s, testSession("exists"))
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = ExecCheckSession(s, testSession("notexists"))
	if ok || err != nil {
		t.Fatal(ok, err)
	}
}
