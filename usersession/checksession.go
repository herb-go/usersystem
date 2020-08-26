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

func WrapCheckSession(h func(ctx context.Context, session *usersystem.Session) (bool, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandCheckSession
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		result, err := h(ctx, usersystem.GetSession(ctx))
		if err != nil {
			return err
		}
		if !result {
			r := GetResult(ctx)
			r.Success = false
			return nil
		}
		return next(ctx)
	}
	return a
}

func ExecCheckSession(s *usersystem.UserSystem, session *usersystem.Session) (bool, error) {
	var err error
	ctx := usersystem.SessionContext(s.Context, session)
	result := &Result{
		Success: true,
	}
	ctx = context.WithValue(ctx, ContextKeyCheckSessionResult, result)

	ctx, err = s.System.ExecActions(ctx, CommandCheckSession)
	if err != nil {
		return false, err
	}
	return result.Success, nil
}
