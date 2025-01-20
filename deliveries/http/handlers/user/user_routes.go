package handlers

import (
	"github.com/labstack/echo/v4"
)

func (dh User) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	userGroup.POST("", dh.create)
}
