package http

import (
	"fmt"

	domainhandler "github.com/dezh-tech/panda/deliveries/http/handlers/domain_handler"
	domainService "github.com/dezh-tech/panda/services/domain"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router   *echo.Echo
	config   Config
	handlers HttpHandlers
}

func New(config Config, userSvc domainService.DomainService) Server {
	return Server{
		Router: echo.New(),
		config: config,

		handlers: HttpHandlers{
			user: domainhandler.New(userSvc),
		},
	}
}

func (s Server) Start() error {
	s.handlers.Start(s.Router)

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
