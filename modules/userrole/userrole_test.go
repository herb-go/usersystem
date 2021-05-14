package userrole

import (
	"testing"

	"github.com/herb-go/herbsecurity/authorize/role"
	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

type testService struct {
}

//Start start service
func (s *testService) Start() error {
	return nil
}

//Stop stop service
func (s *testService) Stop() error {
	return nil
}
func (s *testService) MustRoles(uid string) *role.Roles {
	return role.NewRoles(role.NewRole("test"))
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

	userrole := MustNewAndInstallTo(s)
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != userrole.UserRole {
		t.Fatal()
	}
	userrole.Service = ss
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)

	roles := userrole.MustRoles("test")
	if err != nil {
		panic(err)
	}
	if roles.Data()[0].Name != "test" {
		t.Fatal(roles)
	}
	userpurge.MustExecPurge(s, "test")
}

func TestMustGetModule(t *testing.T) {
	s := usersystem.New()
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != nil {
		t.Fatal()
	}
}
