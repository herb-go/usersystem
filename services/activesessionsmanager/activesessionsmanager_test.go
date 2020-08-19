package activesessionsmanager

import (
	"strconv"
	"testing"
	"time"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/usersession"
)

type testSession struct {
	id       string
	payloads *authority.Payloads
}

func (s testSession) ID() string {
	return ""
}
func (s testSession) Type() usersystem.SessionType {
	return "test"
}
func (s testSession) UID() (string, error) {
	return s.id, nil
}
func (s testSession) SaveUID(string) error {
	return nil
}
func (s testSession) Payloads() (*authority.Payloads, error) {
	return s.payloads, nil
}
func (s testSession) SavePayloads(p *authority.Payloads) error {
	return nil
}

func (s testSession) Destory() (bool, error) {
	return false, nil
}
func (s testSession) Save(key string, v interface{}) error {
	return nil
}
func (s testSession) Load(key string, v interface{}) error {
	return nil
}
func (s testSession) Remove(key string) error {
	return nil
}
func (s testSession) IsNotFoundError(err error) bool {
	return false
}
func newTestSession(id string) *testSession {
	return &testSession{
		id:       id,
		payloads: authority.NewPayloads(),
	}
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
func (s testService) OnSessionActive(session usersystem.Session, uid string) error {
	return nil
}
func (s testService) GetActiveSessions(usersystem.SessionType) ([]*usersession.ActiveSession, bool, error) {
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
	err = usersession.ExecOnSessionActive(s, newTestSession("123"))
	if err != nil {
		panic(err)
	}
	sessions, err := usersession.ExecGetActiveSessions(s, "test")
	if err != nil || len(sessions) != 1 || sessions[0].SessionID != "12345" {
		t.Fatal(sessions, err)
	}
	session := newTestSession("123")
	err = usersession.ExecInitPayloads(s, session)
	if err != nil {
		panic(err)
	}
	sn, err := GetSerialNumber(session)
	if err != nil {
		panic(err)
	}
	if sn == "" {
		t.Fatal(sn)
	}
}
