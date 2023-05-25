package presentation

import (
	"github.com/PBKKE08/FP-BE/V1/business/guest"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo, service *guest.Service) {
	g := e.Group("/guest")
	g.POST("/login", Login(service))
	g.POST("/register", Register(service))
}
