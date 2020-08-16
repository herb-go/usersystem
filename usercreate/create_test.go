package usercreate

import (
	"testing"

	"github.com/herb-go/herb/user"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type testService struct {
	herbsystem.NopService
	IDList map[string]bool
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		WrapExist(func(id string) (bool, error) {
			return s.IDList[id], nil
		}),
		WrapCreate(func(id string, next func() error) error {
			_, ok := s.IDList[id]
			if ok {
				return user.ErrUserExists
			}
			s.IDList[id] = true
			err := next()
			if err != nil {
				delete(s.IDList, id)
			}
			return err
		}),
		WrapRemove(func(id string) error {
			delete(s.IDList, id)
			return nil
		}),
	}
}

func newTestService() *testService {
	return &testService{
		IDList: map[string]bool{},
	}
}
func TestCreate(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	s.InstallService(ss)
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	ok, err := ExecExist(s, "test")
	if ok || err != nil {
		t.Fatal()
	}
	err = ExecCreate(s, "test")
	if err != nil {
		t.Fatal()
	}
	ok, err = ExecExist(s, "test")
	if !ok || err != nil {
		t.Fatal()
	}
	err = ExecCreate(s, "test")
	if err != user.ErrUserExists {
		t.Fatal()
	}
	err = ExecRemove(s, "test")
	if err != nil {
		t.Fatal()
	}
	ok, err = ExecExist(s, "test")
	if ok || err != nil {
		t.Fatal()
	}
}
