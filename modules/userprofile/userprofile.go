package userprofile

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/user/profile"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userdataset"
	"github.com/herb-go/usersystem/userpurge"
)

var ModuleName = "profile"

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
	herbsystem.NopModule
	Services []Service
}

func New() *UserProfile {
	return &UserProfile{}
}

func (s *UserProfile) ModuleName() string {
	return ModuleName
}
func (s *UserProfile) InitModule() {
	s.Services = []Service{}
}

func (s *UserProfile) StartProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	for _, v := range s.Services {
		if err := v.Start(); err != nil {
			panic(err)
		}
	}
	next(ctx, system)
}
func (s *UserProfile) StopProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	for _, v := range s.Services {
		if err := v.Start(); err != nil {
			system.LogSystemError(err)
		}
	}
	next(ctx, system)
}
func (s *UserProfile) PurgeProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	uid := usersystem.GetUID(ctx)
	for _, v := range s.Services {
		if err := v.Purge(uid); err != nil {
			system.LogSystemError(err)
		}
	}
	next(ctx, system)

}
func (s *UserProfile) InitProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {

	system.MountSystemActions(
		herbsystem.CreateStartAction(s.StartProcess),
		herbsystem.CreateStopAction(s.StopProcess),
		userdataset.InitDatasetTypeAction(DatatypeProfile),
		herbsystem.CreateAction(userpurge.Command, s.PurgeProcess),
	)
	next(ctx, system)
}

func (s *UserProfile) loadProfiles(idlist ...string) map[string]*profile.Profile {
	var result = map[string]*profile.Profile{}
	for _, id := range idlist {
		p := profile.NewProfile()
		for _, v := range s.Services {
			fields := v.MustGetProfile(id)
			p.Chain(fields)
		}
		result[id] = p
	}
	return result
}
func (s *UserProfile) MustLoadProfile(uid string) *profile.Profile {
	r := s.loadProfiles(uid)
	return r[uid]
}
func (s *UserProfile) MustLoadProfiles(dataset usersystem.Dataset, passthrough bool, idlist ...string) map[string]*profile.Profile {
	result := map[string]*profile.Profile{}
	unloaded := make([]string, 0, len(idlist))
	for _, v := range idlist {
		if !passthrough {
			p, ok := LoadProfile(dataset, v)
			if ok {
				result[v] = p
				continue
			}

		}
		unloaded = append(unloaded, v)
	}
	loaded := s.loadProfiles(unloaded...)
	for k := range loaded {
		SetProfile(dataset, k, loaded[k])
		result[k] = loaded[k]
	}
	return result
}
func (s *UserProfile) MustUpdateProfile(dataset usersystem.Dataset, id string, p *profile.Profile) {
	for _, v := range s.Services {
		v.MustUpdateProfile(id, p)
	}
	if dataset != nil {
		dataset.Delete(DatatypeProfile, id)
	}
}
func (s *UserProfile) AppendService(service Service) {
	s.Services = append(s.Services, service)
}
func MustNewAndInstallTo(s *usersystem.UserSystem) *InstalledUserProfile {
	p := New()
	s.MustRegisterSystemModule(p)
	i := NewInstalledUserProfile()
	i.UserProfile = p
	i.UserSystem = s
	return i
}

func MustGetModule(s *usersystem.UserSystem) *UserProfile {
	v := herbsystem.MustGetConfigurableModule(s, ModuleName)
	if v == nil {
		return nil
	}
	return v.(*UserProfile)
}
