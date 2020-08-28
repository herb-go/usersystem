package usersession

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"

	"github.com/herb-go/usersystem"
)

func TestRevokeSession(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	session := usersystem.NewSession().WithType("test").WithPayloads(authority.NewPayloads())
	session.Payloads.Set(usersystem.PayloadRevokeCode, []byte(""))
	ok, err := ExecRevokeSession(s, session)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal(ok)
	}
	session = usersystem.NewSession().WithType("test").WithPayloads(authority.NewPayloads())
	session.Payloads.Set(usersystem.PayloadRevokeCode, []byte("notexist"))
	ok, err = ExecRevokeSession(s, session)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal(ok)
	}
	session = usersystem.NewSession().WithType("test").WithPayloads(authority.NewPayloads())
	session.Payloads.Set(usersystem.PayloadRevokeCode, []byte("revokecode"))
	ok, err = ExecRevokeSession(s, session)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal(ok)
	}

}
