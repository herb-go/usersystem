package userstatus

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/reloaduser"
)

var ServiceName = "status"

type UserStauts struct {
	herbsystem.NopService
	StatusService
}

func New() *UserStauts {
	return &UserStauts{}
}
func (s *UserStauts) InitService() error {
	return nil
}
func (s *UserStauts) ServiceName() string {
	return ServiceName
}
func (s *UserStauts) StartService() error {
	return s.Start()
}
func (s *UserStauts) StopService() error {
	return s.Stop()
}
func (s *UserStauts) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		reloaduser.Wrap(func(uid string) error {
			return s.Reload(uid)
		}),
	}
}

func MustNewUserstatus(s *usersystem.UserSystem) *UserStauts {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
}
