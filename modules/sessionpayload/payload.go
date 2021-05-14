package sessionpayload

import (
	"context"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var ModuleName = "sessionpayload"

type SessionPayload struct {
	herbsystem.NopModule
	Builders []Builder
	Checkers []Checker
}

func New() *SessionPayload {
	return &SessionPayload{}
}

func (p *SessionPayload) AppendBuilder(b Builder) {
	p.Builders = append(p.Builders, b)
}
func (p *SessionPayload) AppendChecker(c Checker) {
	p.Checkers = append(p.Checkers, c)
}
func (p *SessionPayload) InitModule() {
	p.Builders = []Builder{}
	p.Checkers = []Checker{}
}

func (p *SessionPayload) ModuleName() string {
	return ModuleName
}
func (s *SessionPayload) StartProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	for _, v := range s.Builders {
		if err := v.Start(); err != nil {
			panic(err)
		}
	}
	for _, v := range s.Checkers {
		if err := v.Start(); err != nil {
			panic(err)
		}
	}
	next(ctx, system)
}
func (s *SessionPayload) StopProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	for _, v := range s.Builders {
		if err := v.Start(); err != nil {
			system.LogSystemError(err)
		}
	}
	for _, v := range s.Checkers {
		if err := v.Start(); err != nil {
			system.LogSystemError(err)
		}
	}
	next(ctx, system)
}

func (p *SessionPayload) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		herbsystem.CreateStartAction(p.StartProcess),
		herbsystem.CreateStopAction(p.StopProcess),
		usersession.WrapInitPayloads(p.MustBuildPayload),
		usersession.WrapCheckSession(p.MustCheckSession),
	)
	next(ctx, system)
}
func (p *SessionPayload) MustBuildPayload(ctx context.Context, st usersystem.SessionType, uid string, payloads *authority.Payloads) {
	for _, v := range p.Builders {
		v.MustBuildPayloads(ctx, st, uid, payloads)

	}
}
func (p *SessionPayload) MustCheckSession(ctx context.Context, session *usersystem.Session) bool {
	for _, v := range p.Checkers {
		ok := v.MustCheckSession(ctx, session)
		if !ok {
			return false
		}
	}
	return true
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledSessionPayload {
	p := New()
	s.MustRegisterSystemModule(p)
	i := NewInstalledSessionPayload()
	i.SessionPayload = p
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *SessionPayload {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*SessionPayload)
}
