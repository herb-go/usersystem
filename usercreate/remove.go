package usercreate

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandRemove = herbsystem.Command("remove")

func WrapRemove(h func(string)) *herbsystem.Action {
	return herbsystem.CreateAction(CommandRemove, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		h(usersystem.GetUID(ctx))
		next(ctx, system)
	})
}

func MustExecRemove(s *usersystem.UserSystem, uid string) {
	ctx := usersystem.UIDContext(s.SystemContext(), uid)
	herbsystem.MustExecActions(ctx, s, CommandRemove)
}
