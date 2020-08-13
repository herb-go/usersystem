package usersession

import (
	"testing"
	"time"

	"github.com/herb-go/usersystem"
)

func TestActiveSessionManagerConfig(t *testing.T) {
	s := usersystem.New()
	s.InstallService(&testService{})
	s.Ready()
	s.Configuring()
	s.Start()
	defer s.Stop()

	config, err := ExecActiveSessionManagerConfig(s, "test")
	if err != nil {
		t.Fatal(err)
	}
	if config.Supported != true || config.Duration != time.Minute {
		t.Fatal(config)
	}
	config, err = ExecActiveSessionManagerConfig(s, "notexist")
	if err != nil {
		t.Fatal(err)
	}
	if config.Supported != false {
		t.Fatal(config)
	}
}
