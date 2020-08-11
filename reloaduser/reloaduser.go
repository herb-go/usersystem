package reloaduser

import (
	"context"

	"github.com/herb-go/usersystem"

	"github.com/herb-go/herbsystem"
)

var Command = herbsystem.Command("reload")

func Wrap(h func(string) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = Command
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(usersystem.GetUID(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecReload(s *usersystem.UserSystem, uid string) error {
	ctx := usersystem.UIDContext(s.Context, uid)
	_, err := s.System.ExecActions(ctx, Command)
	return err
}
