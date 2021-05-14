package userterm

import (
	"testing"

	"github.com/herb-go/herbsystem"
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
func (s *testService) MustCurrentTerm(uid string) string {
	return s.Term
}
func (s *testService) MustStartNewTerm(uid string) string {
	s.Term = "New"
	return s.Term
}

func (t *testService) Purge(uid string) error {
	return nil
}
func newTestService() *testService {
	return &testService{}
}
func TestTerm(t *testing.T) {
	s := usersystem.New()
	ss := newTestService()
	userterm := MustNewAndInstallTo(s)
	herbsystem.MustReady(s)
	userterm.Service = ss
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != userterm.UserTerm {
		t.Fatal()
	}
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	p := usersession.MustExecInitPayloads(s, s.SystemContext(), "test", "test")

	session := testSession("test").WithPayloads(p)
	ok := usersession.MustExecCheckSession(s, session)
	if !ok {
		t.Fatal()
	}
	userterm.MustStartNewTerm("test")
	ok = usersession.MustExecCheckSession(s, session)
	if ok {
		t.Fatal()
	}

}

func TestMustGetModule(t *testing.T) {
	s := usersystem.New()
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != nil {
		t.Fatal()
	}
}
