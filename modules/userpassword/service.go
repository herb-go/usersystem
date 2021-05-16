package userpassword

type Service interface {
	MustVerifyPassword(uid string, password string) bool
	//PasswordChangeable return password changeable
	PasswordChangeable() bool
	//UpdatePassword update user password
	MustUpdatePassword(uid string, password string)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
