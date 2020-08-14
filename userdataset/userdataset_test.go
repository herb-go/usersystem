package userdataset

import (
	"context"
	"testing"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem"
)

var testDatatype = usersystem.DataType("test")

type testService struct {
	herbsystem.NopService
}

func (s *testService) ServiceActions() []*herbsystem.Action {
	return []*herbsystem.Action{
		InitDatasetTypeAction(testDatatype),
		WrapNewDataset(func(s *usersystem.UserSystem) (usersystem.Dataset, error) {
			ds := usersystem.NewPlainDataset()
			err := ExecInitDataset(s, ds)
			if err != nil {
				return nil, err
			}
			return ds, nil
		}),
	}
}

func TestInit(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	ds, err := ExecNewDataset(s)
	if err != nil {
		panic(err)
	}
	pds := ds.(*usersystem.PlainDataset)
	_, ok := pds.Dataset[testDatatype]
	if !ok {
		t.Fatal(pds)
	}
	defer s.Stop()
}
func TestDefaultUserDataset(t *testing.T) {
	s := usersystem.New()
	s.Ready()
	s.Configuring()
	s.Start()
	ds, err := ExecNewDataset(s)
	if err != nil {
		panic(err)
	}
	if ds == nil || ds.(*usersystem.PlainDataset) == nil {
		t.Fatal(ds)
	}
	defer s.Stop()
}

func TestUtil(t *testing.T) {
	if GetDataset(context.Background()) != nil {
		t.Fatal()
	}
	ds := usersystem.NewPlainDataset()
	if GetDataset(WithDataset(context.Background(), ds)) != ds {
		t.Fatal()
	}
}
