package activesessions

import (
	"strconv"
	"testing"
	"time"

	"github.com/herb-go/herbsystem"

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

func (s *testService) MustCreateSerialNumber() string {
	SerialNumber = SerialNumber + 1
	return strconv.Itoa(SerialNumber)
}
func (s *testService) MustConfig(st usersystem.SessionType) *Config {
	return &Config{Supported: true, Duration: time.Minute}
}
func (s *testService) MustOnSessionActive(session *usersystem.Session) {
}
func (s *testService) MustGetActiveSessions(usersystem.SessionType, string) []*Active {
	return []*Active{&Active{SessionID: "12345"}}
}
func (s *testService) MustPurgeActiveSession(st usersystem.SessionType, uid string, serialnumber string) {
}
func (s *testService) Start() error {
	return nil
}
func (s *testService) Stop() error {
	return nil
}

func TestActiveSessionsManager(t *testing.T) {
	s := usersystem.New()
	m := MustNewAndInstallTo(s)
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != m.ActiveSessions {
		t.Fatal()
	}
	m.Service = &testService{}
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	config := m.MustConfig("test")
	if config == nil || config.Supported == false || config.Duration != time.Minute {
		t.Fatal()
	}
	usersession.MustExecOnSessionActive(s, testSession("123"))
	sessions := m.MustGetActiveSessions("test", "test")
	if len(sessions) != 1 || sessions[0].SessionID != "12345" {
		t.Fatal(sessions)
	}

	p := usersession.MustExecInitPayloads(s, s.SystemContext(), "test", "123")
	session := testSession("123").WithPayloads(p)
	sn := MustGetSerialNumber(session)
	if sn == "" {
		t.Fatal(sn)
	}
	m.MustPurgeActiveSession(session)
}

func TestMustGetModule(t *testing.T) {
	s := usersystem.New()
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != nil {
		t.Fatal()
	}
}
