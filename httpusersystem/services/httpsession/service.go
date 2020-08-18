package httpsession

import (
	"net/http"

	"github.com/herb-go/usersystem"
)

type Service interface {
	GetSession(id string, st usersystem.SessionType) (usersystem.Session, error)
	SessionMiddleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	GetRequestSession(r *http.Request, st usersystem.SessionType) (usersystem.Session, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
