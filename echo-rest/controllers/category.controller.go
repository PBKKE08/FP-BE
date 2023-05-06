package controllers

import (
	"net/http"

	"github.com/PBKKE08/FP-BE/echo-rest/models"
	"github.com/labstack/echo/v4"
)

func FetchAllCategory(c echo.Context) error{
	result, err := models.FetchAllCategories()
	
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string] string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}