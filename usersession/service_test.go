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

type testModule struct {
	herbsystem.NopModule
}

func (s *testModule) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		WrapCheckSession(func(ctx context.Context, session *usersystem.Session) bool {
			return session.UID() == "exists"
		}),
		WrapInitPayloads(func(ctx context.Context, st usersystem.SessionType, id string, payloads *authority.Payloads) {
			payloads.Set("test", []byte("testvalue"))
		}),
		WrapOnSessionActive(func(session *usersystem.Session) {
			lastactive = session.UID()
		}),
		WrapGetSession(func(st usersystem.SessionType, id string) *usersystem.Session {
			if st != usersystem.SessionType("test") {
				return nil
			}
			return testSession("got")
		}),
		WrapRevokeSession(func(st usersystem.SessionType, code string) bool {
			if st != usersystem.SessionType("test") {
				return false
			}
			return code == "revokecode"
		}),
	)
	next(ctx, system)
}
