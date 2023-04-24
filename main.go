package main

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/db"
	"github.com/PBKKE08/FP-BE/echo-rest/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	db.Init()
	e := routes.Init()

	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "Hello from echo!")
	})

	e.Logger.Fatal(e.Start(":1234"))
}
