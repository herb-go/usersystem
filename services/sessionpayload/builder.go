package sessionpayload

import (
	"context"

	"github.com/herb-go/herbsecurity/authority"
	"github.com/herb-go/usersystem"
)

type Builder interface {
	BuildPayloads(context.Context, usersystem.SessionType, string, *authority.Payloads) error
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
