package userdataset

import (
	"context"
	"testing"

	"github.com/herb-go/herbsystem"

	"github.com/herb-go/usersystem"
)

var testDatatype = usersystem.DataType("test")

type testModule struct {
	herbsystem.NopModule
}

func (s *testModule) InstallProcess(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
	system.MountSystemActions(
		InitDatasetTypeAction(testDatatype),
		WrapNewDataset(func(s *usersystem.UserSystem) usersystem.Dataset {
			ds := usersystem.NewPlainDataset()
			MustExecInitDataset(s, ds)
			return ds
		}),
	)
	next(ctx, system)
}

func TestInit(t *testing.T) {
	s := usersystem.New()
	s.MustRegisterSystemModule(&testModule{})
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	ds := MustExecNewDataset(s)
	pds := ds.(*usersystem.PlainDataset)
	_, ok := pds.Dataset[testDatatype]
	if !ok {
		t.Fatal(pds)
	}
}
func TestDefaultUserDataset(t *testing.T) {
	s := usersystem.New()
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	ds := MustExecNewDataset(s)
	if ds == nil || ds.(*usersystem.PlainDataset) == nil {
		t.Fatal(ds)
	}
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
