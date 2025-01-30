package handlers

import (
	middleware "github.com/dezh-tech/panda/deliveries/http/middlewares"
	"github.com/labstack/echo/v4"
)

func (dh Identifier) SetIdentifierRoutes(e *echo.Echo) {
	userGroup := e.Group("/identifiers")

	userGroup.POST("", dh.create)
	userGroup.GET("", dh.getAllByPubkey, middleware.Auth("a"))
}
