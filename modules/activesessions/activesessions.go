package activesessions

import (
	"context"
	"time"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

type Config struct {
	Supported bool
	Duration  time.Duration
}
type Active struct {
	SessionID  string
	LastActive time.Time
}

var ModuleName = "activesessionsmanager"
var PayloadSerialNumber = "sessionSerialNumber"

func MustGetSerialNumber(s *usersystem.Session) string {
	return s.Payloads.LoadString(PayloadSerialNumber)
}

type ActiveSessions struct {
	herbsystem.NopModule
	Service
}

func (s *ActiveSessions) ModuleName() string {
	return ModuleName
}
func (s *ActiveSessions) StartService() error {
	return s.Service.Start()
}
func (s *ActiveSessions) StopService() error {
	return s.Service.Stop()
}

func (s *ActiveSessions) MustOnSessionActive(session *usersystem.Session) {
	s.Service.MustOnSessionActive(session)
}
func (s *ActiveSessions) InitProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),
		usersession.WrapOnSessionActive(s.MustOnSessionActive),
		usersession.WrapInitPayloads(func(ctx context.Context, st usersystem.SessionType, uid string, p *authority.Payloads) {
			serialnumber := s.Service.MustCreateSerialNumber()
			if serialnumber != "" {
				p.Set(PayloadSerialNumber, []byte(serialnumber))
			}

		}),
	)
	next(ctx, system)
}
func New() *ActiveSessions {
	return &ActiveSessions{}
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledAcitveSessions {
	as := New()
	s.MustRegisterSystemModule(as)
	i := NewInstalledAcitveSessions()
	i.ActiveSessions = as
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *ActiveSessions {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*ActiveSessions)
}
