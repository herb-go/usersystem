package useravaliable

import (
	"testing"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

type testService struct {
	herbsystem.NopService
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		Wrap(func(id string) (bool, error) {
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
	ok, err := ExecAvaliable(s, "exists")
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = ExecAvaliable(s, "notexists")
	if ok || err != nil {
		t.Fatal(ok, err)
	}
}
