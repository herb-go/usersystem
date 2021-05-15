package userstatus

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usercreate"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"
)

var ModuleName = "status"

type UserStatus struct {
	herbsystem.NopModule
	Service
}

func New() *UserStatus {
	return &UserStatus{}
}

func (s *UserStatus) ModuleName() string {
	return ModuleName
}
func (s *UserStatus) StartService() error {
	return s.Service.Start()
}
func (s *UserStatus) StopService() error {
	return s.Service.Stop()
}
func (s *UserStatus) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),
	)
	system.MountSystemActions(
		usersession.WrapCheckSession(s.MustCheckSession),
		userpurge.Wrap(s),
		usercreate.WrapExist(func(id string) bool {
			_, found := s.Service.MustLoadStatus(id)
			return found
		}),
		usercreate.WrapCreate(func(id string) {
			s.Service.MustCreateStatus(id)
		}, func(id string) {
			s.Service.MustRemoveStatus(id)

		}),
		usercreate.WrapRemove(func(id string) {
			s.Service.MustRemoveStatus(id)
		}),
	)
	next(ctx, system)
}
func (s *UserStatus) MustCheckSession(ctx context.Context, session *usersystem.Session) bool {
	id := session.UID()
	if id == "" {
		return false
	}
	return s.MustIsUserAvaliable(id)
}

func (s *UserStatus) MustIsUserAvaliable(id string) bool {
	st, found := s.Service.MustLoadStatus(id)
	if !found {
		return false
	}
	ok, err := s.Service.IsAvailable(st)
	if err != nil {
		panic(err)
	}
	return ok
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserStatus {
	status := New()
	s.MustRegisterSystemModule(status)
	i := NewInstalledUserStatus()
	i.UserStatus = status
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *UserStatus {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*UserStatus)
}
