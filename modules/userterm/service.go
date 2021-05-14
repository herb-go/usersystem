package userterm

type Service interface {
	MustCurrentTerm(uid string) string
	MustStartNewTerm(uid string) string
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}
