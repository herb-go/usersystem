package useravaliable

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type Result struct {
	Avaliable bool
}

func GetResult(ctx context.Context) *Result {
	return ctx.Value(ContextKeyNotAvaliable).(*Result)
}

var ContextKeyNotAvaliable = usersystem.ContextKey("notavaliable")
var Command = herbsystem.Command("avaliable")

func Wrap(h func(string) (bool, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = Command
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		result, err := h(usersystem.GetUID(ctx))
		if err != nil {
			return err
		}
		if !result {
			r := GetResult(ctx)
			r.Avaliable = false
			return nil
		}
		return next(ctx)
	}
	return a
}

func ExecAvaliable(s *usersystem.UserSystem, uid string) (bool, error) {
	ctx := usersystem.UIDContext(s.Context, uid)
	result := &Result{
		Avaliable: true,
	}
	ctx = context.WithValue(ctx, ContextKeyNotAvaliable, result)
	ctx, err := s.System.ExecActions(ctx, Command)
	if err != nil {
		return false, err
	}
	return result.Avaliable, nil
}
