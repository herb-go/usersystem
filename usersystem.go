package usersystem

import (
	"context"

	"github.com/herb-go/herbsystem"
)

type UserSystem struct {
	Context context.Context
	*herbsystem.System
}

func New() *UserSystem {
	s := &UserSystem{
		System: herbsystem.NewSystem(),
	}
	s.Context = context.WithValue(context.Background(), ContextKeyUsersystem, s)
	return s
}
