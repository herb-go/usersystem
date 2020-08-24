package activesessions

import (
	"github.com/herb-go/usersystem"
)

type Service interface {
	Config(st usersystem.SessionType) (*Config, error)
	OnSessionActive(session *usersystem.Session) error
	GetActiveSessions(usersystem.SessionType, string) ([]*Active, error)
	PurgeActiveSession(st usersystem.SessionType, uid string, serialnumber string) error
	CreateSerialNumber() (string, error)
	Start() error
	Stop() error
}
