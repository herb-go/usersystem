package useraccount

import "github.com/herb-go/user"

type Service interface {
	//Accounts return accounts of give uid.
	MustAccounts(uid string) *user.Accounts
	//AccountToUID query uid by user account.
	//Return user id.
	//Return empty string as userid if account not found.
	MustAccountToUID(account *user.Account) (uid string)
	//BindAccount bind account to user.
	//If account exists,user.ErrAccountBindingExists should be rasied.
	MustBindAccount(uid string, account *user.Account)
	//UnbindAccount unbind account from user.
	//If account not exists,user.ErrAccountUnbindingNotExists should be rasied.
	MustUnbindAccount(uid string, account *user.Account)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
