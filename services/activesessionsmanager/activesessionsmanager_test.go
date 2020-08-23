package activesessionsmanager

import (
	"strconv"
	"testing"
	"time"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

func testSession(id string) *usersystem.Session {
	p := authority.NewPayloads()
	p.Set(usersystem.PayloadUID, []byte(id))
	return usersystem.NewSession().WithType("test").WithPayloads(p)
}

type testService struct {
}

var SerialNumber int

func (s testService) CreateSerialNumber() (string, error) {
	SerialNumber = SerialNumber + 1
	return strconv.Itoa(SerialNumber), nil
}
func (s testService) Config(st usersystem.SessionType) (*usersession.Config, error) {
	return &usersession.Config{Supported: true, Duration: time.Minute}, nil
}
func (s testService) OnSessionActive(session *usersystem.Session) error {
	return nil
}
func (s testService) GetActiveSessions(usersystem.SessionType, string) ([]*usersession.ActiveSession, bool, error) {
	return []*usersession.ActiveSession{&usersession.ActiveSession{SessionID: "12345"}}, true, nil
}
func (s testService) Start() error {
	return nil
}
func (s testService) Stop() error {
	return nil
}

func TestActiveSessionsManager(t *testing.T) {
	s := usersystem.New()
	m := MustNewAndInstallTo(s)
	s.Ready()
	s.Configuring()
	m.Service = &testService{}
	s.Start()
	defer s.Stop()
	config, err := usersession.ExecActiveSessionManagerConfig(s, "test")
	if config == nil || config.Supported == false || config.Duration != time.Minute || err != nil {
		t.Fatal()
	}
	err = usersession.ExecOnSessionActive(s, testSession("123"))
	if err != nil {
		panic(err)
	}
	sessions, err := usersession.ExecGetActiveSessions(s, "test", "test")
	if err != nil || len(sessions) != 1 || sessions[0].SessionID != "12345" {
		t.Fatal(sessions, err)
	}

	p, err := usersession.ExecInitPayloads(s, s.Context, "test", "123")
	if err != nil {
		panic(err)
	}
	session := testSession("123").WithPayloads(p)
	sn, err := GetSerialNumber(session)
	if err != nil {
		panic(err)
	}
	if sn == "" {
		t.Fatal(sn)
	}
}
