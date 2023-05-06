package controllers

import (
	"net/http"
	"strconv"

	"github.com/PBKKE08/FP-BE/echo-rest/models"
	"github.com/labstack/echo/v4"
)

func ReadAllCities(c echo.Context) error{
	result, err := models.ReadAllCities()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCity(c echo.Context) error{
	name := c.FormValue("name")
	result, err := models.CreateCity(name)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCity(c echo.Context) error{
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	name := c.FormValue("name")
	
	result, err := models.UpdateCity(id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCity(c echo.Context) error{
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	result, err := models.DeleteCity(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}