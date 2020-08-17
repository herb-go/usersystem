package userrole

import (
	"testing"

	"github.com/herb-go/herbsecurity/authorize/role"

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
func (s *testService) Roles(uid string) (*role.Roles, error) {
	return role.NewRoles(role.NewRole("test")), nil
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
	userstatus := MustNewAndInstallTo(s)
	s.Ready()
	s.Configuring()
	userstatus.Service = ss
	s.Start()
	defer s.Stop()
	roles, err := userstatus.Roles("test")
	if err != nil {
		panic(err)
	}
	if roles.Data()[0].Name != "test" {
		t.Fatal(roles)
	}
	err = userpurge.ExecPurge(s, "test")
	if err != nil {
		panic(err)
	}
}
