package websession

import (
	"context"
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

func (s *InstalledWebSession) GetRequestSession(r *http.Request) (*usersystem.Session, error) {
	var cs *ContextSession
	v := r.Context().Value(s.Type)
	if v == nil {
		rs, err := s.Service.GetRequestSession(r, s.Type)
		if err != nil {
			return nil, err
		}
		cs = &ContextSession{
			Session: rs,
		}
		ctx := context.WithValue(r.Context(), s.Type, cs)
		req := r.WithContext(ctx)
		*r = *req
	} else {
		cs = v.(*ContextSession)
	}
	return cs.Session, nil
}
func (s *InstalledWebSession) Logout(r *http.Request) (bool, error) {
	ok, err := s.Service.LogoutRequestSession(r)
	if err != nil {
		return false, err
	}
	if ok {
		ctx := context.WithValue(r.Context(), s.Type, nil)
		req := r.WithContext(ctx)
		*r = *req
	}
	return ok, nil
}

func (s *InstalledWebSession) Login(r *http.Request, uid string) (*usersystem.Session, error) {
	ctx := httpusersystem.RequestContext(s.UserSystem.Context, r)
	p, err := usersession.ExecInitPayloads(s.UserSystem, ctx, s.Type, uid)
	if err != nil {
		return nil, err
	}
	us, err := s.Service.LoginRequestSession(r, p)
	if err != nil {
		return nil, err
	}
	cs := &ContextSession{
		Session: us,
	}
	rctx := context.WithValue(r.Context(), s.Type, cs)
	req := r.WithContext(rctx)
	*r = *req
	return us, nil
}

func (s *InstalledWebSession) LogoutMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	_, err := s.Logout(r)
	if err != nil {
		panic(err)
	}
	next(w, r)
}
