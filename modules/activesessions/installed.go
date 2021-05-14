package activesessions

import "github.com/herb-go/usersystem"

type InstalledAcitveSessions struct {
	*ActiveSessions
	UserSystem *usersystem.UserSystem
}

func (s *InstalledAcitveSessions) MustPurgeActiveSession(session *usersystem.Session) {
	if session == nil {
		return
	}
	s.Service.MustPurgeActiveSession(session.Type, session.UID(), session.Payloads.LoadString(PayloadSerialNumber))
}

func NewInstalledAcitveSessions() *InstalledAcitveSessions {
	return &InstalledAcitveSessions{}
}
