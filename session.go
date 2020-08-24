package usersystem

import (
	"github.com/herb-go/herbsecurity/authority"
)

type SessionType string

type Session struct {
	ID       string
	Type     SessionType
	Payloads *authority.Payloads
}

func (s *Session) WithID(id string) *Session {
	s.ID = id
	return s
}
func (s *Session) WithType(t SessionType) *Session {
	s.Type = t
	return s
}
func (s *Session) WithPayloads(p *authority.Payloads) *Session {
	s.Payloads = p
	return s
}
func (s *Session) UID() string {
	return s.Payloads.LoadString(PayloadUID)
}
func (s *Session) RevokeCode() string {
	return s.Payloads.LoadString(PayloadRevokeCode)
}
func NewSession() *Session {
	return &Session{}
}

var PayloadUID = "uid"
var PayloadRevokeCode = "revokecode"

type SessionStore interface {
	GetSession(sessiontype SessionType, id string) (*Session, error)
	RevokeSession(sessiontype SessionType, code string) (bool, error)
}
