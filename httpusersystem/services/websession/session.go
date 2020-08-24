package websession

import (
	"net/http"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

const SessionKeyPrefix = "."

var ServiceName = "websession"

var SessionType = usersystem.SessionType("web")

type WebSession struct {
	herbsystem.NopService
	Service
	Name       string
	UserSystem *usersystem.UserSystem
	Type       usersystem.SessionType
}

func (s *WebSession) InitService() error {
	return nil
}
func (s *WebSession) ServiceName() string {
	return s.Name
}
func (s *WebSession) StartService() error {
	return s.Service.Start()
}
func (s *WebSession) StopService() error {
	return s.Service.Stop()
}

func (s *WebSession) GetRequestSession(r *http.Request) (*usersystem.Session, error) {
	return s.Service.GetRequestSession(r, s.Type)
}
func (s *WebSession) Logout(r *http.Request) (bool, error) {
	return s.Service.LogoutRequestSession(r)
}
func (s *WebSession) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapGetSession(func(st usersystem.SessionType, id string) (*usersystem.Session, error) {
			if st != s.Type {
				return nil, nil
			}
			return s.Service.GetSession(s.Type, id)
		}),
		usersession.WrapRevokeSession(func(st usersystem.SessionType, code string) (bool, error) {
			if st != s.Type {
				return false, nil
			}
			return s.Service.RevokeSession(st, code)
		}),
	}
}
func New() *WebSession {
	return &WebSession{}
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledWebSession {
	session := New()
	session.Name = ServiceName
	session.Type = SessionType
	session.UserSystem = s
	err := s.InstallService(session)
	if err != nil {
		panic(err)
	}
	i := NewInstalledWebSession()
	i.WebSession = session
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*WebSession, error) {
	return GetServiceByName(s, ServiceName)
}

func GetServiceByName(s *usersystem.UserSystem, servicename string) (*WebSession, error) {
	v, err := s.GetConfigurableService(servicename)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*WebSession), nil
}
