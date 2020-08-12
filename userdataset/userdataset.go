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

func WrapNewDataset(h func(s *usersystem.UserSystem) (usersystem.Dataset, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = NewDatasetCommand
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {
		ds, err := h(usersystem.GetUsersystem(ctx))
		if err != nil {
			return err
		}
		ctx = WithDataset(ctx, ds)
		return next(ctx)
	}
	return a
}

func NewDefaultDataset(s *usersystem.UserSystem) (usersystem.Dataset, error) {
	ds := usersystem.NewPlainDataset()
	err := ExecInitDataset(s, ds)
	if err != nil {
		return nil, err
	}
	return ds, nil
}

func ExecNewDataset(s *usersystem.UserSystem) (usersystem.Dataset, error) {
	ctx, err := s.System.ExecActions(s.Context, NewDatasetCommand)
	if err != nil {
		return nil, err
	}
	v := ctx.Value(ContextKeyDataset)
	ds, ok := v.(usersystem.Dataset)
	if !ok {
		return NewDefaultDataset(s)
	}
	return ds, nil
}
