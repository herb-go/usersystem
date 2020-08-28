package usersession

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var payloads = authority.NewPayloads()

var lastactive = ""

func testSession(id string) *usersystem.Session {
	p := authority.NewPayloads()
	p.Set(usersystem.PayloadUID, []byte(id))
	return usersystem.NewSession().WithType("test").WithPayloads(p)
}

type testService struct {
	herbsystem.NopService
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		WrapCheckSession(func(ctx context.Context, session *usersystem.Session) (bool, error) {
			return session.UID() == "exists", nil
		}),
		WrapInitPayloads(func(ctx context.Context, st usersystem.SessionType, id string, payloads *authority.Payloads) error {
			payloads.Set("test", []byte("testvalue"))
			return nil
		}),
		WrapOnSessionActive(func(session *usersystem.Session) error {
			lastactive = session.UID()
			return nil
		}),
		WrapGetSession(func(st usersystem.SessionType, id string) (*usersystem.Session, error) {
			if st != usersystem.SessionType("test") {
				return nil, nil
			}
			return testSession("got"), nil
		}),
		WrapRevokeSession(func(st usersystem.SessionType, code string) (bool, error) {
			if st != usersystem.SessionType("test") {
				return false, nil
			}
			return code == "revokecode", nil
		}),
	}
}
