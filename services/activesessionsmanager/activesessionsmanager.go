package activesessionsmanager

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

var ServiceName = "activesessionsmanager"
var PayloadSerialNumber = "sessionSerialNumber"

func GetSerialNumber(s *usersystem.Session) (string, error) {
	return s.Payloads.LoadString(PayloadSerialNumber), nil
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
func (m *Manager) OnSessionActive(session *usersystem.Session) error {
	return m.Service.OnSessionActive(session)
}
func (m *Manager) GetActiveSessions(st usersystem.SessionType, id string) ([]*usersession.ActiveSession, bool, error) {
	return m.Service.GetActiveSessions(st, id)
}
func (m *Manager) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapOnSessionActive(m.OnSessionActive),
		usersession.WrapActiveSessionManagerConfig(m.Config),
		usersession.WrapGetActiveSessions(m.GetActiveSessions),
		usersession.WrapInitPayloads(func(ctx context.Context, st usersystem.SessionType, uid string, p *authority.Payloads) error {
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
