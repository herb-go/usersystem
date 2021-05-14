package usersystem

import (
	"context"

	"github.com/herb-go/herbsystem"
)

type UserSystem struct {
	Keyword Keyword
	*herbsystem.BasicSystem
	context context.Context
}

func (u *UserSystem) WithKeyword(k Keyword) *UserSystem {
	u.Keyword = k
	return u
}
func (u *UserSystem) SystemContext() context.Context {
	return u.context
}
func New() *UserSystem {
	s := &UserSystem{
		BasicSystem: herbsystem.New(),
	}
	s.context = context.WithValue(context.Background(), ContextKeyUsersystem, s)
	return s
}
