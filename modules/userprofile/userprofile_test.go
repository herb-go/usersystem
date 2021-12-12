package userprofile

import (
	"testing"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem/userpurge"

	"github.com/herb-go/user/profile"

	"github.com/herb-go/usersystem"
)

type testService struct {
	profiles map[string]*profile.Profile
}

func (s *testService) MustGetProfile(id string) *profile.Profile {
	p, ok := s.profiles[id]
	if !ok {
		return profile.NewProfile()
	}
	return p.Clone()
}
func (s *testService) MustUpdateProfile(id string, p *profile.Profile) {
	uprofile, ok := s.profiles[id]
	if !ok {
		uprofile = profile.NewProfile()
		s.profiles[id] = uprofile
	}
	uprofile.With("test", p.Load("test"))
}

//Start start service
func (s *testService) Start() error {
	return nil
}

//Stop stop service
func (s *testService) Stop() error {
	return nil
}

func (t *testService) Purge(uid string) error {
	return nil
}
func newTestService() *testService {
	return &testService{
		profiles: map[string]*profile.Profile{},
	}
}
func TestUserProfile(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	profileservice := MustNewAndInstallTo(s)
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	profileservice.AppendService(ss)
	if MustGetModule(s) != profileservice.UserProfile {
		t.Fatal()
	}
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	result := profileservice.MustLoadProfiles("test")
	if len(result) != 1 || len(result["test"].Data()) != 0 {
		t.Fatal(result)
	}
	newprofile := profile.NewProfile().With("test", "testvalue").With("test2", "test2value")
	profileservice.MustUpdateProfile("test", newprofile)
	result = profileservice.MustLoadProfiles("test")
	if len(result) != 1 || len(result["test"].Data()) != 1 || result["test"].Load("test") != "testvalue" {
		t.Fatal(result)
	}
	result = profileservice.MustLoadProfiles("test")
	if len(result) != 1 || len(result["test"].Data()) != 1 || result["test"].Load("test") != "testvalue" {
		t.Fatal(result)
	}
	newprofile = profile.NewProfile().With("test", "newvalue")
	profileservice.MustUpdateProfile("test", newprofile)
	if err != nil {
		panic(err)
	}
	result = profileservice.MustLoadProfiles("test")
	if len(result) != 1 || result["test"].Load("test") != "newvalue" {
		t.Fatal(result)
	}
	userpurge.MustExecPurge(s, "test")

}

func TestMustGetModule(t *testing.T) {
	s := usersystem.New()
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != nil {
		t.Fatal()
	}
}
