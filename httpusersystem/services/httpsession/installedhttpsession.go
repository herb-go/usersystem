package httpsession

import (
	"net/http"

	"github.com/herb-go/usersystem/httpusersystem"
	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/usersystem"
)

type InstalledHTTPSession struct {
	*HTTPSession
	UserSystem *usersystem.UserSystem
}

func NewInstalledHTTPSession() *InstalledHTTPSession {
	return &InstalledHTTPSession{}
}

func (s *InstalledHTTPSession) Middleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		s.Service.SessionMiddleware()(w, r, func(w http.ResponseWriter, r *http.Request) {
			session, err := s.GetRequestSession(r)
			if err != nil {
				panic(err)
			}
			err = usersession.ExecOnSessionActive(s.UserSystem, session)
			if err != nil {
				panic(err)
			}
			next(w, r)
		})
	}
}

func (s *InstalledHTTPSession) Login(r *http.Request, uid string) (*usersystem.Session, error) {
	ctx := httpusersystem.RequestContext(s.UserSystem.Context, r)
	p, err := usersession.ExecInitPayloads(s.UserSystem, ctx, s.Type, uid)
	if err != nil {
		return nil, err
	}
	return s.Service.LoginRequestSession(r, p)
}
