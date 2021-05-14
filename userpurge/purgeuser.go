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
	return herbsystem.CreateAction(Command, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		err := h.Purge(usersystem.GetUID(ctx))
		if err != nil {
			panic(err)
		}
		next(ctx, system)
	})
}

func MustExecPurge(s *usersystem.UserSystem, uid string) {
	ctx := usersystem.UIDContext(s.SystemContext(), uid)
	herbsystem.MustExecActions(ctx, s, Command)
}
