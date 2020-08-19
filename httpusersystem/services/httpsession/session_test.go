package httpsession

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/herb-go/usersystem/usersession"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type testSession struct {
	uid         string
	sessiontype usersystem.SessionType
}

func (s *testSession) ID() string {
	return ""
}
func (s *testSession) Type() usersystem.SessionType {
	return s.sessiontype
}
func (s *testSession) UID() (string, error) {
	return s.uid, nil
}
func (s *testSession) SaveUID(string) error {
	return nil
}
func (s *testSession) Payloads() (*authority.Payloads, error) {
	return authority.NewPayloads(), nil
}
func (s *testSession) SavePayloads(p *authority.Payloads) error {
	return nil
}

func (s *testSession) Destory() (bool, error) {
	return false, nil
}
func (s *testSession) Save(key string, v interface{}) error {
	return nil
}
func (s *testSession) Load(key string, v interface{}) error {
	return nil
}
func (s *testSession) Remove(key string) error {
	return nil
}
func (s *testSession) IsNotFoundError(err error) bool {
	return false
}

type testService struct {
	sessions map[string]*testSession
}

func (s *testService) GetSession(id string, st usersystem.SessionType) (usersystem.Session, error) {
	session, ok := s.sessions[id]
	if !ok {
		return nil, nil
	}
	session.sessiontype = st
	return session, nil
}
func (s *testService) SessionMiddleware() func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		req := r.WithContext(context.WithValue(r.Context(), "session", &testSession{uid: "requestsession"}))
		*r = *req
		next(w, r)
	}
}
func (s *testService) GetRequestSession(r *http.Request, st usersystem.SessionType) (usersystem.Session, error) {
	v := r.Context().Value("session")
	session, ok := v.(*testSession)
	if !ok {
		return nil, nil
	}
	session.sessiontype = st
	return session, nil
}
func (s *testService) Start() error {
	return nil
}
func (s *testService) Stop() error {
	return nil
}

func newTestService() *testService {
	return &testService{
		sessions: map[string]*testSession{},
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
	ss.sessions["test"] = &testSession{uid: "test"}
	us, err := usersession.ExecGetSession(s, SessionType, "test")
	if us == nil || err != nil {
		t.Fatal()
	}
	us, err = usersession.ExecGetSession(s, usersystem.SessionType("notexists"), "test")
	if us != nil || err != nil {
		t.Fatal()
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session.Service.SessionMiddleware()(w, r, func(w http.ResponseWriter, r *http.Request) {
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
			uid, err := s.UID()
			w.Write([]byte(uid))
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
}
