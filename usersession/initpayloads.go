package usersession

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandInitPayloads = herbsystem.Command("initpayloads")

func WrapInitPayloads(h func(context.Context, usersystem.SessionType, string, *authority.Payloads)) *herbsystem.Action {
	return herbsystem.CreateAction(CommandInitPayloads, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		h(ctx, usersystem.GetSessionType(ctx), usersystem.GetUID(ctx), usersystem.GetPayloads(ctx))
		next(ctx, system)
	})
}

func MustExecInitPayloads(s *usersystem.UserSystem, ctx context.Context, st usersystem.SessionType, uid string) *authority.Payloads {
	payloads := authority.NewPayloads()
	payloads.Set(usersystem.PayloadUID, []byte(uid))
	ctx = usersystem.UIDContext(ctx, uid)
	ctx = context.WithValue(ctx, usersystem.ContextKeyPayloads, payloads)
	ctx = usersystem.SessionTypeContext(ctx, st)
	ctx = herbsystem.MustExecActions(ctx, s, CommandInitPayloads)
	return usersystem.GetPayloads(ctx)
}
