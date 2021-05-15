package userstatus

import "github.com/herb-go/user/status"

type Service interface {
	//MustLoadStatus load user status
	//Return uid status and exists
	//Status.StatusUnkown and false will be retruned if user not exists
	MustLoadStatus(string) (status.Status, bool)
	//MustUpdateStatus update user status.
	MustUpdateStatus(string, status.Status)
	//MustCreateStatus create user status.
	MustCreateStatus(string)
	//MustRemoveStatus remove user status
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
