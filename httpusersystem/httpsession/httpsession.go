package httpsession

import (
	"context"
	"net/http"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/httpusersystem"
)

var Command = herbsystem.Command("httpsession")

func Wrap(h func(*http.Request) (usersystem.Session, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = Command
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		session, err := h(httpusersystem.GetRequest(ctx))
		if err != nil {
			return err
		}
		ctx = usersystem.SessionContext(ctx, session)
		return next(ctx)
	}
	return a
}

func ExecSession(s *usersystem.UserSystem, req *http.Request) (usersystem.Session, error) {
	ctx := httpusersystem.RequestContext(s.Context, req)
	ctx = usersystem.SessionContext(ctx, nil)
	ctx, err := s.System.ExecActions(ctx, Command)
	if err != nil {
		return nil, err
	}
	return usersystem.GetSession(ctx), nil
}
