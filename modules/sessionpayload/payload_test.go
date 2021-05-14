package sessionpayload

import (
	"context"
	"testing"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type testBuilder string

func (t testBuilder) MustBuildPayloads(ctx context.Context, st usersystem.SessionType, id string, p *authority.Payloads) {
	p.Set("testfield", []byte(t))
}

//Start start service
func (t testBuilder) Start() error {
	return nil
}

//Stop stop service
func (t testBuilder) Stop() error {
	return nil
}

type testChecker string

func (c testChecker) MustCheckSession(ctx context.Context, session *usersystem.Session) bool {
	return session.Payloads.LoadString("testfield") == string(c)
}

//Start start service
func (c testChecker) Start() error {
	return nil
}

//Stop stop service
func (c testChecker) Stop() error {
	return nil
}

func TestPayload(t *testing.T) {
	s := usersystem.New()
	b := testBuilder("testpayload")
	c := testChecker("testpayload")
	up := MustNewAndInstallTo(s)
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != up.SessionPayload {
		t.Fatal()
	}
	up.AppendBuilder(b)
	up.AppendChecker(c)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	p := authority.NewPayloads()
	up.MustBuildPayload(nil, "", "", p)
	if p.LoadString("testfield") != string(b) {
		t.Fatal()
	}
	session := usersystem.NewSession().WithPayloads(p)
	ok := up.MustCheckSession(context.Background(), session)
	if !ok {
		t.Fatal()
	}
	session = usersystem.NewSession().WithPayloads(authority.NewPayloads())
	ok = up.MustCheckSession(context.Background(), session)
	if ok {
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
