package userstatus

import "github.com/herb-go/user/status"

type Service interface {
	LoadStatus(string) (status.Status, error)
	UpdateStatus(string, status.Status) error
	CreateStatus(string) error
	RemoveStatus(string) error
	ListUsersByStatus(last string, limit int, statuses ...status.Status) ([]string, bool, error)
	status.Service
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
