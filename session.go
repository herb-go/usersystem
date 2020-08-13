package usersystem

import (
	"github.com/herb-go/herbsecurity/authority"
)

type SessionType string

type Session interface {
	ID() string
	Type() SessionType
	UID() (string, error)
	Payloads() (*authority.Payloads, error)
	Destory() error
	Save(key string, v interface{}) error
	Load(key string, v interface{}) error
	Remove(key string) error
	IsNotFoundError(err error) bool
}
