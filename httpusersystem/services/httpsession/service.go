package httpsession

type Service interface {
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
}
