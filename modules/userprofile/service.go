package userprofile

import (
	"github.com/herb-go/herbsystem"
	"github.com/herb-go/user/profile"
)

type Service interface {
	MustGetProfile(id string) *profile.Profile
	MustUpdateProfile(id string, p *profile.Profile)
	//Start start service
	Start() error
	//Stop stop service
	Stop() error
	//Purge purge user data cache
	Purge(string) error
}

type Services []Service

func (s Services) MustGetProfile(id string) *profile.Profile {
	p := profile.NewProfile()
	for _, v := range s {
		fields := v.MustGetProfile(id)
		p.Chain(fields)
	}
	return p
}
func (s Services) MustUpdateProfile(id string, p *profile.Profile) {
	for _, v := range s {
		v.MustUpdateProfile(id, p)
	}
}

//Start start service
func (s Services) Start() error {
	for _, v := range s {
		if err := v.Start(); err != nil {
			return err
		}
	}
	return nil
}

//Stop stop service
func (s Services) Stop() error {
	var errs = herbsystem.NewErrors()
	for _, v := range s {
		if err := v.Start(); err != nil {
			errs = errs.Append(err)
		}
	}
	return errs.NilOrError()
}

//Purge purge user data cache
func (s Services) Purge(uid string) error {
	var errs = herbsystem.NewErrors()
	for _, v := range s {
		if err := v.Purge(uid); err != nil {
			errs = errs.Append(err)
		}
	}
	return errs.NilOrError()
}
