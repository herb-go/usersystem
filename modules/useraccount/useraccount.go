package useraccount

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ModuleName = "account"

type UserAccount struct {
	herbsystem.NopModule
	Service
}

func New() *UserAccount {
	return &UserAccount{}
}

func (s *UserAccount) ModuleName() string {
	return ModuleName
}
func (s *UserAccount) StartService() error {
	return s.Service.Start()
}
func (s *UserAccount) StopService() error {
	return s.Service.Stop()
}
func (s *UserAccount) InitProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),
		userpurge.Wrap(s),
	)
	next(ctx, system)
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserAccount {
	a := New()
	s.MustRegisterSystemModule(a)
	i := NewInstalledUserAccount()
	i.UserAccount = a
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *UserAccount {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*UserAccount)
}
