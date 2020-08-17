package httpsession

import (
	"net/http"

	"github.com/herb-go/usersystem"
)

type Service interface {
	SessionType() usersystem.SessionType
	GetSession(id string) (RequestSession, error)
	Middleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
