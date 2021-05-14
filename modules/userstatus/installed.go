package userstatus

import "github.com/herb-go/usersystem"

type InstalledUserStatus struct {
	*UserStatus
	UserSystem *usersystem.UserSystem
}

func NewInstalledUserStatus() *InstalledUserStatus {
	return &InstalledUserStatus{}
}
