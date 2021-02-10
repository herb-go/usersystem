package userstatus

import (
	"sort"
	"testing"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/user"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem/usercreate"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/user/status"
	"github.com/herb-go/usersystem"
)

func testSession(id string) *usersystem.Session {
	p := authority.NewPayloads()
	p.Set(usersystem.PayloadUID, []byte(id))
	return usersystem.NewSession().WithType("test").WithPayloads(p)
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
	st, ok := s.Statuses[id]
	if !ok {
		return status.StatusUnkown, user.ErrUserNotExists
	}
	return st, nil
}
func (s *testService) UpdateStatus(id string, st status.Status) error {
	_, ok := s.Statuses[id]
	if !ok {
		return user.ErrUserNotExists
	}
	s.Statuses[id] = st
	return nil
}
func (s *testService) CreateStatus(id string) error {
	_, ok := s.Statuses[id]
	if ok {
		return user.ErrUserExists
	}
	s.Statuses[id] = status.StatusUnkown
	return nil
}
func (s *testService) RemoveStatus(id string) error {
	_, ok := s.Statuses[id]
	if !ok {
		return user.ErrUserNotExists
	}
	delete(s.Statuses, id)
	return nil
}
func (s *testService) getAfterLast(last string, users []string, reverse bool) []string {
	if reverse {
		sort.Sort(sort.Reverse(sort.StringSlice(users)))
	} else {
		sort.Strings(users)
	}
	if last == "" {
		return users
	}
	for k := range users {
		if users[k] > last {
			return users[k:]
		}
	}
	return []string{}
}
func (s *testService) ListUsersByStatus(last string, limit int, reverse bool, st ...status.Status) ([]string, error) {
	m := map[status.Status]bool{}
	for _, v := range st {
		m[v] = true
	}
	alluser := []string{}
	for k, v := range s.Statuses {
		if m[v] {
			alluser = append(alluser, k)
		}
	}
	result := s.getAfterLast(last, alluser, reverse)
	if limit > 0 && limit < len(result) {
		return result[:limit], nil
	}
	return result, nil
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

type testUsersystem struct {
	herbsystem.NopService
}

func (t *testUsersystem) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		usercreate.WrapCreate(func(id string, next func() error) error {
			if id == "testcreateexsits" {
				return user.ErrUserExists
			}
			return nil
		}),
	}
}
func TestStatus(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	userstatus := MustNewAndInstallTo(s)
	s.InstallService(&testUsersystem{})
	s.Ready()
	s.Configuring()
	userstatus.Service = ss
	s.Start()
	defer s.Stop()
	err = userstatus.CreateStatus("test")
	if err != nil {
		panic(err)
	}
	err = userstatus.CreateStatus("test2")
	if err != nil {
		panic(err)
	}
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
	err = userpurge.ExecPurge(s, "testcreate")
	if err != nil {
		panic(err)
	}

	ok, err = usercreate.ExecExist(s, "testcreate")
	if ok || err != nil {
		t.Fatal()
	}
	err = usercreate.ExecCreate(s, "testcreate")
	if err != nil {
		t.Fatal()
	}
	ok, err = usercreate.ExecExist(s, "testcreate")
	if !ok || err != nil {
		t.Fatal()
	}
	err = usercreate.ExecCreate(s, "testcreate")
	if err != user.ErrUserExists {
		t.Fatal()
	}
	err = usercreate.ExecRemove(s, "testcreate")
	if err != nil {
		t.Fatal()
	}
	ok, err = usercreate.ExecExist(s, "testcreate")
	if ok || err != nil {
		t.Fatal()
	}
	err = usercreate.ExecCreate(s, "testcreateexsits")
	if err != user.ErrUserExists {
		t.Fatal()
	}
	ok, err = usercreate.ExecExist(s, "testcreateexsits")
	if ok || err != nil {
		t.Fatal()
	}
	ids, err := userstatus.ListUsersByStatus("", 0, false, status.StatusNormal)
	if len(ids) != 1 || err != nil {
		t.Fatal(ids, err)
	}
}
