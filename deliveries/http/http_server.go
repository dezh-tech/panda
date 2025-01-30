package http

import (
	"fmt"

	domain "github.com/dezh-tech/panda/deliveries/http/handlers/domain"
	identifier "github.com/dezh-tech/panda/deliveries/http/handlers/identifier"
	user "github.com/dezh-tech/panda/deliveries/http/handlers/user"
	domainService "github.com/dezh-tech/panda/services/domain"
	identifierService "github.com/dezh-tech/panda/services/identifier"
	userService "github.com/dezh-tech/panda/services/user"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router   *echo.Echo
	config   Config
	handlers Handlers
}

func New(config Config, domainService domainService.Domain, userService userService.User, identifierService identifierService.Identifier) Server {
	return Server{
		Router: echo.New(),
		config: config,

		handlers: Handlers{
			domain:     domain.NewDomainService(domainService),
			user:       user.NewUserService(userService),
			identifier: identifier.NewIdentifierService(identifierService),
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
