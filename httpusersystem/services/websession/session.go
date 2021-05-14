package websession

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

const SessionKeyPrefix = "."

var ModuleName = "websession"

var SessionType = usersystem.SessionType("web")

type ContextSession struct {
	Session *usersystem.Session
}

type WebSession struct {
	herbsystem.NopModule
	Service
	Name       string
	UserSystem *usersystem.UserSystem
	Type       usersystem.SessionType
}

func (s *WebSession) ModuleName() string {
	return s.Name
}
func (s *WebSession) StartService() error {
	return s.Service.Start()
}
func (s *WebSession) StopService() error {
	return s.Service.Stop()
}
func (s *WebSession) MustGetSession(id string) *usersystem.Session {
	return s.Service.MustGetSession(s.Type, id)
}

func (s *WebSession) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),

		usersession.WrapGetSession(func(st usersystem.SessionType, id string) *usersystem.Session {
			if st != s.Type {
				return nil
			}
			return s.Service.MustGetSession(s.Type, id)

		}),
		usersession.WrapRevokeSession(func(st usersystem.SessionType, code string) bool {
			if st != s.Type {
				return false
			}
			return s.Service.MustRevokeSession(code)
		}),
	)
	next(ctx, system)
}

func New() *WebSession {
	return &WebSession{}
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledWebSession {
	session := New()
	session.Name = ModuleName
	session.Type = SessionType
	session.UserSystem = s
	s.MustRegisterSystemModule(session)

	i := NewInstalledWebSession()
	i.WebSession = session
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *WebSession {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*WebSession)
}
