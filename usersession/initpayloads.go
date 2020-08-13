package usersession

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandInitPayloads = herbsystem.Command("initpayloads")

func WrapInitPayloads(h func(usersystem.Session, string, *authority.Payloads) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandInitPayloads
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(usersystem.GetSession(ctx), usersystem.GetUID(ctx), GetPayloads(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecInitPayloads(s *usersystem.UserSystem, session usersystem.Session) error {
	uid, err := session.UID()
	if err != nil {
		return err
	}
	payloads, err := session.Payloads()
	if err != nil {
		return err
	}
	ctx := usersystem.SessionContext(s.Context, session)
	ctx = usersystem.UIDContext(ctx, uid)
	ctx = context.WithValue(ctx, ContextKeyPayloads, payloads)
	_, err = s.System.ExecActions(ctx, CommandInitPayloads)
	if err != nil {
		return err
	}
	return nil
}