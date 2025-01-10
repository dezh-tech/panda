package userhandler

import "github.com/dezh-tech/geb/service/user"

type Handler struct {
	userSvc user.Service
}

func New(userSvc user.Service) Handler {
	return Handler{
		userSvc: userSvc,
	}
}
