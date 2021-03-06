package useraccount

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ServiceName = "account"

type UserAccount struct {
	herbsystem.NopService
	Service
}

func New() *UserAccount {
	return &UserAccount{}
}

func (s *UserAccount) InitService() error {
	return nil
}
func (s *UserAccount) ServiceName() string {
	return ServiceName
}
func (s *UserAccount) StartService() error {
	return s.Service.Start()
}
func (s *UserAccount) StopService() error {
	return s.Service.Stop()
}
func (s *UserAccount) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userpurge.Wrap(s),
	}
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserAccount {
	a := New()
	err := s.InstallService(a)
	if err != nil {
		panic(err)
	}
	i := NewInstalledUserAccount()
	i.UserAccount = a
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*UserAccount, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*UserAccount), nil
}
