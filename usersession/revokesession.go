package usersession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var ContextSessionRevokeCode = usersystem.ContextKey("sessionrevokecode")

func GetSessionRevokeCode(ctx context.Context) string {
	return ctx.Value(ContextSessionRevokeCode).(string)
}

var CommandRevokeSession = herbsystem.Command("revokesession")

func WrapRevokeSession(h func(st usersystem.SessionType, code string) bool) *herbsystem.Action {
	return herbsystem.CreateAction(CommandRevokeSession, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		ok := h(usersystem.GetSessionType(ctx), GetSessionRevokeCode(ctx))
		if ok {
			r := GetResult(ctx)
			r.Success = true
			return
		}
		next(ctx, system)

	})
}

func MustExecRevokeSession(s *usersystem.UserSystem, session *usersystem.Session) bool {
	code := session.RevokeCode()
	if code == "" {
		return false
	}
	st := session.Type
	ctx := usersystem.SessionTypeContext(s.SystemContext(), st)

	ctx = context.WithValue(ctx, ContextSessionRevokeCode, code)
	result := &Result{
		Success: false,
	}
	ctx = context.WithValue(ctx, ContextKeyCheckSessionResult, result)

	herbsystem.MustExecActions(ctx, s, CommandRevokeSession)
	return result.Success
}
