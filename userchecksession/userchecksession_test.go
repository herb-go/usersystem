package userchecksession

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type testSession string

func (s testSession) ID() string {
	return ""
}
func (s testSession) Type() usersystem.SessionType {
	return ""
}
func (s testSession) UID() (string, error) {
	return string(s), nil
}
func (s testSession) Payloads() (*authority.Payloads, error) {
	return nil, nil
}
func (s testSession) Destory() error {
	return nil
}
func (s testSession) Save(key string, v interface{}) error {
	return nil
}
func (s testSession) Load(key string, v interface{}) error {
	return nil
}
func (s testSession) Remove(key string) error {
	return nil
}
func (s testSession) IsNotFoundError(err error) bool {
	return false
}

type testService struct {
	herbsystem.NopService
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		Wrap(func(session usersystem.Session, id string, payloads *authority.Payloads) (bool, error) {
			return id == "exists", nil
		}),
	}
}
func TestAvaliable(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	ok, err := ExecCheckSession(s, testSession("exists"))
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = ExecCheckSession(s, testSession("notexists"))
	if ok || err != nil {
		t.Fatal(ok, err)
	}
}
