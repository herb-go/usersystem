package userterm

import "github.com/herb-go/usersystem"

type InstalledUserTerm struct {
	*UserTerm
	UserSystem *usersystem.UserSystem
}

func NewInstalledUserTerm() *InstalledUserTerm {
	return &InstalledUserTerm{}
}
