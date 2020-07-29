package usersystem

import (
	"github.com/herb-go/herbsecurity/authority/service/token"
)

type Service struct {
	Keyword         Keyword
	SourceService   SourceService
	AccountsService AccountsService
	ProfileService  ProfileService
	RolesService    RolesService
	StatusService   StatusService
	PasswordService PasswordService
	TokenService    token.Manager
}
