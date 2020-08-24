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

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserPassword {
	password := New()
	err := s.InstallService(password)
	if err != nil {
		panic(err)
	}
	i := NewInstalledUserPassword()
	i.UserPassword = password
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*UserPassword, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*UserPassword), nil
}
