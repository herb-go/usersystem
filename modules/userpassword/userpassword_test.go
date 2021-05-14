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
func (s *testService) VerifyPassword(uid string, password string) (bool, error) {
	return password == s.Password, nil
}

//PasswordChangeable return password changeable
func (s *testService) PasswordChangeable() bool {
	return true
}

//UpdatePassword update user password
//Return any error if raised
func (s *testService) UpdatePassword(uid string, password string) error {
	s.Password = password
	return nil
}

func (t *testService) Purge(uid string) error {
	return nil
}
func newTestService() *testService {
	return &testService{}
}
func TestUserPassword(t *testing.T) {
	var err error
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
	err = userpassword.UpdatePassword("id", "old")
	if err != nil {
		panic(err)
	}
	ok, err = userpassword.VerifyPassword("id", "old")
	if !ok || err != nil {
		t.Fatal()
	}
	err = userpassword.UpdatePassword("id", "new")
	if err != nil {
		panic(err)
	}
	ok, err = userpassword.VerifyPassword("id", "old")
	if ok || err != nil {
		t.Fatal()
	}
	ok, err = userpassword.VerifyPassword("id", "new")
	if !ok || err != nil {
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
