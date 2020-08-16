package userstatus

import "github.com/herb-go/herb/user/status"

type Service interface {
	LoadStatus(string) (status.Status, error)
	UpdateStatus(string, status.Status) error
	CreateStatus(string) error
	status.Service
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
