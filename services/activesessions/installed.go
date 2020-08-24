package activesessions

import "github.com/herb-go/usersystem"

type InstalledManager struct {
	*ActiveSessions
	UserSystem *usersystem.UserSystem
}

func NewInstalledManager() *InstalledManager {
	return &InstalledManager{}
}
