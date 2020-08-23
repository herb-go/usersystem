package httpsession

import (
	"net/http"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type Service interface {
	usersystem.SessionStore
	SessionMiddleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	GetRequestSession(r *http.Request, st usersystem.SessionType) (*usersystem.Session, error)
	LoginRequestSession(r *http.Request, payloads *authority.Payloads) error
	LogoutRequestSession(r *http.Request) (bool, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
