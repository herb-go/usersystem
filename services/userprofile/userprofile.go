package userprofile

import (
	"github.com/herb-go/herb/user/profile"
	"github.com/herb-go/herb/user/status"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userdataset"
	"github.com/herb-go/usersystem/userpurge"
)

var ServiceName = "profile"

var DatatypeProfile = usersystem.DataType("profile")

func LoadProfile(ds usersystem.Dataset, id string) (*profile.Profile, bool) {
	p, ok := ds.Get(DatatypeProfile, id)
	if !ok {
		return nil, false
	}
	return p.(*profile.Profile), true
}

func SetProfile(ds usersystem.Dataset, id string, p *profile.Profile) {
	ds.Set(DatatypeProfile, id, p)
}

type UserProfile struct {
	herbsystem.NopService
	Services []Service
}

func New() *UserProfile {
	return &UserProfile{}
}
func (s *UserProfile) ConfigurService() error {
	s.Services = []Service{}
	return nil
}

func (s *UserProfile) InitService() error {
	return nil
}
func (s *UserProfile) ServiceName() string {
	return ServiceName
}
func (s *UserProfile) StartService() error {
	errs := herbsystem.NewErrors()
	for _, v := range s.Services {
		errs.Add(v.Start())
	}
	return errs.ToError()
}
func (s *UserProfile) StopService() error {
	errs := herbsystem.NewErrors()
	for _, v := range s.Services {
		errs.Add(v.Stop())
	}
	return errs.ToError()
}
func (s *UserProfile) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userdataset.InitDatasetTypeAction(DatatypeProfile),
		userpurge.Wrap(s),
	}
}

func (s *UserProfile) LoadProfile(dataset usersystem.Dataset, passthrough bool, idlist ...string) (map[string]status.Status, error) {
	result := map[string]status.Status{}
	unloaded := make([]string, 0, len(idlist))
	for _, v := range idlist {
		if !passthrough {
			st, ok := LoadStatus(dataset, v)
			if ok {
				result[v] = st
				continue
			}

		}
		unloaded = append(unloaded, v)
	}
	loaded, err := s.Service.LoadStatus(unloaded...)
	if err != nil {
		return nil, err
	}
	for k := range loaded {
		SetStatus(dataset, k, loaded[k])
		result[k] = loaded[k]
	}
	return result, nil
}
func (s *UserProfile) UpdateProfile(dataset usersystem.Dataset, id string, p *profile.Profile) error {
	err := s.Service.UpdateProfile(id, st)
	if err != nil {
		return err
	}
	if dataset != nil {
		dataset.Delete(DatatypeStatus, id)
	}
	return nil
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *UserProfile {
	p := New()
	err := s.InstallService(p)
	if err != nil {
		panic(err)
	}
	return p
}
