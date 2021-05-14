package activesessions

import (
	"github.com/herb-go/usersystem"
)

type Service interface {
	MustConfig(st usersystem.SessionType) *Config
	MustOnSessionActive(session *usersystem.Session)
	MustGetActiveSessions(usersystem.SessionType, string) []*Active
	MustPurgeActiveSession(st usersystem.SessionType, uid string, serialnumber string)
	MustCreateSerialNumber() string
	Start() error
	Stop() error
}
