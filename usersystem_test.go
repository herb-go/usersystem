package usersystem

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
)

func testSession(id string) *Session {
	p := authority.NewPayloads()
	p.Set(PayloadUID, []byte(id))
	return NewSession().WithType("test").WithPayloads(p)
}

func TestUserSystem(t *testing.T) {
	s := New()
	if GetUsersystem(s.context) != s {
		t.Fatal(s)
	}
	if GetUID(UIDContext(s.context, "test")) != "test" {
		t.Fatal(s)
	}
	session := testSession("test")
	if GetSession(SessionContext(s.context, session)) != session {
		t.Fatal(s)
	}
}
