package usersession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var ContextKeySessionID = usersystem.ContextKey("sessionid")

func GetSessionID(ctx context.Context) string {
	return ctx.Value(ContextKeySessionID).(string)
}

var CommandGetSession = herbsystem.Command("getsession")

func WrapGetSession(h func(st usersystem.SessionType, id string) (usersystem.Session, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandGetSession
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		session, err := h(GetSessionType(ctx), GetSessionID(ctx))
		if err != nil {
			return err
		}
		if session != nil {
			ctx = usersystem.SessionContext(ctx, session)
		}
		return next(ctx)
	}
	return a
}

func ExecGetSession(s *usersystem.UserSystem, st usersystem.SessionType, id string) (usersystem.Session, error) {
	var session usersystem.Session = nil
	ctx := usersystem.SessionContext(s.Context, session)
	ctx = SessionTypeContext(ctx, st)
	ctx = context.WithValue(ctx, ContextKeySessionID, id)
	ctx, err := s.System.ExecActions(ctx, CommandGetSession)
	if err != nil {
		return nil, err
	}
	v := ctx.Value(usersystem.ContextKeySession)
	if v == nil {
		return nil, nil
	}
	return v.(usersystem.Session), nil
}
