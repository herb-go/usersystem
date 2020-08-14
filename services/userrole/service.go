package userrole

import "github.com/herb-go/herbsecurity/authorize/role"

type Service interface {
	Roles(uid string) (*role.Roles, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
