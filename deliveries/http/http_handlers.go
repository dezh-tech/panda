package http

import (
	domainhandler "github.com/dezh-tech/panda/deliveries/http/handlers/domain_handler"
	_ "github.com/dezh-tech/panda/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Panda Swagger
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

type HttpHandlers struct {
	user domainhandler.Handler
}

func (h *HttpHandlers) Start(r *echo.Echo) {
	h.user.SetRoutes(r)

	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
