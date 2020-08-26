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
	//Purge purge user data cache
	Purge(string) error
}
