package userpassword

import (
	"testing"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem"
)

type testService struct {
	Password string
}

//Start start service
func (s *testService) Start() error {
	return nil
}

//Stop stop service
func (s *testService) Stop() error {
	return nil
}
func (s *testService) MustVerifyPassword(uid string, password string) bool {
	return password == s.Password
}

//PasswordChangeable return password changeable
func (s *testService) PasswordChangeable() bool {
	return true
}

//UpdatePassword update user password
//Return any error if raised
func (s *testService) MustUpdatePassword(uid string, password string) {
	s.Password = password
}

func (t *testService) Purge(uid string) error {
	return nil
}
func newTestService() *testService {
	return &testService{}
}
func TestUserPassword(t *testing.T) {
	s := usersystem.New()
	ss := newTestService()
	userpassword := MustNewAndInstallTo(s)
	userpassword.Service = ss
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != userpassword.UserPassword {
		t.Fatal()
	}
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	ok := userpassword.PasswordChangeable()
	if !ok {
		t.Fatal()
	}
	userpassword.MustUpdatePassword("id", "old")

	ok = userpassword.MustVerifyPassword("id", "old")
	if !ok {
		t.Fatal()
	}
	userpassword.MustUpdatePassword("id", "new")
	ok = userpassword.MustVerifyPassword("id", "old")
	if ok {
		t.Fatal()
	}
	ok = userpassword.MustVerifyPassword("id", "new")
	if !ok {
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
