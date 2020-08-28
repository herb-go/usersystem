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

var ServiceName = "activesessionsmanager"
var PayloadSerialNumber = "sessionSerialNumber"

func GetSerialNumber(s *usersystem.Session) (string, error) {
	return s.Payloads.LoadString(PayloadSerialNumber), nil
}

type ActiveSessions struct {
	herbsystem.NopService
	Service
}

func (s *ActiveSessions) InitService() error {
	return nil
}
func (s *ActiveSessions) ServiceName() string {
	return ServiceName
}
func (s *ActiveSessions) StartService() error {
	return s.Service.Start()
}
func (s *ActiveSessions) StopService() error {
	return s.Service.Stop()
}

func (s *ActiveSessions) OnSessionActive(session *usersystem.Session) error {
	return s.Service.OnSessionActive(session)
}
func (s *ActiveSessions) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapOnSessionActive(s.OnSessionActive),
		usersession.WrapInitPayloads(func(ctx context.Context, st usersystem.SessionType, uid string, p *authority.Payloads) error {
			serialnumber, err := s.Service.CreateSerialNumber()
			if err != nil {
				return err
			}
			p.Set(PayloadSerialNumber, []byte(serialnumber))
			return nil
		}),
	}
}
func New() *ActiveSessions {
	return &ActiveSessions{}
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledAcitveSessions {
	as := New()
	err := s.InstallService(as)
	if err != nil {
		panic(err)
	}
	i := NewInstalledAcitveSessions()
	i.ActiveSessions = as
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*ActiveSessions, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*ActiveSessions), nil
}
