package userstatus

import (
	"testing"

	"github.com/herb-go/usersystem/useravaliable"

	"github.com/herb-go/usersystem/userpurge"

	"github.com/herb-go/usersystem/userdataset"

	"github.com/herb-go/herb/user/status"
	"github.com/herb-go/usersystem"
)

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
func (s *testService) LoadStatus(idlist ...string) (map[string]status.Status, error) {
	result := map[string]status.Status{}
	for _, v := range idlist {
		st, ok := s.Statuses[v]
		if ok {
			result[v] = st
		}
	}
	return result, nil
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
	delete(t.Statuses, uid)
	return nil
}

func TestStatus(t *testing.T) {
	s := usersystem.New()
	ss := newTestService()
	userstatus := MustNewUserstatus(s)
	userstatus.StatusService = ss
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()
	ds, err := userdataset.ExecNewDataset(s)
	if err != nil {
		panic(err)
	}
	err = userstatus.UpdateStatus(ds, "test", status.StatusBanned)
	if err != nil {
		panic(err)
	}
	err = userstatus.UpdateStatus(ds, "test2", status.StatusBanned)
	if err != nil {
		panic(err)
	}
	ok, err := userstatus.IsUserAvaliable("test")
	if ok || err != nil {
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
	result, err := userstatus.LoadStatus(ds, false, "test", "notexist")
	if err != nil {
		panic(err)
	}
	if len(result) != 1 || result["test"] != status.StatusBanned {
		t.Fatal(result)
	}
	result, err = userstatus.LoadStatus(ds, false, "test", "test2", "notexist")
	if err != nil {
		panic(err)
	}
	if len(result) != 2 || result["test"] != status.StatusBanned || result["test2"] != status.StatusBanned || result["notexist"] != status.StatusUnkown {
		t.Fatal(result["test2"])
	}
	err = userstatus.UpdateStatus(nil, "test", status.StatusNormal)
	if err != nil {
		t.Fatal(err)
	}
	result, err = userstatus.LoadStatus(ds, false, "test", "test2", "notexist")
	if err != nil {
		panic(err)
	}
	if len(result) != 2 || result["test"] != status.StatusBanned || result["test2"] != status.StatusBanned || result["notexist"] != status.StatusUnkown {
		t.Fatal(result["test2"])
	}
	result, err = userstatus.LoadStatus(ds, true, "test", "test2", "notexist")
	if err != nil {
		panic(err)
	}
	if len(result) != 2 || result["test"] != status.StatusNormal || result["test2"] != status.StatusBanned || result["notexist"] != status.StatusUnkown {
		t.Fatal(result["test2"])
	}
	ok, err = useravaliable.ExecAvaliable(s, "test")
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = useravaliable.ExecAvaliable(s, "test2")
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	ok, err = useravaliable.ExecAvaliable(s, "notexist")
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	userpurge.ExecPurge(s, "test")
	_, ok = ss.Statuses["test"]
	if ok {
		t.Fatal(ss)
	}
}
