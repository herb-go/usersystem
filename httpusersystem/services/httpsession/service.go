package httpsession

import "github.com/herb-go/usersystem"

type Service interface {
	SessionType() usersystem.SessionType
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
