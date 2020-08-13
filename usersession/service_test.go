package usersession

import (
	"time"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var payloads = authority.NewPayloads()

var lastactive = ""

type testSession string

func (s testSession) ID() string {
	return ""
}
func (s testSession) Type() usersystem.SessionType {
	return "test"
}
func (s testSession) UID() (string, error) {
	return string(s), nil
}
func (s testSession) SaveUID(string) error {
	return nil
}
func (s testSession) Payloads() (*authority.Payloads, error) {
	return authority.NewPayloads(), nil
}
func (s testSession) SavePayloads(p *authority.Payloads) error {
	payloads = p
	return nil
}

func (s testSession) Destory() error {
	return nil
}
func (s testSession) Save(key string, v interface{}) error {
	return nil
}
func (s testSession) Load(key string, v interface{}) error {
	return nil
}
func (s testSession) Remove(key string) error {
	return nil
}
func (s testSession) IsNotFoundError(err error) bool {
	return false
}

type testService struct {
	herbsystem.NopService
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		WrapCheckSession(func(session usersystem.Session, id string, payloads *authority.Payloads) (bool, error) {
			return id == "exists", nil
		}),
		WrapInitPayloads(func(session usersystem.Session, id string, payloads *authority.Payloads) error {
			payloads.Set("test", []byte("testvalue"))
			return nil
		}),
		WrapOnSessionActive(func(session usersystem.Session, id string) error {
			lastactive = id
			return nil
		}),
		WrapActiveSessionManagerConfig(func(st usersystem.SessionType) (*Config, error) {
			if st != usersystem.SessionType("test") {
				return nil, nil
			}
			return &Config{
				Supported: true,
				Duration:  time.Minute,
			}, nil
		}),
		WrapGetActiveSessions(func(st usersystem.SessionType) ([]*ActiveSession, bool, error) {
			if st != usersystem.SessionType("test") {
				return nil, false, nil
			}
			return []*ActiveSession{
				&ActiveSession{
					SessionID: "active",
				},
			}, true, nil
		}),
	}
}
