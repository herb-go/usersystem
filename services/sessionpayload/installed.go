package sessionpayload

import "github.com/herb-go/usersystem"

type InstalledSessionPayload struct {
	*SessionPayload
	UserSystem *usersystem.UserSystem
}

func NewInstalledSessionPayload() *InstalledSessionPayload {
	return &InstalledSessionPayload{}
}
