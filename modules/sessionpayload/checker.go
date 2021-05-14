package sessionpayload

import (
	"context"

	"github.com/herb-go/usersystem"
)

type Checker interface {
	MustCheckSession(ctx context.Context, session *usersystem.Session) bool
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}

type CheckerFunc func(ctx context.Context, session *usersystem.Session) bool

func (f CheckerFunc) MustCheckSession(ctx context.Context, session *usersystem.Session) bool {
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
