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

func WrapGetSession(h func(st usersystem.SessionType, id string) *usersystem.Session) *herbsystem.Action {
	return herbsystem.CreateAction(CommandGetSession, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		session := h(usersystem.GetSessionType(ctx), GetSessionID(ctx))
		if session != nil {
			ctx = usersystem.SessionContext(ctx, session)
		}
		next(ctx, system)

	})
}

func MustExecGetSession(s *usersystem.UserSystem, st usersystem.SessionType, id string) *usersystem.Session {
	var session *usersystem.Session
	ctx := usersystem.SessionContext(s.SystemContext(), session)
	ctx = usersystem.SessionTypeContext(ctx, st)
	ctx = context.WithValue(ctx, ContextKeySessionID, id)
	ctx = herbsystem.MustExecActions(ctx, s, CommandGetSession)
	v := ctx.Value(usersystem.ContextKeySession)
	if v == nil {
		return nil
	}
	return v.(*usersystem.Session)
}
