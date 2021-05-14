package userterm

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"
)

var ModuleName = "userterm"

var PayloadTerm = "term"

type UserTerm struct {
	herbsystem.NopModule
	Service
}

func New() *UserTerm {
	return &UserTerm{}
}

func (s *UserTerm) ModuleName() string {
	return ModuleName
}
func (s *UserTerm) StartService() error {
	return s.Service.Start()
}
func (s *UserTerm) StopService() error {
	return s.Service.Stop()
}
func (s *UserTerm) MustInitPayloads(ctx context.Context, st usersystem.SessionType, uid string, payloads *authority.Payloads) {
	term := s.Service.MustCurrentTerm(uid)
	payloads.Set(PayloadTerm, []byte(term))
}
func (s *UserTerm) MustCheckSession(ctx context.Context, session *usersystem.Session) bool {
	uid := session.UID()
	if uid == "" {
		return false
	}
	term := s.Service.MustCurrentTerm(uid)
	sessionterm := session.Payloads.LoadString(PayloadTerm)
	return term == sessionterm
}
func (s *UserTerm) InitProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.WrapStartOrPanicAction(s.StartService),
		herbsystem.WrapStopOrPanicAction(s.StopService),
	)
	system.MountSystemActions(
		userpurge.Wrap(s),
		usersession.WrapInitPayloads(s.MustInitPayloads),
		usersession.WrapCheckSession(s.MustCheckSession),
	)
	next(ctx, system)
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserTerm {
	term := New()
	s.MustRegisterSystemModule(term)
	i := NewInstalledUserTerm()
	i.UserTerm = term
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *UserTerm {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*UserTerm)
}
