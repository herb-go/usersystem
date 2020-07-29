package usersystem

import (
	"sync"

	"github.com/herb-go/herb/user"
	"github.com/herb-go/herb/user/profile"
	"github.com/herb-go/herb/user/status"
	"github.com/herb-go/herbsecurity/authorize/role"
)

type Source interface {
	LoadProfile(service *Service, idlist ...string) (map[string]*profile.Profile, error)
	SetProfile(map[string]*profile.Profile) error
	LoadAccounts(service *Service, idlist ...string) (map[string]*user.Accounts, error)
	SetAccounts(map[string]*user.Accounts) error
	LoadRoles(service *Service, idlist ...string) (map[string]*role.Roles, error)
	SetRoles(map[string]status.Status) error
	LoadStatus(service *Service, idlist ...string) (map[string]status.Status, error)
	SetStatus(map[string]status.Status) error
}

type CachedSource struct {
	profilemap  sync.Map
	accountsmap sync.Map
	rolesmap    sync.Map
	statusmap   sync.Map
}

func (s *CachedSource) LoadProfile(service *Service, idlist ...string) (map[string]*profile.Profile, error) {
	result := map[string]*profile.Profile{}
	for k := range idlist {
		v, ok := s.profilemap.Load(idlist[k])
		if ok {
			result[idlist[k]] = v.(*profile.Profile)
		}
	}
	return result, nil
}
func (s *CachedSource) SetProfile(m map[string]*profile.Profile) error {
	for k := range m {
		s.profilemap.Store(k, m[k])
	}
	return nil
}
func (s *CachedSource) LoadAccounts(service *Service, idlist ...string) (map[string]*user.Accounts, error) {
	result := map[string]*user.Accounts{}
	for k := range idlist {
		v, ok := s.accountsmap.Load(idlist[k])
		if ok {
			result[idlist[k]] = v.(*user.Accounts)
		}
	}
	return result, nil
}
func (s *CachedSource) SetAccounts(m map[string]*user.Accounts) error {
	for k := range m {
		s.accountsmap.Store(k, m[k])
	}
	return nil
}
func (s *CachedSource) LoadRoles(service *Service, idlist ...string) (map[string]*role.Roles, error) {
	result := map[string]*role.Roles{}
	for k := range idlist {
		v, ok := s.rolesmap.Load(idlist[k])
		if ok {
			result[idlist[k]] = v.(*role.Roles)
		}
	}
	return result, nil
}
func (s *CachedSource) SetRoles(m map[string]status.Status) error {
	for k := range m {
		s.rolesmap.Store(k, m[k])
	}
	return nil
}
func (s *CachedSource) LoadStatus(service *Service, idlist ...string) (map[string]status.Status, error) {
	result := map[string]status.Status{}
	for k := range idlist {
		v, ok := s.statusmap.Load(idlist[k])
		if ok {
			result[idlist[k]] = v.(status.Status)
		}
	}
	return result, nil

}
func (s *CachedSource) SetStatus(m map[string]status.Status) error {
	for k := range m {
		s.statusmap.Store(k, m[k])
	}
	return nil
}

func NewCachedSource() *CachedSource {
	return &CachedSource{}
}

type SourceService interface {
	CreateSource() (Source, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}

type SourceServiceFunc func() (Source, error)

func (f SourceServiceFunc) CreateSource() (Source, error) {
	return f()
}

func (f SourceServiceFunc) Start() error {
	return nil
}
func (f SourceServiceFunc) Stop() error {
	return nil
}

var CachedSourceService = SourceServiceFunc(func() (Source, error) {
	return NewCachedSource(), nil
})
