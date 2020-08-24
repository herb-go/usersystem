package usersystem

import (
	"context"

	"github.com/herb-go/herbsystem"
)

type UserSystem struct {
	Keyword Keyword
	Context context.Context
	*herbsystem.System
}

func (u *UserSystem) WithKeyword(k Keyword) *UserSystem {
	u.Keyword = k
	return u
}

func New() *UserSystem {
	s := &UserSystem{
		System: herbsystem.NewSystem(),
	}
	s.Context = context.WithValue(context.Background(), ContextKeyUsersystem, s)
	return s
}
