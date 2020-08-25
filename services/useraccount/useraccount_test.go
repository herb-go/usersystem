package useraccount

import (
	"testing"

	"github.com/herb-go/user"
	"github.com/herb-go/usersystem"
)

type testService struct {
	accounts map[string]*user.Accounts
}

//Start start service
func (s *testService) Start() error {
	return nil
}

//Stop stop service
func (s *testService) Stop() error {
	return nil
}

//Account return account of give uid.
func (s *testService) Account(uid string) (*user.Accounts, error) {
	v, ok := s.accounts[uid]
	if !ok {
		return user.NewAccounts(), nil
	}
	return v, nil
}

//AccountToUID query uid by user account.
//Return user id and any error if raised.
//Return empty string as userid if account not found.
func (s *testService) AccountToUID(account *user.Account) (uid string, err error) {
	for k, v := range s.accounts {
		if v.Exists(account) {
			return k, nil
		}
	}
	return "", nil
}

//BindAccount bind account to user.
//Return any error if raised.
//If account exists,user.ErrAccountBindingExists should be rasied.
func (s *testService) BindAccount(uid string, account *user.Account) error {
	for _, v := range s.accounts {
		if v.Exists(account) {
			return user.ErrAccountBindingExists
		}
	}
	a, ok := s.accounts[uid]
	if !ok {
		a = user.NewAccounts()
		s.accounts[uid] = a
	}
	return a.Bind(account)
}

//UnbindAccount unbind account from user.
//Return any error if raised.
//If account not exists,user.ErrAccountUnbindingNotExists should be rasied.
func (s *testService) UnbindAccount(uid string, account *user.Account) error {
	return s.accounts[uid].Unbind(account)
}

func (t *testService) Purge(uid string) error {
	return nil
}
func newTestService() *testService {
	return &testService{
		accounts: map[string]*user.Accounts{},
	}
}
func TestUserAccount(t *testing.T) {
	var err error
	s := usersystem.New()
	ss := newTestService()
	useraccount := MustNewAndInstallTo(s)
	s.Ready()
	s.Configuring()
	useraccount.Service = ss
	s.Start()
	defer s.Stop()
	account, err := user.CaseSensitiveAcountProvider.NewAccount("test", "test")
	if err != nil {
		panic(err)
	}
	account2, err := user.CaseSensitiveAcountProvider.NewAccount("test", "test2")
	if err != nil {
		panic(err)
	}
	accounts, err := useraccount.Account("test")
	if err != nil {
		panic(err)
	}
	if len(accounts.Data()) != 0 {
		t.Fatal()
	}
	uid, err := useraccount.AccountToUID(account)
	if err != nil {
		panic(err)
	}
	if uid != "" {
		t.Fatal(uid)
	}
	err = useraccount.BindAccount("test", account)
	if err != nil {
		panic(err)
	}
	accounts, err = useraccount.Account("test")
	if err != nil {
		panic(err)
	}
	if len(accounts.Data()) != 1 {
		t.Fatal()
	}
	uid, err = useraccount.AccountToUID(account)
	if err != nil {
		panic(err)
	}
	if uid != "test" {
		t.Fatal(uid)
	}
	err = useraccount.BindAccount("test", account)
	if err != user.ErrAccountBindingExists {
		t.Fatal(err)
	}
	err = useraccount.UnbindAccount("test", account2)
	if err != user.ErrAccountUnbindingNotExists {
		t.Fatal(err)
	}
	err = useraccount.UnbindAccount("test", account)
	if err != nil {
		t.Fatal(err)
	}
	uid, err = useraccount.AccountToUID(account)
	if err != nil {
		panic(err)
	}
	if uid != "" {
		t.Fatal(uid)
	}
	err = useraccount.UnbindAccount("test", account)
	if err != user.ErrAccountUnbindingNotExists {
		t.Fatal(err)
	}
}
