package userterm

type Service interface {
	CurrentTerm(uid string) (string, error)
	StartNewTerm(uid string) (string, error)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
