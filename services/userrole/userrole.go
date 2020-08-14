package userrole

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ServiceName = "role"

type UserRole struct {
	herbsystem.NopService
	Service
}

func New() *UserRole {
	return &UserRole{}
}

func (s *UserRole) InitService() error {
	return nil
}
func (s *UserRole) ServiceName() string {
	return ServiceName
}
func (s *UserRole) StartService() error {
	return s.Service.Start()
}
func (s *UserRole) StopService() error {
	return s.Service.Stop()
}
func (s *UserRole) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userpurge.Wrap(s),
	}
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *UserRole {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
}
