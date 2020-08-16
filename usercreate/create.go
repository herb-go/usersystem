package usercreate

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandCreate = herbsystem.Command("create")

func WrapCreate(h func(string, func() error) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandCreate
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		nextregister := func() error {
			return next(ctx)
		}
		return h(usersystem.GetUID(ctx), nextregister)

	}
	return a
}

func ExecCreate(s *usersystem.UserSystem, uid string) error {
	ctx := usersystem.UIDContext(s.Context, uid)
	_, err := s.System.ExecActions(ctx, CommandCreate)
	return err
}
