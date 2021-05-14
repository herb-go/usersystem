package sessionpayload

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type Builder interface {
	MustBuildPayloads(context.Context, usersystem.SessionType, string, *authority.Payloads)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}

type BuilderFunc func(context.Context, usersystem.SessionType, string, *authority.Payloads) error

func (f BuilderFunc) BuildPayloads(ctx context.Context, st usersystem.SessionType, uid string, p *authority.Payloads) error {
	return f(ctx, st, uid, p)
}

//Start start service
func (f BuilderFunc) Start() error {
	return nil
}

//Stop stop service
func (f BuilderFunc) Stop() error {
	return nil
}
