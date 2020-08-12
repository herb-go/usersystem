package userstatus

import (
	"github.com/herb-go/herb/user/status"
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
	"github.com/herb-go/usersystem/useravaliable"
	"github.com/herb-go/usersystem/userdataset"
)

var ServiceName = "status"

var DatatypeStatus = usersystem.Datatype("status")

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

type UserStauts struct {
	herbsystem.NopService
	StatusService
}

func New() *UserStauts {
	return &UserStauts{}
}
func (s *UserStauts) InitService() error {
	return nil
}
func (s *UserStauts) ServiceName() string {
	return ServiceName
}
func (s *UserStauts) StartService() error {
	return s.StatusService.Start()
}
func (s *UserStauts) StopService() error {
	return s.StatusService.Stop()
}
func (s *UserStauts) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		userdataset.InitDatasetTypeAction(DatatypeStatus),
		useravaliable.Wrap(s.IsUserAvaliable),
	}
}
func (s *UserStauts) IsUserAvaliable(id string) (bool, error) {
	result, err := s.StatusService.LoadStatus(id)
	if err != nil {
		return false, err
	}
	st, ok := result[id]
	if !ok {
		return false, nil
	}
	return s.StatusService.IsAvailable(st)
}
func (s *UserStauts) LoadStatus(dataset usersystem.Dataset, passthrough bool, idlist ...string) (map[string]status.Status, error) {
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
	loaded, err := s.StatusService.LoadStatus(unloaded...)
	if err != nil {
		return nil, err
	}
	for k := range loaded {
		SetStatus(dataset, k, loaded[k])
		result[k] = loaded[k]
	}
	return result, nil
}
func (s *UserStauts) UpdateStatus(dataset usersystem.Dataset, id string, st status.Status) error {
	err := s.StatusService.UpdateStatus(id, st)
	if err != nil {
		return err
	}
	if dataset != nil {
		dataset.Delete(DatatypeStatus, id)
	}
	return nil
}

func MustNewUserstatus(s *usersystem.UserSystem) *UserStauts {
	status := New()
	err := s.InstallService(status)
	if err != nil {
		panic(err)
	}
	return status
}
