package userrole

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ModuleName = "role"

type UserRole struct {
	herbsystem.NopModule
	Service
}

func New() *UserRole {
	return &UserRole{}
}

func (s *UserRole) ModuleName() string {
	return ModuleName
}
func (s *UserRole) StartService() error {
	return s.Service.Start()
}
func (s *UserRole) StopService() error {
	return s.Service.Stop()
}
func (s *UserRole) InitProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),
	)
	system.MountSystemActions(
		userpurge.Wrap(s),
	)
	next(ctx, system)
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserRole {
	role := New()
	s.MustRegisterSystemModule(role)
	i := NewInstalledUserRole()
	i.UserRole = role
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *UserRole {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*UserRole)
}
