package userpurge

import (
	"testing"

	"github.com/herb-go/usersystem"

	"github.com/herb-go/herbsystem"
)

type testService struct {
	herbsystem.NopService
	Cached map[string]string
}

func (t *testService) Purge(uid string) error {
	delete(t.Cached, uid)
	return nil
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		Wrap(s),
	}
}

func TestPurge(t *testing.T) {
	s := usersystem.New()
	service := &testService{Cached: map[string]string{"test": "test"}}
	s.InstallService(service)
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	err := ExecPurge(s, "test")
	if err != nil {
		panic(err)
	}
	if service.Cached["test"] != "" {
		t.Fatal()
	}
}
