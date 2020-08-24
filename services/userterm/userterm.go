package userterm

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"
)

var ServiceName = "userterm"

var PayloadTerm = "term"

type UserTerm struct {
	herbsystem.NopService
	Service
}

func New() *UserTerm {
	return &UserTerm{}
}

func (s *UserTerm) InitService() error {
	return nil
}
func (s *UserTerm) ServiceName() string {
	return ServiceName
}
func (s *UserTerm) StartService() error {
	return s.Service.Start()
}
func (s *UserTerm) StopService() error {
	return s.Service.Stop()
}
func (s *UserTerm) InitPayloads(ctx context.Context, st usersystem.SessionType, uid string, payloads *authority.Payloads) error {
	term, err := s.Service.CurrentTerm(uid)
	if err != nil {
		return err
	}
	payloads.Set(PayloadTerm, []byte(term))
	return nil
}
func (s *UserTerm) CheckSession(session *usersystem.Session) (bool, error) {
	uid := session.UID()
	if uid == "" {
		return false, nil
	}
	term, err := s.Service.CurrentTerm(uid)
	if err != nil {
		return false, err
	}
	sessionterm := session.Payloads.LoadString(PayloadTerm)
	return term == sessionterm, nil
}
func (s *UserTerm) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userpurge.Wrap(s),
		usersession.WrapInitPayloads(s.InitPayloads),
		usersession.WrapCheckSession(s.CheckSession),
	}
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserTerm {
	term := New()
	err := s.InstallService(term)
	if err != nil {
		panic(err)
	}
	i := NewInstalledUserTerm()
	i.UserTerm = term
	i.UserSystem = s
	return i
}

func GetService(s *usersystem.UserSystem) (*UserTerm, error) {
	v, err := s.GetConfigurableService(ServiceName)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return v.(*UserTerm), nil
}
