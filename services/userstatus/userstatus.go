package userstatus

import (
	"github.com/herb-go/herb/user"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usercreate"
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
		usercreate.WrapExist(func(id string) (bool, error) {
			_, err := s.Service.LoadStatus(id)
			if err != nil {
				if err == user.ErrUserNotExists {
					return false, nil
				}
				return false, err
			}
			return true, nil
		}),
		usercreate.WrapCreate(func(id string, next func() error) error {
			err := s.Service.CreateStatus(id)
			if err != nil {
				return err
			}
			err = next()
			if err != nil {
				s.Service.RemoveStatus(id)
				return err
			}
			return nil
		}),
		usercreate.WrapRemove(func(id string) error {
			return s.Service.RemoveStatus(id)
		}),
	}
}
func (s *UserStatus) CheckSession(session *usersystem.Session) (bool, error) {
	id := session.UID()
	if id == "" {
		return false, nil
	}
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

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserStatus {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	i := NewInstalledUserStatus()
	i.UserStatus = status
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*UserStatus, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*UserStatus), nil
}
