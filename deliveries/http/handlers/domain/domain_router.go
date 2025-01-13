package handlers

import (
	"github.com/labstack/echo/v4"
)

func (dh Domain) SetDomainRoutes(e *echo.Echo) {
	userGroup := e.Group("/domains")

	userGroup.POST("", dh.create)
	userGroup.GET("", dh.getAll)
}
