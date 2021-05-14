package userpassword

import "github.com/herb-go/usersystem"

type InstalledUserPassword struct {
	*UserPassword
	UserSystem *usersystem.UserSystem
}

func NewInstalledUserPassword() *InstalledUserPassword {
	return &InstalledUserPassword{}
}
