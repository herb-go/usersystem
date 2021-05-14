package userstatus

import "github.com/herb-go/user/status"

type Service interface {
	MustLoadStatus(string) status.Status
	MustUpdateStatus(string, status.Status)
	MustCreateStatus(string)
	MustRemoveStatus(string)
	MustListUsersByStatus(last string, limit int, reverse bool, statuses ...status.Status) []string
	status.Service
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
