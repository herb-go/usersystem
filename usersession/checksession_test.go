package usersession

import (
	"testing"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem"
)

func TestCheckSession(t *testing.T) {
	s := usersystem.New()
	s.MustRegisterSystemModule(&testModule{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	ok := MustExecCheckSession(s, testSession("exists"))
	if !ok {
		t.Fatal(ok)
	}
	ok = MustExecCheckSession(s, testSession("notexists"))
	if ok {
		t.Fatal(ok)
	}
}
