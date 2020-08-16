package userstatus

import (
	"github.com/herb-go/herb/user"
	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"
)

var ServiceName = "status"

type UserStatus struct {
	herbsystem.NopService
	Service
}

func New() *UserStatus {
	return &UserStatus{}
}
func (s *UserStatus) InitService() error {
	return nil
}
func (s *UserStatus) ServiceName() string {
	return ServiceName
}
func (s *UserStatus) StartService() error {
	return s.Service.Start()
}
func (s *UserStatus) StopService() error {
	return s.Service.Stop()
}
func (s *UserStatus) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapCheckSession(s.CheckSession),
		userpurge.Wrap(s),
	}
}
func (s *UserStatus) CheckSession(session usersystem.Session, id string, payloads *authority.Payloads) (bool, error) {
	return s.IsUserAvaliable(id)
}
func (s *UserStatus) IsUserAvaliable(id string) (bool, error) {
	st, err := s.Service.LoadStatus(id)
	if err != nil {
		if err == user.ErrUserNotExists {
			return false, nil
		}
		return false, err
	}
	return s.Service.IsAvailable(st)
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *UserStatus {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
}
