package usersession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type Result struct {
	Success bool
}

var ContextKeyCheckSessionResult = usersystem.ContextKey("usersession.result")

func GetResult(ctx context.Context) *Result {
	return ctx.Value(ContextKeyCheckSessionResult).(*Result)
}

var CommandCheckSession = herbsystem.Command("checksession")

func WrapCheckSession(h func(ctx context.Context, session *usersystem.Session) bool) *herbsystem.Action {
	return herbsystem.CreateAction(CommandCheckSession, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		result := h(ctx, usersystem.GetSession(ctx))
		if !result {
			r := GetResult(ctx)
			r.Success = false
			return
		}
		next(ctx, system)
	})
}

func MustExecCheckSession(s *usersystem.UserSystem, session *usersystem.Session) bool {
	ctx := usersystem.SessionContext(s.SystemContext(), session)
	result := &Result{
		Success: true,
	}
	ctx = context.WithValue(ctx, ContextKeyCheckSessionResult, result)
	ctx = herbsystem.MustExecActions(ctx, s, CommandCheckSession)
	return result.Success
}
