package userstatus

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herb/user/status"
	"github.com/herb-go/usersystem"
)

type testSession string

func (s testSession) ID() string {
	return ""
}
func (s testSession) Type() usersystem.SessionType {
	return ""
}
func (s testSession) UID() (string, error) {
	return string(s), nil
}
func (s testSession) Payloads() (*authority.Payloads, error) {
	return nil, nil
}
func (s testSession) SavePayloads(*authority.Payloads) error {
	return nil
}
func (s testSession) Destory() error {
	return nil
}
func (s testSession) Save(key string, v interface{}) error {
	return nil
}
func (s testSession) SaveUID(key string) error {
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

type testService struct {
	status.Service
	Statuses map[string]status.Status
}

//Start start service
func (s *testService) Start() error {
	return nil
}

//Stop stop service
func (s *testService) Stop() error {
	return nil
}
func (s *testService) LoadStatus(id string) (status.Status, error) {
	return s.Statuses[id], nil
}
func (s *testService) UpdateStatus(uid string, st status.Status) error {
	s.Statuses[uid] = st
	return nil
}

func newTestService() *testService {
	return &testService{
		Service:  status.NormalOrBannedService,
		Statuses: map[string]status.Status{},
	}
}
func (t *testService) Purge(uid string) error {
	return nil
}

func TestStatus(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	userstatus := MustNewAndInstallTo(s)
	userstatus.Service = ss
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	err = userstatus.UpdateStatus("test", status.StatusNormal)
	if err != nil {
		panic(err)
	}
	err = userstatus.UpdateStatus("test2", status.StatusBanned)
	if err != nil {
		panic(err)
	}
	ok, err := userstatus.IsUserAvaliable("test")
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = userstatus.IsUserAvaliable("test2")
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = userstatus.IsUserAvaliable("notexist")
	if ok || err != nil {
		t.Fatal(ok, err)
	}

	ok, err = usersession.ExecCheckSession(s, testSession("test"))
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = usersession.ExecCheckSession(s, testSession("test2"))
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = usersession.ExecCheckSession(s, testSession("notexist"))
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	err = userpurge.ExecPurge(s, "test")
	if err != nil {
		panic(err)
	}
}
