package service

import (
	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/dezh-tech/panda/repositories"
)

type Domain struct {
	repo      *repositories.Domain
	validator *validator.Validator
}

func NewDomainService(repo *repositories.Domain) Domain {
	return Domain{repo: repo, validator: validator.NewValidator()}
}
