package usersession

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandInitPayloads = herbsystem.Command("initpayloads")

func WrapInitPayloads(h func(context.Context, usersystem.SessionType, string, *authority.Payloads) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandInitPayloads
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(ctx, usersystem.GetSessionType(ctx), usersystem.GetUID(ctx), usersystem.GetPayloads(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecInitPayloads(s *usersystem.UserSystem, ctx context.Context, st usersystem.SessionType, uid string) (*authority.Payloads, error) {
	var err error
	payloads := authority.NewPayloads()
	payloads.Set(usersystem.PayloadUID, []byte(uid))
	ctx = usersystem.UIDContext(ctx, uid)
	ctx = context.WithValue(ctx, usersystem.ContextKeyPayloads, payloads)
	ctx = usersystem.SessionTypeContext(ctx, st)
	ctx, err = s.System.ExecActions(ctx, CommandInitPayloads)
	if err != nil {
		return nil, err
	}
	return usersystem.GetPayloads(ctx), nil
}
