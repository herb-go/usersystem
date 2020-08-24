package userterm

import (
	"testing"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

func testSession(id string) *usersystem.Session {
	p := authority.NewPayloads()
	p.Set(usersystem.PayloadUID, []byte(id))
	return usersystem.NewSession().WithType("test").WithPayloads(p)
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
	s.Ready()
	s.Configuring()
	userterm.Service = ss
	s.Start()
	defer s.Stop()
	p, err := usersession.ExecInitPayloads(s, s.Context, "test", "test")
	if err != nil {
		panic(err)
	}
	session := testSession("test").WithPayloads(p)
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
