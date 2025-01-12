package domainhandler

import (
	"github.com/labstack/echo/v4"
)

func (h Handler) SetRoutes(e *echo.Echo) {
	userGroup := e.Group("/domains")

	userGroup.POST("", h.domainCreate)
	userGroup.GET("", h.domainGetAll)
}
