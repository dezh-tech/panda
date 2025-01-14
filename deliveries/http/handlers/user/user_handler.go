package handlers

import service "github.com/dezh-tech/panda/services/user"

type User struct {
	service service.User
}

func NewUserService(UserSvc service.User) User {
	return User{
		service: UserSvc,
	}
}
