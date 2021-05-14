package userpassword

type Service interface {
	VerifyPassword(uid string, password string) (bool, error)
	//PasswordChangeable return password changeable
	PasswordChangeable() bool
	//UpdatePassword update user password
	//Return any error if raised
	UpdatePassword(uid string, password string) error
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
