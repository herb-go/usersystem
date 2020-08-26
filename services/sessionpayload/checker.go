package sessionpayload

import "github.com/herb-go/usersystem"

type Checker interface {
	CheckSession(*usersystem.Session) (bool, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
