package websession

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

var testErrNotFound = errors.New("not found")

func testSession(id string) *usersystem.Session {
	p := authority.NewPayloads()
	p.Set(usersystem.PayloadUID, []byte(id))
	return usersystem.NewSession().WithType("test").WithPayloads(p).WithID(id)
}

type testService struct {
	sessions map[string]*usersystem.Session
	values   map[string][]byte
}

func (s *testService) GetSession(st usersystem.SessionType, id string) (*usersystem.Session, error) {
	session, ok := s.sessions[id]
	if !ok {
		return nil, nil
	}
	session.WithType(st)
	return session, nil
}
func (s *testService) RevokeSession(code string) (bool, error) {
	_, ok := s.sessions[code]
	if !ok {
		return false, nil
	}
	delete(s.sessions, code)
	return true, nil
}
func (s *testService) SessionMiddleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		req := r.WithContext(context.WithValue(r.Context(), "session", testSession("requestsession")))
		*r = *req
		next(w, r)
	}
}
func (s *testService) GetRequestSession(r *http.Request, st usersystem.SessionType) (*usersystem.Session, error) {
	v := r.Context().Value("session")
	session, ok := v.(*usersystem.Session)
	if !ok {
		return nil, nil
	}
	return session, nil
}
func (s *testService) LoginRequestSession(r *http.Request, payloads *authority.Payloads) (*usersystem.Session, error) {
	session := usersystem.NewSession().WithPayloads(payloads).WithID("id")
	req := r.WithContext(context.WithValue(r.Context(), "session", session))
	*r = *req
	return session, nil
}
func (s *testService) LogoutRequestSession(r *http.Request) (bool, error) {
	req := r.WithContext(context.WithValue(r.Context(), "session", nil))
	*r = *req
	return true, nil
}
func (s *testService) Set(r *http.Request, fieldname string, v interface{}) error {
	bs, err := json.Marshal(v)
	if err != nil {
		return err
	}
	s.values[fieldname] = bs
	return nil
}

//Get get session by field name with given value.
func (s *testService) Get(r *http.Request, fieldname string, v interface{}) error {
	bs, ok := s.values[fieldname]
	if !ok {
		return testErrNotFound
	}
	return json.Unmarshal(bs, v)
}

// Del del session value by field name .
func (s *testService) Del(r *http.Request, fieldname string) error {
	delete(s.values, fieldname)
	return nil
}

// IsNotFoundError check if given error is session not found error.
func (s *testService) IsNotFoundError(err error) bool {
	return err == testErrNotFound
}
func (s *testService) Start() error {
	return nil
}
func (s *testService) Stop() error {
	return nil
}

func newTestService() *testService {
	return &testService{
		sessions: map[string]*usersystem.Session{},
		values:   map[string][]byte{},
	}
}
func TestHTTPSession(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	session := MustNewAndInstallTo(s)
	s.Ready()
	s.Configuring()
	session.Service = ss
	s.Start()
	defer s.Stop()
	ss.sessions["test"] = testSession("test")
	us, err := usersession.ExecGetSession(s, usersystem.SessionType("notexists"), "test")
	if us != nil || err != nil {
		t.Fatal()
	}
	us, err = usersession.ExecGetSession(s, SessionType, "test")
	if us == nil || err != nil {
		t.Fatal()
	}
	sessionnotexist := usersystem.NewSession().WithType("notexists").WithPayloads(authority.NewPayloads())
	sessionnotexist.Payloads.Set(usersystem.PayloadRevokeCode, []byte("test"))
	ok, err := usersession.ExecRevokeSession(s, sessionnotexist)
	if ok || err != nil {
		t.Fatal()
	}
	us, err = usersession.ExecGetSession(s, SessionType, "test")
	if us == nil || err != nil {
		t.Fatal()
	}
	us = usersystem.NewSession().WithType(SessionType).WithPayloads(authority.NewPayloads())
	us.Payloads.Set(usersystem.PayloadRevokeCode, []byte("test"))
	ok, err = usersession.ExecRevokeSession(s, us)
	if !ok || err != nil {
		t.Fatal()
	}
	us, err = usersession.ExecGetSession(s, SessionType, "test")
	if us != nil || err != nil {
		t.Fatal()
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session.Middleware()(w, r, func(w http.ResponseWriter, r *http.Request) {
			s, err := session.GetRequestSession(r)
			if err != nil {
				panic(err)
			}
			if s == nil {
				w.Write([]byte(""))
				return
			}
			if err != nil {
				panic(err)
			}
			w.Write([]byte(s.UID()))
			return
		})
	}))
	defer server.Close()
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Fatal(err)
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != 200 || string(bs) != "requestsession" {
		t.Fatal()
	}
	req, err := http.NewRequest("GET", server.URL, nil)
	if err != nil {
		t.Fatal(err)
	}
	us, err = session.Login(req, "ttt")
	if err != nil {
		t.Fatal(err)
	}
	if us.UID() != "ttt" {
		t.Fatal(us.UID())
	}
	us2, err := session.GetRequestSession(req)
	if err != nil {
		t.Fatal(err)
	}
	if us2.ID != us.ID {
		t.Fatal(us2, us)
	}
	ok, err = session.Logout(req)
	if !ok || err != nil {
		t.Fatal(ok, err)
	}
	us2, err = session.GetRequestSession(req)
	if err != nil {
		t.Fatal(err)
	}
	if us2 != nil {
		t.Fatal(us2)
	}
	var result = ""
	err = session.Get(req, "test", &result)
	if err == nil || !session.IsNotFoundError(err) {
		t.Fatal()
	}
	err = session.Set(req, "test", "testvalue")
	if err != nil {
		t.Fatal()
	}
	err = session.Get(req, "test", &result)
	if err != nil {
		t.Fatal()
	}
	if result != "testvalue" {
		t.Fatal(result)
	}
	err = session.Del(req, "test")
	if err != nil {
		t.Fatal()
	}
	err = session.Get(req, "test", &result)
	if err == nil || !session.IsNotFoundError(err) {
		t.Fatal()
	}
}
