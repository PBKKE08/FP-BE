package routes

import (
	"database/sql"
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/controllers"
	"github.com/labstack/echo/v4"
)

func GetRoutes(db *sql.DB, e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, this is root echo")
	})

	//select all
	e.GET("/category", controllers.ReadAllCategories(db))
	e.GET("/city", controllers.ReadAllCities(db))

	e.POST("/category", controllers.CreateCategory(db))
	e.POST("/city", controllers.CreateCity(db))

	e.PUT("/category", controllers.UpdateCategory(db))
	e.PUT("/city", controllers.UpdateCity(db))

	e.DELETE("/category", controllers.DeleteCategory(db))
	e.DELETE("/city", controllers.DeleteCity(db))
}
