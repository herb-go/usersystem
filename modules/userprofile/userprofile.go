package userprofile

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/user/profile"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userpurge"
)

var ModuleName = "profile"

type UserProfile struct {
	herbsystem.NopModule
	Services Services
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
	if err := s.Services.Start(); err != nil {
		panic(err)
	}
	next(ctx, system)
}
func (s *UserProfile) StopProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	if err := s.Services.Stop(); err != nil {
		system.LogSystemError(err)
	}
	next(ctx, system)
}
func (s *UserProfile) PurgeProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	uid := usersystem.GetUID(ctx)
	if err := s.Services.Purge(uid); err != nil {
		system.LogSystemError(err)
	}
	next(ctx, system)

}
func (s *UserProfile) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {

	system.MountSystemActions(
		herbsystem.CreateStartAction(s.StartProcess),
		herbsystem.CreateStopAction(s.StopProcess),
		herbsystem.CreateAction(userpurge.Command, s.PurgeProcess),
	)
	next(ctx, system)
}

func (s *UserProfile) loadProfiles(idlist ...string) map[string]*profile.Profile {
	var result = map[string]*profile.Profile{}
	for _, id := range idlist {
		result[id] = s.Services.MustGetProfile(id)
	}
	return result
}
func (s *UserProfile) MustLoadProfile(uid string) *profile.Profile {
	r := s.loadProfiles(uid)
	return r[uid]
}
func (s *UserProfile) MustLoadProfiles(idlist ...string) map[string]*profile.Profile {
	result := map[string]*profile.Profile{}
	loaded := s.loadProfiles(idlist...)
	for k := range loaded {
		result[k] = loaded[k]
	}
	return result
}
func (s *UserProfile) MustUpdateProfile(id string, p *profile.Profile) {
	s.Services.MustUpdateProfile(id, p)
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
