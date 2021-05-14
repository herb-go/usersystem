package usercreate

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandExist = herbsystem.Command("exist")

type contextkey string

var contextkeyResult = contextkey("result")

type result struct {
	exists bool
}

func WrapExist(h func(string) bool) *herbsystem.Action {
	return herbsystem.CreateAction(CommandExist, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		ok := h(usersystem.GetUID(ctx))
		if ok {
			r := ctx.Value(contextkeyResult).(*result)
			r.exists = true
			return
		}
		next(ctx, system)
	})
}

func MustExecExist(s *usersystem.UserSystem, uid string) bool {
	r := &result{exists: false}
	ctx := usersystem.UIDContext(s.SystemContext(), uid)
	ctx = context.WithValue(ctx, contextkeyResult, r)
	herbsystem.MustExecActions(ctx, s, CommandExist)
	return r.exists
}
