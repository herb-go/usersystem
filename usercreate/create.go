package usercreate

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandCreate = herbsystem.Command("create")

func WrapCreate(h func(string), onFail func(string)) *herbsystem.Action {
	return herbsystem.CreateAction(CommandCreate, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		uid := usersystem.GetUID(ctx)
		defer func() {
			if onFail != nil && !herbsystem.IsFinished(ctx) {
				onFail(uid)
			}
		}()
		h(uid)
		next(ctx, system)
	})
}

func MustExecCreate(s *usersystem.UserSystem, uid string) {
	ctx := usersystem.UIDContext(s.SystemContext(), uid)
	herbsystem.MustExecActions(ctx, s, CommandCreate)
}
