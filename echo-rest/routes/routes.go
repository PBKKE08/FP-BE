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

	//select all
	e.GET("/category", controllers.FetchAllCategory)
	e.GET("/city", controllers.FetchAllCities)

	e.POST("/city", controllers.PostCity)
	e.POST("/category", controllers.PostCategory)

	return e
}