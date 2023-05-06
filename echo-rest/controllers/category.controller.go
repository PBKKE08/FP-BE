package controllers

import (
	"net/http"
	"strconv"

	"github.com/PBKKE08/FP-BE/echo-rest/models"
	"github.com/labstack/echo/v4"
)

func ReadAllCategories(c echo.Context) error{
	result, err := models.ReadAllCategories()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func CreateCategory(c echo.Context) error{
	name := c.FormValue("name") 
	result, err := models.CreateCategory(name)
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCategory(c echo.Context) error{
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}
	
	name := c.FormValue("name") 
	
	result, err := models.UpdateCategory(id, name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteCategory(c echo.Context) error{
	id, err := strconv.Atoi(c.FormValue("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}
	
	result, err := models.DeleteCategory(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}