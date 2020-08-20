package httpsession

import (
	"net/http"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

const SessionKeyPrefix = "."

var ServiceName = "httpsession"

var SessionType = usersystem.SessionType("http")

type HTTPSession struct {
	herbsystem.NopService
	Service
	Name string
	Type usersystem.SessionType
}

func (s *HTTPSession) InitService() error {
	return nil
}
func (s *HTTPSession) ServiceName() string {
	return s.Name
}
func (s *HTTPSession) StartService() error {
	return s.Service.Start()
}
func (s *HTTPSession) StopService() error {
	return s.Service.Stop()
}
func (s *HTTPSession) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapGetSession(func(st usersystem.SessionType, id string) (usersystem.Session, error) {
			if st != s.Type {
				return nil, nil
			}
			return s.Service.GetSession(id, s.Type)
		}),
	}
}
func (s *HTTPSession) GetRequestSession(r *http.Request) (usersystem.Session, error) {
	return s.Service.GetRequestSession(r, s.Type)
}
func New() *HTTPSession {
	return &HTTPSession{}
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *HTTPSession {
	session := New()
	session.Name = ServiceName
	session.Type = SessionType
	err := s.InstallService(session)
	if err != nil {
		panic(err)
	}
	return session
}
