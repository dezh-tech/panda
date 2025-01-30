package service

import (
	"github.com/dezh-tech/panda/pkg/validator"
	"github.com/dezh-tech/panda/repositories"
)

type User struct {
	repo      *repositories.User
	validator *validator.Validator
}

func NewUserService(repo *repositories.User) User {
	return User{repo: repo, validator: validator.NewValidator()}
}
