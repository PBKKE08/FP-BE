package controllers

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/models"
	"github.com/labstack/echo/v4"
)

func FetchAllCities(c echo.Context) error{
	result, err := models.FetchAllCities()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostCity(c echo.Context) error{
	name := c.FormValue("name")
	result, err := models.PostCity(name)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}