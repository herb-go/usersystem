package userstatus

import (
	"context"
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
func (s *testService) MustLoadStatus(id string) status.Status {
	st, ok := s.Statuses[id]
	if !ok {
		panic(user.ErrUserNotExists)
	}
	return st
}
func (s *testService) MustUpdateStatus(id string, st status.Status) {
	_, ok := s.Statuses[id]
	if !ok {
		panic(user.ErrUserNotExists)
	}
	s.Statuses[id] = st
}
func (s *testService) MustCreateStatus(id string) {
	_, ok := s.Statuses[id]
	if ok {
		panic(user.ErrUserExists)
	}
	s.Statuses[id] = status.StatusUnkown
}
func (s *testService) MustRemoveStatus(id string) {
	_, ok := s.Statuses[id]
	if !ok {
		panic(user.ErrUserNotExists)
	}
	delete(s.Statuses, id)
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
func (s *testService) MustListUsersByStatus(last string, limit int, reverse bool, st ...status.Status) []string {
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
		return result[:limit]
	}
	return result
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
	herbsystem.NopModule
}

func (t *testUsersystem) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		usercreate.WrapCreate(func(id string) {
			if id == "testcreateexsits" {
				panic(user.ErrUserExists)
			}
		}, nil),
	)
	next(ctx, system)
}
func TestStatus(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	userstatus := MustNewAndInstallTo(s)
	s.MustRegisterSystemModule(&testUsersystem{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != userstatus.UserStatus {
		t.Fatal()
	}
	userstatus.Service = ss
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	userstatus.MustCreateStatus("test")
	userstatus.MustCreateStatus("test2")
	userstatus.MustUpdateStatus("test", status.StatusNormal)
	userstatus.MustUpdateStatus("test2", status.StatusBanned)

	ok := userstatus.MustIsUserAvaliable("test")
	if !ok {
		t.Fatal(ok)
	}
	ok = userstatus.MustIsUserAvaliable("test2")
	if ok {
		t.Fatal(ok)
	}
	ok = userstatus.MustIsUserAvaliable("notexist")
	if ok {
		t.Fatal(ok)
	}

	ok = usersession.MustExecCheckSession(s, testSession("test"))
	if !ok {
		t.Fatal(ok)
	}
	ok = usersession.MustExecCheckSession(s, testSession("test2"))
	if ok {
		t.Fatal(ok)
	}
	ok = usersession.MustExecCheckSession(s, testSession("notexist"))
	if ok || err != nil {
		t.Fatal(ok, err)
	}
	userpurge.MustExecPurge(s, "testcreate")

	ok = usercreate.MustExecExist(s, "testcreate")
	if ok {
		t.Fatal()
	}
	usercreate.MustExecCreate(s, "testcreate")

	ok = usercreate.MustExecExist(s, "testcreate")
	if !ok {
		t.Fatal()
	}
	err = herbsystem.Catch(
		func() {
			usercreate.MustExecCreate(s, "testcreate")
		})
	if err != user.ErrUserExists {
		t.Fatal()
	}
	err = herbsystem.Catch(
		func() {
			usercreate.MustExecRemove(s, "testcreate")
		})
	if err != user.ErrUserNotExists {
		t.Fatal()
	}
	ok = usercreate.MustExecExist(s, "testcreate")
	if ok {
		t.Fatal()
	}

	err = herbsystem.Catch(
		func() {
			usercreate.MustExecCreate(s, "testcreateexsits")
		})
	if err != user.ErrUserExists {
		t.Fatal()
	}
	ok = usercreate.MustExecExist(s, "testcreateexsits")
	if ok {
		t.Fatal()
	}
	ids := userstatus.MustListUsersByStatus("", 0, false, status.StatusNormal)
	if len(ids) != 1 {
		t.Fatal(ids)
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
