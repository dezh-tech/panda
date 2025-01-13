package http

import (
	handlers "github.com/dezh-tech/panda/deliveries/http/handlers/domain"
	_ "github.com/dezh-tech/panda/docs" // revive:disable-line:blank-imports Justification: Required for Swagger documentation
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Panda Swagger
// @version 1.0
// @description Panda is a NOSTR NIP-05 management service developed by Dezh.tech (Dezh technologies).
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.dezh.tech/
// @contact.email hi@dezh.tech

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

type Handlers struct {
	domain handlers.Domain
}

func (h *Handlers) Start(r *echo.Echo) {
	h.domain.SetDomainRoutes(r)

	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
