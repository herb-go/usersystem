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

func WrapRevokeSession(h func(st usersystem.SessionType, code string) (bool, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandRevokeSession
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		ok, err := h(GetSessionType(ctx), GetSessionRevokeCode(ctx))
		if err != nil {
			return err
		}
		if ok {
			r := GetResult(ctx)
			r.Success = true
			return nil
		}
		return next(ctx)
	}
	return a
}

func ExecRevokeSession(s *usersystem.UserSystem, st usersystem.SessionType, code string) (bool, error) {
	if code == "" {
		return false, nil
	}
	ctx := SessionTypeContext(s.Context, st)

	ctx = context.WithValue(ctx, ContextSessionRevokeCode, code)
	result := &Result{
		Success: false,
	}
	ctx = context.WithValue(ctx, ContextKeyCheckSessionResult, result)

	ctx, err := s.System.ExecActions(ctx, CommandRevokeSession)
	if err != nil {
		return false, err
	}
	return result.Success, nil
}
