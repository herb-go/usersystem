package usercreate

import (
	"context"

	"github.com/herb-go/herbsystem"
	"github.com/herb-go/usersystem"
)

var CommandExist = herbsystem.Command("exist")

type contextkey string

var contextkeyResult = contextkey("result")

type result struct {
	exists bool
}

func WrapExist(h func(string) (bool, error)) *herbsystem.Action {
	a := herbsystem.NewAction()
	a.Command = CommandExist
	a.Handler = func(ctx context.Context, next func(context.Context) error) error {

		ok, err := h(usersystem.GetUID(ctx))
		if err != nil {
			return err
		}
		if ok {
			r := ctx.Value(contextkeyResult).(*result)
			r.exists = true
			return nil
		}
		return next(ctx)
	}
	return a
}

func ExecExist(s *usersystem.UserSystem, uid string) (bool, error) {
	r := &result{exists: false}
	ctx := usersystem.UIDContext(s.Context, uid)
	ctx = context.WithValue(ctx, contextkeyResult, r)
	_, err := s.System.ExecActions(ctx, CommandExist)
	if err != nil {
		return false, err
	}
	return r.exists, nil
}
