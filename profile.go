package usersystem

import "github.com/herb-go/herb/user/profile"

type ProfileService interface {
	LoadProfile(...string) (map[string]*profile.Profile, error)
	UpdateProfile(string, *profile.Profile) error
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
