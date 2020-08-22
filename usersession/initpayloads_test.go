package usersession

import (
	"testing"

	"github.com/herb-go/usersystem"
)

func TestInitPayloads(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	payloads, err := ExecInitPayloads(s, s.Context, "exists")
	if err != nil {
		t.Fatal(err)
	}
	if payloads.LoadString(usersystem.PayloadUID) != "exists" {
		t.Fatal()
	}
	if payloads.LoadString("test") != "testvalue" {
		t.Fatal()
	}
}
