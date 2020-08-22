package usersession

import (
	"context"
	"time"

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
		WrapCheckSession(func(session *usersystem.Session) (bool, error) {
			return session.UID() == "exists", nil
		}),
		WrapInitPayloads(func(ctx context.Context, id string, payloads *authority.Payloads) error {
			payloads.Set("test", []byte("testvalue"))
			return nil
		}),
		WrapOnSessionActive(func(session *usersystem.Session) error {
			lastactive = session.UID()
			return nil
		}),
		WrapActiveSessionManagerConfig(func(st usersystem.SessionType) (*Config, error) {
			if st != usersystem.SessionType("test") {
				return nil, nil
			}
			return &Config{
				Supported: true,
				Duration:  time.Minute,
			}, nil
		}),
		WrapGetActiveSessions(func(st usersystem.SessionType, uid string) ([]*ActiveSession, bool, error) {
			if st != usersystem.SessionType("test") || uid != "test" {
				return nil, false, nil
			}
			return []*ActiveSession{
				&ActiveSession{
					SessionID: "active",
				},
			}, true, nil
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
