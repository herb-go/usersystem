package userpurge

import (
	"context"
	"testing"

	"github.com/herb-go/usersystem"

	"github.com/herb-go/herbsystem"
)

type testModule struct {
	herbsystem.NopModule
	Cached map[string]string
}

func (m *testModule) Purge(uid string) error {
	delete(m.Cached, uid)
	return nil
}

func (m *testModule) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(Wrap(m))
	next(ctx, system)
}

func TestPurge(t *testing.T) {
	s := usersystem.New()
	m := &testModule{Cached: map[string]string{"test": "test"}}
	s.MustRegisterSystemModule(m)
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	MustExecPurge(s, "test")
	if m.Cached["test"] != "" {
		t.Fatal()
	}
}
