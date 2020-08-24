package userprofile

import "github.com/herb-go/usersystem"

type InstalledUserProfile struct {
	*UserProfile
	UserSystem *usersystem.UserSystem
}

func NewInstalledUserProfile() *InstalledUserProfile {
	return &InstalledUserProfile{}
}
