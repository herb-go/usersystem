package usersession

import (
	"context"
	"time"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type ActiveSession struct {
	SessionID  string
	LastActive time.Time
}

var CommandGetActiveSessions = herbsystem.Command("getactivesessions")

var ContextKeyActiveSessions = usersystem.ContextKey("usersession.activesessions")

func GetActiveSessions(ctx context.Context) []*ActiveSession {
	v := ctx.Value(ContextKeyActiveSessions)
	if v == nil {
		return nil
	}
	return v.([]*ActiveSession)
}

func WrapGetActiveSessions(h func(usersystem.SessionType) ([]*ActiveSession, bool, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandGetActiveSessions
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		sessions, ok, err := h(GetSessionType(ctx))
		if err != nil {
			return err
		}
		if ok {
			ctx = context.WithValue(ctx, ContextKeyActiveSessions, sessions)
		}
		return next(ctx)
	}
	return a
}

func ExecGetActiveSessions(s *usersystem.UserSystem, st usersystem.SessionType) ([]*ActiveSession, error) {
	ctx := SessionTypeContext(s.Context, st)
	ctx = context.WithValue(ctx, ContextKeyActiveSessions, nil)

	ctx, err := s.System.ExecActions(ctx, CommandGetActiveSessions)
	if err != nil {
		return nil, err
	}
	return GetActiveSessions(ctx), nil
}
