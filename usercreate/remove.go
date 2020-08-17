package usercreate

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandRemove = herbsystem.Command("remove")

func WrapRemove(h func(string) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandRemove
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {

		err := h(usersystem.GetUID(ctx))
		if err != nil {
			return herbsystem.MergeError(next(ctx), err)
		}
		return next(ctx)
	}
	return a
}

func ExecRemove(s *usersystem.UserSystem, uid string) error {
	ctx := usersystem.UIDContext(s.Context, uid)
	_, err := s.System.ExecActions(ctx, CommandRemove)
	return err
}
