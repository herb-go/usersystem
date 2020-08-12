package userpurge

import (
	"context"

	"github.com/herb-go/usersystem"

	"github.com/herb-go/herbsystem"
)

type Purgeable interface {
	Purge(string) error
}

var Command = herbsystem.Command("purge")

func Wrap(h Purgeable) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = Command
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h.Purge(usersystem.GetUID(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecPurge(s *usersystem.UserSystem, uid string) error {
	ctx := usersystem.UIDContext(s.Context, uid)
	_, err := s.System.ExecActions(ctx, Command)
	return err
}
