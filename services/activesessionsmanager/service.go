package activesessionsmanager

import (
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

type Service interface {
	Config(st usersystem.SessionType) (*usersession.Config, error)
	OnSessionActive(session *usersystem.Session) error
	GetActiveSessions(usersystem.SessionType, string) ([]*usersession.ActiveSession, bool, error)
	CreateSerialNumber() (string, error)
	Start() error
	Stop() error
}
