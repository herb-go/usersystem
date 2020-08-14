package userterm

import (
	"testing"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type testSession string

func (s testSession) ID() string {
	return ""
}
func (s testSession) Type() usersystem.SessionType {
	return "test"
}
func (s testSession) UID() (string, error) {
	return string(s), nil
}
func (s testSession) SaveUID(string) error {
	return nil
}
func (s testSession) Payloads() (*authority.Payloads, error) {
	return authority.NewPayloads(), nil
}
func (s testSession) SavePayloads(p *authority.Payloads) error {
	return nil
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
	Term string
}

//Start start service
func (s *testService) Start() error {
	return nil
}

//Stop stop service
func (s *testService) Stop() error {
	return nil
}
func (s *testService) CurrentTerm(uid string) (string, error) {
	return s.Term, nil
}
func (s *testService) StartNewTerm(uid string) (string, error) {
	s.Term = "New"
	return s.Term, nil
}

func (t *testService) Purge(uid string) error {
	return nil
}
func newTestService() *testService {
	return &testService{}
}
func TestStatus(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	userterm := MustNewAndInstallTo(s)
	userterm.Service = ss
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	session := testSession("test")
	err = usersession.ExecInitPayloads(s, session)
	if err != nil {
		panic(err)
	}
	ok, err := usersession.ExecCheckSession(s, session)
	if !ok || err != nil {
		t.Fatal()
	}
	_, err = userterm.StartNewTerm("test")
	ok, err = usersession.ExecCheckSession(s, session)
	if ok || err != nil {
		t.Fatal()
	}

}
