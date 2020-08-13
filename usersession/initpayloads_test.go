package usersession

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

func TestInitPayloads(t *testing.T) {
	payloads = authority.NewPayloads()
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	if payloads.LoadString("test") != "" {
		t.Fatal()
	}
	err := ExecInitPayloads(s, testSession("exists"))
	if err != nil {
		t.Fatal(err)
	}
	if payloads.LoadString("test") != "testvalue" {
		t.Fatal()
	}
}
