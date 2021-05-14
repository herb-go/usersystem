package userpassword

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ModuleName = "password"

type UserPassword struct {
	herbsystem.NopModule
	Service
}

func New() *UserPassword {
	return &UserPassword{}
}

func (s *UserPassword) ModuleName() string {
	return ModuleName
}
func (s *UserPassword) StartService() error {
	return s.Service.Start()
}
func (s *UserPassword) StopService() error {
	return s.Service.Stop()
}
func (s *UserPassword) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),
		userpurge.Wrap(s),
	)
	next(ctx, system)
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserPassword {
	password := New()
	s.MustRegisterSystemModule(password)
	i := NewInstalledUserPassword()
	i.UserPassword = password
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *UserPassword {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*UserPassword)
}
