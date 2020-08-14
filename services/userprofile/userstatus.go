package userstatus

import (
	"github.com/herb-go/herb/user/status"
	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/userdataset"
	"github.com/herb-go/usersystem/userpurge"
	"github.com/herb-go/usersystem/usersession"
)

var ServiceName = "status"

var DatatypeStatus = usersystem.DataType("status")

func LoadStatus(ds usersystem.Dataset, id string) (status.Status, bool) {
	st, ok := ds.Get(DatatypeStatus, id)
	if !ok {
		return status.StatusUnkown, false
	}
	return st.(status.Status), true
}

func SetStatus(ds usersystem.Dataset, id string, st status.Status) {
	ds.Set(DatatypeStatus, id, st)
}

type UserStatus struct {
	herbsystem.NopService
	Service
}

func New() *UserStatus {
	return &UserStatus{}
}
func (s *UserStatus) InitService() error {
	return nil
}
func (s *UserStatus) ServiceName() string {
	return ServiceName
}
func (s *UserStatus) StartService() error {
	return s.Service.Start()
}
func (s *UserStatus) StopService() error {
	return s.Service.Stop()
}
func (s *UserStatus) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userdataset.InitDatasetTypeAction(DatatypeStatus),
		usersession.WrapCheckSession(s.CheckSession),
		userpurge.Wrap(s),
	}
}
func (s *UserStatus) CheckSession(session usersystem.Session, id string, payloads *authority.Payloads) (bool, error) {
	return s.IsUserAvaliable(id)
}
func (s *UserStatus) IsUserAvaliable(id string) (bool, error) {
	result, err := s.Service.LoadStatus(id)
	if err != nil {
		return false, err
	}
	st, ok := result[id]
	if !ok {
		return false, nil
	}
	return s.Service.IsAvailable(st)
}
func (s *UserStatus) LoadStatus(dataset usersystem.Dataset, passthrough bool, idlist ...string) (map[string]status.Status, error) {
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
func (s *UserStatus) UpdateStatus(dataset usersystem.Dataset, id string, st status.Status) error {
	err := s.Service.UpdateStatus(id, st)
	if err != nil {
		return err
	}
	if dataset != nil {
		dataset.Delete(DatatypeStatus, id)
	}
	return nil
}

func MustNewAndInstallTo(s *usersystem.UserSystem) *UserStatus {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
}
