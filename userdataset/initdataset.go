package userdataset

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var InitDatasetCommand = herbsystem.Command("initdataset")

func WrapInitDataset(h func(ds usersystem.Dataset)) *herbsystem.Action {
	return herbsystem.CreateAction(InitDatasetCommand, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		h(GetDataset(ctx))
		next(ctx, system)
	})
}

func InitDatasetTypeAction(dt usersystem.DataType) *herbsystem.Action {
	return WrapInitDataset(func(ds usersystem.Dataset) {
		ds.InitType(dt)
	})
}

func MustExecInitDataset(s *usersystem.UserSystem, ds usersystem.Dataset) {
	herbsystem.MustExecActions(WithDataset(s.SystemContext(), ds), s, InitDatasetCommand)
}
