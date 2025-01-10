package http

import (
	"fmt"

	userh "github.com/dezh-tech/geb/delivery/http/user_handler"
	users "github.com/dezh-tech/geb/service/user"
	"github.com/labstack/echo/v4"
)

type Server struct {
	config      Config
	userHandler userh.Handler
	Router      *echo.Echo
}

func New(config Config, userSvc users.Service) Server {
	return Server{
		Router:      echo.New(),
		config:      config,
		userHandler: userh.New(userSvc),
	}
}

func (s Server) Start() error {
	s.userHandler.SetRoutes(s.Router)

	address := fmt.Sprintf(":%d", s.config.Port)
	if err := s.Router.Start(address); err != nil {
		return err
	}

	return nil
}

func (s Server) Stop() error {
	if err := s.Router.Close(); err != nil {
		return err
	}

	return nil
}
