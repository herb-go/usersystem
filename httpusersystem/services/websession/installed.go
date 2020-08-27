package websession

import (
	"net/http"

	"github.com/herb-go/usersystem/httpusersystem"
	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/usersystem"
)

type InstalledWebSession struct {
	*WebSession
	UserSystem *usersystem.UserSystem
}

func NewInstalledWebSession() *InstalledWebSession {
	return &InstalledWebSession{}
}

func (s *InstalledWebSession) Middleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
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

func (s *InstalledWebSession) Login(r *http.Request, uid string) (*usersystem.Session, error) {
	ctx := httpusersystem.RequestContext(s.UserSystem.Context, r)
	p, err := usersession.ExecInitPayloads(s.UserSystem, ctx, s.Type, uid)
	if err != nil {
		return nil, err
	}
	return s.Service.LoginRequestSession(r, p)
}

func (s *InstalledWebSession) IdentifyRequest(r *http.Request) (uid string, err error) {
	session, err := s.GetRequestSession(r)
	if err != nil {
		return "", err
	}
	if session == nil {
		return "", nil
	}
	ok, err := usersession.ExecCheckSession(s.UserSystem, session)
	if !ok {
		return "", nil
	}
	return session.UID(), nil
}
