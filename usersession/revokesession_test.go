package usersession

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem"
)

func TestRevokeSession(t *testing.T) {
	s := usersystem.New()
	s.MustRegisterSystemModule(&testModule{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)

	session := usersystem.NewSession().WithType("test").WithPayloads(authority.NewPayloads())
	session.Payloads.Set(usersystem.PayloadRevokeCode, []byte(""))
	ok := MustExecRevokeSession(s, session)
	if ok {
		t.Fatal(ok)
	}
	session = usersystem.NewSession().WithType("test").WithPayloads(authority.NewPayloads())
	session.Payloads.Set(usersystem.PayloadRevokeCode, []byte("notexist"))
	ok = MustExecRevokeSession(s, session)
	if ok {
		t.Fatal(ok)
	}
	session = usersystem.NewSession().WithType("test").WithPayloads(authority.NewPayloads())
	session.Payloads.Set(usersystem.PayloadRevokeCode, []byte("revokecode"))
	ok = MustExecRevokeSession(s, session)
	if !ok {
		t.Fatal(ok)
	}

}
