package sessionpayload

import (
	"context"
	"testing"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type testBuilder string

func (t testBuilder) BuildPayloads(ctx context.Context, st usersystem.SessionType, id string, p *authority.Payloads) error {
	p.Set("testfield", []byte(t))
	return nil
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

func (c testChecker) CheckSession(ctx context.Context, session *usersystem.Session) (bool, error) {
	return session.Payloads.LoadString("testfield") == string(c), nil
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
	var err error
	s := usersystem.New()
	b := testBuilder("testpayload")
	c := testChecker("testpayload")
	up := MustNewAndInstallTo(s)
	s.Ready()
	s.Configuring()
	up.AppendBuilder(b)
	up.AppendChecker(c)
	s.Start()
	defer s.Stop()
	p := authority.NewPayloads()
	err = up.BuildPayload(nil, "", "", p)
	if err != nil {
		t.Fatal(err)
	}
	if p.LoadString("testfield") != string(b) {
		t.Fatal()
	}
	session := usersystem.NewSession().WithPayloads(p)
	ok, err := up.CheckSession(context.Background(), session)
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal()
	}
	session = usersystem.NewSession().WithPayloads(authority.NewPayloads())
	ok, err = up.CheckSession(context.Background(), session)
	if err != nil {
		t.Fatal(err)
	}
	if ok {
		t.Fatal()
	}
}
