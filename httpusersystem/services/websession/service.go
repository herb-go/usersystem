package websession

import (
	"net/http"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type Service interface {
	usersystem.SessionStore
	SessionMiddleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
	MustGetRequestSession(r *http.Request, st usersystem.SessionType) *usersystem.Session
	MustLoginRequestSession(r *http.Request, payloads *authority.Payloads) *usersystem.Session
	MustLogoutRequestSession(r *http.Request) bool
	// Set set session by field name with given value.
	Set(r *http.Request, fieldname string, v interface{}) error
	//Get get session by field name with given value.
	Get(r *http.Request, fieldname string, v interface{}) error
	// Del del session value by field name .
	Del(r *http.Request, fieldname string) error
	// IsNotFoundError check if given error is session not found error.
	IsNotFoundError(err error) bool
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
