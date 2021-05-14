package websession

import (
	"context"
	"net/http"

	"github.com/herb-go/herbsystem"

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
			session := s.MustGetRequestSession(r)
			usersession.MustExecOnSessionActive(s.UserSystem, session)
			next(w, r)
		})
	}
}

func (s *InstalledWebSession) IdentifyRequest(r *http.Request) (string, error) {
	var uid string
	err := herbsystem.Catch(func() {
		uid = ""
		session := s.MustGetRequestSession(r)
		if session == nil {
			return
		}
		ok := usersession.MustExecCheckSession(s.UserSystem, session)
		if !ok {
			return
		}
		uid = session.UID()
	})
	if err != nil {
		return "", err
	}
	return uid, nil
}

func (s *InstalledWebSession) MustGetRequestSession(r *http.Request) *usersystem.Session {
	var cs *ContextSession
	v := r.Context().Value(s.Type)
	if v == nil {
		rs := s.Service.MustGetRequestSession(r, s.Type)

		cs = &ContextSession{
			Session: rs,
		}
		ctx := context.WithValue(r.Context(), s.Type, cs)
		req := r.WithContext(ctx)
		*r = *req
	} else {
		cs = v.(*ContextSession)
	}
	return cs.Session
}
func (s *InstalledWebSession) MustLogout(r *http.Request) bool {
	ok := s.Service.MustLogoutRequestSession(r)

	if ok {
		ctx := context.WithValue(r.Context(), s.Type, nil)
		req := r.WithContext(ctx)
		*r = *req
	}
	return ok
}

func (s *InstalledWebSession) MustLogin(r *http.Request, uid string) *usersystem.Session {
	ctx := httpusersystem.RequestContext(s.UserSystem.SystemContext(), r)
	p := usersession.MustExecInitPayloads(s.UserSystem, ctx, s.Type, uid)

	us := s.Service.MustLoginRequestSession(r, p)
	cs := &ContextSession{
		Session: us,
	}
	rctx := context.WithValue(r.Context(), s.Type, cs)
	req := r.WithContext(rctx)
	*r = *req
	return us
}

func (s *InstalledWebSession) LogoutMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	s.MustLogout(r)
	next(w, r)
}
