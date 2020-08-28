package activesessions

import "github.com/herb-go/usersystem"

type InstalledAcitveSessions struct {
	*ActiveSessions
	UserSystem *usersystem.UserSystem
}

func (s *InstalledAcitveSessions) PurgeActiveSession(session *usersystem.Session) error {
	if session == nil {
		return nil
	}
	return s.Service.PurgeActiveSession(session.Type, session.UID(), session.Payloads.LoadString(PayloadSerialNumber))
}

func NewInstalledAcitveSessions() *InstalledAcitveSessions {
	return &InstalledAcitveSessions{}
}
