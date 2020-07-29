package usersystem

type User struct {
	users *Users
	ID    string
}

type Users struct {
	Service *Service
	source  Source
}

func NewUsers() *Users {
	return &Users{}
}
func (u *Users) User(id string) *User {
	return &User{
		users: u,
		ID:    id,
	}
}
