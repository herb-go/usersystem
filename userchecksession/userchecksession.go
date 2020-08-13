package userchecksession

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type Result struct {
	Success bool
}

var ContextKeyCheckSessionResult = usersystem.ContextKey("checksession.result")

func GetResult(ctx context.Context) *Result {
	return ctx.Value(ContextKeyCheckSessionResult).(*Result)
}

var ContextKeyPayloads = usersystem.ContextKey("checksession.payloads")

func GetPayloads(ctx context.Context) *authority.Payloads {
	return ctx.Value(ContextKeyPayloads).(*authority.Payloads)
}

var Command = herbsystem.Command("checksession")

func Wrap(h func(usersystem.Session, string, *authority.Payloads) (bool, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = Command
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		result, err := h(usersystem.GetSession(ctx), usersystem.GetUID(ctx), GetPayloads(ctx))
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

func ExecCheckSession(s *usersystem.UserSystem, session usersystem.Session) (bool, error) {
	uid, err := session.UID()
	if err != nil {
		return false, err
	}
	payloads, err := session.Payloads()
	if err != nil {
		return false, err
	}
	ctx := usersystem.SessionContext(s.Context, session)
	ctx = usersystem.UIDContext(ctx, uid)
	ctx = context.WithValue(ctx, ContextKeyPayloads, payloads)
	result := &Result{
		Success: true,
	}
	ctx = context.WithValue(ctx, ContextKeyCheckSessionResult, result)

	ctx, err = s.System.ExecActions(ctx, Command)
	if err != nil {
		return false, err
	}
	return result.Success, nil
}
