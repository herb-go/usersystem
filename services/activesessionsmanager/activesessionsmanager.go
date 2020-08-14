package activesessionsmanager

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

var ServiceName = "activesessionsmanager"

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
