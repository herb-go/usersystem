package activesessionsmanager

import (
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

type Service interface {
	Config(st usersystem.SessionType) (*usersession.Config, error)
	OnSessionActive(session usersystem.Session, uid string) error
	GetActiveSessions(usersystem.SessionType) ([]*usersession.ActiveSession, bool, error)
	Start() error
	Stop() error
}
