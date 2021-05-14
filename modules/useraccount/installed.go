package useraccount

import "github.com/herb-go/usersystem"

type InstalledUserAccount struct {
	*UserAccount
	UserSystem *usersystem.UserSystem
}

func NewInstalledUserAccount() *InstalledUserAccount {
	return &InstalledUserAccount{}
}
