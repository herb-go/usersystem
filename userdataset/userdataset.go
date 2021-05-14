package userdataset

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var ContextKeyDataset = usersystem.ContextKey("dataset")

func WithDataset(ctx context.Context, dataset usersystem.Dataset) context.Context {
	return context.WithValue(ctx, ContextKeyDataset, dataset)
}

func GetDataset(ctx context.Context) usersystem.Dataset {
	v := ctx.Value(ContextKeyDataset)
	if v == nil {
		return nil
	}
	return v.(usersystem.Dataset)
}

var NewDatasetCommand = herbsystem.Command("newdataset")

func WrapNewDataset(h func(s *usersystem.UserSystem) usersystem.Dataset) *herbsystem.Action {
	return herbsystem.CreateAction(NewDatasetCommand, func(ctx context.Context, system herbsystem.System, next func(context.Context, herbsystem.System)) {
		ds := h(system.(*usersystem.UserSystem))
		ctx = WithDataset(ctx, ds)
		next(ctx, system)
	})
}

func NewDefaultDataset(s *usersystem.UserSystem) usersystem.Dataset {
	ds := usersystem.NewPlainDataset()
	MustExecInitDataset(s, ds)
	return ds
}

func MustExecNewDataset(s *usersystem.UserSystem) usersystem.Dataset {
	ctx := herbsystem.MustExecActions(s.SystemContext(), s, NewDatasetCommand)
	v := ctx.Value(ContextKeyDataset)
	ds, ok := v.(usersystem.Dataset)
	if !ok {
		return NewDefaultDataset(s)
	}
	return ds
}
