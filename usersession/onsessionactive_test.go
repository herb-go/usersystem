package usersession

import (
	"testing"

	"github.com/herb-go/usersystem"
)

func TestOnSessionActive(t *testing.T) {
	lastactive = ""
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	if lastactive != "" {
		t.Fatal()
	}
	err := ExecOnSessionActive(s, testSession("exists"))
	if err != nil {
		t.Fatal(err)
	}
	if lastactive != "exists" {
		t.Fatal()
	}
}
