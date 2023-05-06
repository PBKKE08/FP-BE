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
	e.GET("/category", controllers.ReadAllCategories)
	e.GET("/city", controllers.ReadAllCities)

	e.POST("/category", controllers.CreateCategory)
	e.POST("/city", controllers.CreateCity)

	e.PUT("/category", controllers.UpdateCategory)
	e.PUT("/city", controllers.UpdateCity)

	e.DELETE("/category", controllers.DeleteCategory)
	e.DELETE("/city", controllers.DeleteCity)

	return e
}