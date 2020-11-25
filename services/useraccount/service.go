package useraccount

import "github.com/herb-go/user"

type Service interface {
	//Accounts return accounts of give uid.
	Accounts(uid string) (*user.Accounts, error)
	//AccountToUID query uid by user account.
	//Return user id and any error if raised.
	//Return empty string as userid if account not found.
	AccountToUID(account *user.Account) (uid string, err error)
	//BindAccount bind account to user.
	//Return any error if raised.
	//If account exists,user.ErrAccountBindingExists should be rasied.
	BindAccount(uid string, account *user.Account) error
	//UnbindAccount unbind account from user.
	//Return any error if raised.
	//If account not exists,user.ErrAccountUnbindingNotExists should be rasied.
	UnbindAccount(uid string, account *user.Account) error
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
