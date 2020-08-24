package userrole

import "github.com/herb-go/usersystem"

type InstalledUserRole struct {
	*UserRole
	UserSystem *usersystem.UserSystem
}

func NewInstalledUserRole() *InstalledUserRole {
	return &InstalledUserRole{}
}
