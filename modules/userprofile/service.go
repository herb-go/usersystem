package userprofile

import "github.com/herb-go/user/profile"

type Service interface {
	MustGetProfile(id string) *profile.Profile
	MustUpdateProfile(id string, p *profile.Profile)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
