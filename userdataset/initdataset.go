package userdataset

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var InitDatasetCommand = herbsystem.Command("initdataset")

func WrapInitDataset(h func(ds usersystem.Dataset) error) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = InitDatasetCommand
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		err := h(GetDataset(ctx))
		if err != nil {
			return err
		}
		return next(ctx)
	}
	return a
}

func InitDatasetTypeAction(dt usersystem.DataType) *herbsystem.Action {
	return WrapInitDataset(func(ds usersystem.Dataset) error {
		ds.InitType(dt)
		return nil
	})
}

func ExecInitDataset(s *usersystem.UserSystem, ds usersystem.Dataset) error {
	_, err := s.System.ExecActions(WithDataset(s.Context, ds), InitDatasetCommand)
	return err
}
