package useravaliable

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

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
			ctx = context.WithValue(ctx, ContextKeyNotAvaliable, true)
			return nil
		}
		return next(ctx)
	}
	return a
}

func ExecAvaliable(s *usersystem.UserSystem, uid string) (bool, error) {
	ctx := usersystem.UIDContext(s.Context, uid)
	ctx = context.WithValue(ctx, ContextKeyNotAvaliable, false)
	ctx, err := s.System.ExecActions(ctx, Command)
	if err != nil {
		return false, err
	}
	return !ctx.Value(ContextKeyNotAvaliable).(bool), nil
}
