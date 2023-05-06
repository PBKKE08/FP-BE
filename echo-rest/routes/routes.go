package routes

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/controllers"
	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello, this is root echo")
	})

	e.GET("/category", controllers.FetchAllCategory)

	return e
}