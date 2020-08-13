package usersession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandOnSessionActive = herbsystem.Command("onsessionactive")

func WrapOnSessionActive(h func(usersystem.Session, string) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandOnSessionActive
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(usersystem.GetSession(ctx), usersystem.GetUID(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecOnSessionActive(s *usersystem.UserSystem, session usersystem.Session) error {
	uid, err := session.UID()
	if err != nil {
		return err
	}
	ctx := usersystem.SessionContext(s.Context, session)
	ctx = usersystem.UIDContext(ctx, uid)
	_, err = s.System.ExecActions(ctx, CommandOnSessionActive)
	if err != nil {
		return err
	}
	return nil
}
