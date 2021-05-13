package usersystem

import (
	"context"

	"github.com/herb-go/herbsystem"
)

type UserSystem struct {
	Keyword Keyword
	*herbsystem.BasicSystem
	Context context.Context
}

func (u *UserSystem) WithKeyword(k Keyword) *UserSystem {
	u.Keyword = k
	return u
}

func New() *UserSystem {
	s := &UserSystem{
		BasicSystem: herbsystem.New(),
	}
	s.Context = context.WithValue(context.Background(), ContextKeyUsersystem, s)
	return s
}
