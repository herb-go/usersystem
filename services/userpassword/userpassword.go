package userpassword

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ServiceName = "password"

type UserPassword struct {
	herbsystem.NopService
	Service
}

func New() *UserPassword {
	return &UserPassword{}
}

func (s *UserPassword) InitService() error {
	return nil
}
func (s *UserPassword) ServiceName() string {
	return ServiceName
}
func (s *UserPassword) StartService() error {
	return s.Service.Start()
}
func (s *UserPassword) StopService() error {
	return s.Service.Stop()
}
func (s *UserPassword) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userpurge.Wrap(s),
	}
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *UserPassword {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
}
