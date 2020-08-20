package activesessionsmanager

import (
	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

var ServiceName = "activesessionsmanager"
var PayloadSerialNumber = "sessionSerialNumber"

func GetSerialNumber(s usersystem.Session) (string, error) {
	p, err := s.Payloads()
	if err != nil {
		return "", err
	}
	return p.LoadString(PayloadSerialNumber), nil
}

type Manager struct {
	herbsystem.NopService
	Service
}

func (m *Manager) InitService() error {
	return nil
}
func (m *Manager) ServiceName() string {
	return ServiceName
}
func (m *Manager) StartService() error {
	return m.Service.Start()
}
func (m *Manager) StopService() error {
	return m.Service.Stop()
}
func (m *Manager) Config(st usersystem.SessionType) (*usersession.Config, error) {
	return m.Service.Config(st)
}
func (m *Manager) OnSessionActive(session usersystem.Session, uid string) error {
	return m.Service.OnSessionActive(session, uid)
}
func (m *Manager) GetActiveSessions(st usersystem.SessionType) ([]*usersession.ActiveSession, bool, error) {
	return m.Service.GetActiveSessions(st)
}
func (m *Manager) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapOnSessionActive(m.OnSessionActive),
		usersession.WrapActiveSessionManagerConfig(m.Config),
		usersession.WrapGetActiveSessions(m.GetActiveSessions),
		usersession.WrapInitPayloads(func(s usersystem.Session, id string, p *authority.Payloads) error {
			serialnumber, err := m.Service.CreateSerialNumber()
			if err != nil {
				return err
			}
			p.Set(PayloadSerialNumber, []byte(serialnumber))
			return nil
		}),
	}
}
func New() *Manager {
	return &Manager{}
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *Manager {
	m := New()
	err := s.InstallService(m)
	if err != nil {
		panic(err)
	}
	return m
}

func GetService(s *usersystem.UserSystem) (*Manager, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*Manager), nil
}
