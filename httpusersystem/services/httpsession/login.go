package httpsession

import (
	"context"
	"net/http"

	"github.com/herb-go/usersystem/httpusersystem"
	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandLogin = herbsystem.Command("httpsession.login")

func WrapLogin(h func(st usersystem.SessionType, r *http.Request, p *authority.Payloads) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandLogin
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(usersystem.GetSessionType(ctx), httpusersystem.GetRequest(ctx), usersystem.GetPayloads(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func ExecLoginByType(s *usersystem.UserSystem, st usersystem.SessionType, uid string, r *http.Request) (*usersystem.Session, error) {
	var err error
	ctx := httpusersystem.RequestContext(s.Context, r)
	ctx = usersystem.UIDContext(ctx, uid)
	p, err := usersession.ExecInitPayloads(s, ctx, st, uid)
	if err != nil {
		return nil, err
	}
	ctx = usersystem.SessionTypeContext(s.Context, st)
	ctx = httpusersystem.RequestContext(ctx, r)
	ctx = usersystem.PayloadsContext(ctx, p)
	_, err = s.System.ExecActions(ctx, CommandLogin)
	if err != nil {
		return nil, err
	}
	session := usersystem.NewSession()
	session.WithType(st).WithPayloads(p)
	return session, nil
}

func ExecLogin(s *usersystem.UserSystem, uid string, r *http.Request) (*usersystem.Session, error) {
	return ExecLoginByType(s, SessionType, uid, r)
}
