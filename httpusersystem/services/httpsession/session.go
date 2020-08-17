package httpsession

import (
	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

const SessionKeyPrefix = "."

type RequestSession interface {
	Set(name string, v interface{}) error
	Get(name string, v interface{}) error
	Del(name string) error
	Destory() error
	IsNotFoundError()
}
type Session struct {
	id          string
	sessionType usersystem.SessionType
	session     RequestSession
}

func (s *Session) ID() string {
	return s.id
}
func (s *Session) Type() usersystem.SessionType {
	return s.sessionType
}
func (s *Session) UID() (string, error) {
	id := ""
	err := s.session.Get("uid", &id)
	if err != nil {
		return "", err
	}
	return id, nil
}
func (s *Session) SaveUID(id string) error {
	return s.session.Set("uid", id)
}
func (s *Session) Payloads() (*authority.Payloads, error) {
	payload := authority.NewPayloads()
	err := s.session.Get("payloads", &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
func (s *Session) SavePayloads(payload *authority.Payloads) error {
	return s.session.Set("payloads", payload)

}
func (s *Session) Destory() error {
	return s.Destory()
}
func (s *Session) Save(key string, v interface{}) error {
	return s.session.Set(SessionKeyPrefix+key, v)
}
func (s *Session) Load(key string, v interface{}) error {
	return s.session.Get(SessionKeyPrefix+key, v)

}
func (s *Session) Remove(key string) error {
	return s.session.Del(SessionKeyPrefix + key)
}
func (s *Session) IsNotFoundError(err error) bool {
	return s.IsNotFoundError(err)
}

type HTTPSession struct {
	herbsystem.NopService
	Services map[usersystem.SessionType]Service
}
