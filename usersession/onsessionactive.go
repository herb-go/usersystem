package usersession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandOnSessionActive = herbsystem.Command("onsessionactive")

func WrapOnSessionActive(h func(*usersystem.Session)) *herbsystem.Action {
	return herbsystem.CreateAction(CommandOnSessionActive, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		h(usersystem.GetSession(ctx))
		next(ctx, system)
	})
}

func MustExecOnSessionActive(s *usersystem.UserSystem, session *usersystem.Session) {
	ctx := usersystem.SessionContext(s.SystemContext(), session)
	herbsystem.MustExecActions(ctx, s, CommandOnSessionActive)
}
