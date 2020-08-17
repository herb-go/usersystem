package userprofile

import (
	"testing"

	"github.com/herb-go/usersystem/userpurge"

	"github.com/herb-go/usersystem/userdataset"

	"github.com/herb-go/herb/user/profile"

	"github.com/herb-go/usersystem"
)

type testService struct {
	profiles map[string]*profile.Profile
}

func (s *testService) GetProfile(id string) (*profile.Profile, error) {
	p, ok := s.profiles[id]
	if !ok {
		return profile.NewProfile(), nil
	}
	return p.Clone(), nil
}
func (s *testService) UpdateProfile(id string, p *profile.Profile) error {
	uprofile, ok := s.profiles[id]
	if !ok {
		uprofile = profile.NewProfile()
		s.profiles[id] = uprofile
	}
	uprofile.With("test", p.Load("test"))
	return nil
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
	s.Ready()
	s.Configuring()
	profileservice.AppendService(ss)
	s.Start()
	defer s.Stop()
	ds, err := userdataset.ExecNewDataset(s)
	if err != nil {
		panic(err)
	}
	result, err := profileservice.LoadProfiles(ds, false, "test")
	if err != nil {
		panic(err)
	}
	if len(result) != 1 || len(result["test"].Data()) != 0 {
		t.Fatal(result)
	}
	newprofile := profile.NewProfile().With("test", "testvalue").With("test2", "test2value")
	err = profileservice.UpdateProfile(ds, "test", newprofile)
	if err != nil {
		panic(err)
	}
	result, err = profileservice.LoadProfiles(ds, false, "test")
	if err != nil {
		panic(err)
	}
	if len(result) != 1 || len(result["test"].Data()) != 1 || result["test"].Load("test") != "testvalue" {
		t.Fatal(result)
	}
	result, err = profileservice.LoadProfiles(ds, false, "test")
	if err != nil {
		panic(err)
	}
	if len(result) != 1 || len(result["test"].Data()) != 1 || result["test"].Load("test") != "testvalue" {
		t.Fatal(result)
	}
	newprofile = profile.NewProfile().With("test", "newvalue")
	err = profileservice.UpdateProfile(nil, "test", newprofile)
	if err != nil {
		panic(err)
	}
	result, err = profileservice.LoadProfiles(ds, false, "test")
	if err != nil {
		panic(err)
	}
	if len(result) != 1 || result["test"].Load("test") != "testvalue" {
		t.Fatal(result)
	}
	result, err = profileservice.LoadProfiles(ds, true, "test")
	if err != nil {
		panic(err)
	}
	if len(result) != 1 || result["test"].Load("test") != "newvalue" {
		t.Fatal(result)
	}
	err = userpurge.ExecPurge(s, "test")
	if err != nil {
		panic(err)
	}
}
