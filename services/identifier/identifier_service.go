package service

import (
	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/dezh-tech/panda/repositories"
	domainService "github.com/dezh-tech/panda/services/domain"
	userService "github.com/dezh-tech/panda/services/user"
)

type Identifier struct {
	repo          *repositories.Identifier
	domainService domainService.Domain
	userService   userService.User
	validator     *validator.Validator
}

func NewIdentifierService(repo *repositories.Identifier, domainService domainService.Domain, userService userService.User) Identifier {
	return Identifier{repo: repo, validator: validator.NewValidator(), domainService: domainService, userService: userService}
}
