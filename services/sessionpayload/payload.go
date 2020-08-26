package sessionpayload

import (
	"context"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var ServiceName = "sessionpayload"

type SessionPayload struct {
	herbsystem.NopService
	Builders []Builder
	Checkers []Checker
}

func New() *SessionPayload {
	return &SessionPayload{}
}

func (p *SessionPayload) InitService() error {
	return nil
}
func (p *SessionPayload) ConfigurService() error {
	p.Builders = []Builder{}
	p.Checkers = []Checker{}
	return nil
}

func (p *SessionPayload) ServiceName() string {
	return ServiceName
}
func (p *SessionPayload) StartService() error {
	for _, v := range p.Builders {
		err := v.Start()
		if err != nil {
			return err
		}
	}
	for _, v := range p.Checkers {
		err := v.Start()
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *SessionPayload) StopService() error {
	var err error
	for _, v := range p.Builders {
		err = herbsystem.MergeError(err, v.Stop())
	}
	for _, v := range p.Checkers {
		err = herbsystem.MergeError(err, v.Stop())
	}
	return err
}
func (p *SessionPayload) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usersession.WrapInitPayloads(p.BuildPayload),
		usersession.WrapCheckSession(p.CheckSession),
	}
}
func (p *SessionPayload) BuildPayload(ctx context.Context, st usersystem.SessionType, uid string, payloads *authority.Payloads) error {
	for _, v := range p.Builders {
		err := v.BuildPayloads(ctx, st, uid, payloads)
		if err != nil {
			return err
		}
	}
	return nil
}
func (p *SessionPayload) CheckSession(session *usersystem.Session) (bool, error) {
	for _, v := range p.Checkers {
		ok, err := v.CheckSession(session)
		if err != nil {
			return false, err
		}
		if !ok {
			return false, nil
		}
	}
	return true, nil
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledSessionPayload {
	p := New()
	err := s.InstallService(p)
	if err != nil {
		panic(err)
	}
	i := NewInstalledSessionPayload()
	i.SessionPayload = p
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*SessionPayload, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*SessionPayload), nil
}
