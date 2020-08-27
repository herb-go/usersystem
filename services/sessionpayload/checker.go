package sessionpayload

import (
	"context"

	"github.com/herb-go/usersystem"
)

type Checker interface {
	CheckSession(ctx context.Context, session *usersystem.Session) (bool, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}

type CheckerFunc func(ctx context.Context, session *usersystem.Session) (bool, error)

func (f CheckerFunc) CheckSession(ctx context.Context, session *usersystem.Session) (bool, error) {
	return f(ctx, session)
}

//Start start service
func (f CheckerFunc) Start() error {
	return nil
}

//Stop stop service
func (f CheckerFunc) Stop() error {
	return nil
}
