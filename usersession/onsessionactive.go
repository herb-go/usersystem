package usersession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandOnSessionActive = herbsystem.Command("onsessionactive")

func WrapOnSessionActive(h func(*usersystem.Session) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandOnSessionActive
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(usersystem.GetSession(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecOnSessionActive(s *usersystem.UserSystem, session *usersystem.Session) error {
	ctx := usersystem.SessionContext(s.Context, session)
	_, err := s.System.ExecActions(ctx, CommandOnSessionActive)
	return err
}
