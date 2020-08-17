package userprofile

import "github.com/herb-go/herb/user/profile"

type Service interface {
	GetProfile(id string) (*profile.Profile, error)
	UpdateProfile(id string, p *profile.Profile) error
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
