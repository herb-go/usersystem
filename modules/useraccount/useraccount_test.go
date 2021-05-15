package useraccount

import (
	"testing"

	"github.com/herb-go/herbsystem"

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
func (s *testService) MustAccounts(uid string) *user.Accounts {
	v, ok := s.accounts[uid]
	if !ok {
		return user.NewAccounts()
	}
	return v
}

//AccountToUID query uid by user account.
//Return empty string as userid if account not found.
func (s *testService) MustAccountToUID(account *user.Account) (uid string) {
	for k, v := range s.accounts {
		if v.Exists(account) {
			return k
		}
	}
	return ""
}

//BindAccount bind account to user.
//If account exists,user.ErrAccountBindingExists should be rasied.
func (s *testService) MustBindAccount(uid string, account *user.Account) {
	for _, v := range s.accounts {
		if v.Exists(account) {
			panic(user.ErrAccountBindingExists)
		}
	}
	a, ok := s.accounts[uid]
	if !ok {
		a = user.NewAccounts()
		s.accounts[uid] = a
	}
	err := a.Bind(account)
	if err != nil {
		panic(err)
	}
}

//UnbindAccount unbind account from user.
//If account not exists,user.ErrAccountUnbindingNotExists should be rasied.
func (s *testService) MustUnbindAccount(uid string, account *user.Account) {
	err := s.accounts[uid].Unbind(account)
	if err != nil {
		panic(err)
	}
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
	herbsystem.MustReady(s)
	useraccount.Service = ss
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != useraccount.UserAccount {
		t.Fatal()
	}
	herbsystem.MustStart(s)
	defer herbsystem.MustStop(s)
	account, err := user.CaseSensitiveAcountProvider.NewAccount("test", "test")
	if err != nil {
		panic(err)
	}
	account2, err := user.CaseSensitiveAcountProvider.NewAccount("test", "test2")
	if err != nil {
		panic(err)
	}
	accounts := useraccount.MustAccounts("test")
	if len(accounts.Data()) != 0 {
		t.Fatal()
	}
	uid := useraccount.MustAccountToUID(account)
	if uid != "" {
		t.Fatal(uid)
	}
	useraccount.MustBindAccount("test", account)

	accounts = useraccount.MustAccounts("test")

	if len(accounts.Data()) != 1 {
		t.Fatal()
	}
	uid = useraccount.MustAccountToUID(account)
	if err != nil {
		panic(err)
	}
	if uid != "test" {
		t.Fatal(uid)
	}
	useraccount.MustBindAccount("test", account)
	if err != user.ErrAccountBindingExists {
		t.Fatal(err)
	}
	useraccount.MustUnbindAccount("test", account2)
	if err != user.ErrAccountUnbindingNotExists {
		t.Fatal(err)
	}
	useraccount.MustUnbindAccount("test", account)
	if err != nil {
		t.Fatal(err)
	}
	uid = useraccount.MustAccountToUID(account)
	if uid != "" {
		t.Fatal(uid)
	}
	useraccount.MustUnbindAccount("test", account)
}

func TestMustGetModule(t *testing.T) {
	s := usersystem.New()
	herbsystem.MustReady(s)
	herbsystem.MustConfigure(s)
	if MustGetModule(s) != nil {
		t.Fatal()
	}
}
