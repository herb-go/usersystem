package usercreate

import (
	"context"
	"testing"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/user"
	"github.com/herb-go/usersystem"
)

type testModule struct {
	herbsystem.NopModule
	IDList map[string]bool
}

func (s *testModule) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(

		WrapExist(func(id string) bool {
			return s.IDList[id]
		}),
		WrapCreate(func(id string) {
			_, ok := s.IDList[id]
			if ok {
				panic(user.ErrUserExists)
			}
			s.IDList[id] = true
		}, func(id string) {
			delete(s.IDList, id)
		}),
		WrapRemove(func(id string) {
			delete(s.IDList, id)
		}),
	)
	next(ctx, system)
}

func newTestModule() *testModule {
	return &testModule{
		IDList: map[string]bool{},
	}
}
func TestCreate(t *testing.T) {
	s := usersystem.New()
	m := newTestModule()
	s.MustRegisterSystemModule(m)
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	ok := MustExecExist(s, "test")
	if ok {
		t.Fatal()
	}
	MustExecCreate(s, "test")
	ok = MustExecExist(s, "test")
	if !ok {
		t.Fatal()
	}
	err := herbsystem.Catch(func() { MustExecCreate(s, "test") })
	if err != user.ErrUserExists {
		t.Fatal()
	}
	MustExecRemove(s, "test")

	ok = MustExecExist(s, "test")
	if ok {
		t.Fatal()
	}
}
