package userterm

import (
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
func (s *UserTerm) InitPayloads(session usersystem.Session, uid string, payloads *authority.Payloads) error {
	term, err := s.Service.CurrentTerm(uid)
	if err != nil {
		return err
	}
	payloads.Set(PayloadTerm, []byte(term))
	return nil
}
func (s *UserTerm) CheckSession(session usersystem.Session, uid string, payloads *authority.Payloads) (bool, error) {
	term, err := s.Service.CurrentTerm(uid)
	if err != nil {
		return false, err
	}
	sessionterm := payloads.LoadString(PayloadTerm)
	return term == sessionterm, nil
}
func (s *UserTerm) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userpurge.Wrap(s),
		usersession.WrapInitPayloads(s.InitPayloads),
		usersession.WrapCheckSession(s.CheckSession),
	}
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *UserTerm {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
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
