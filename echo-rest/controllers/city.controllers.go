package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/PBKKE08/FP-BE/echo-rest/models"
	"github.com/labstack/echo/v4"
)

func ReadAllCities(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := models.ReadAllCities(db)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}

func CreateCity(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		result, err := models.CreateCity(db, name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}

func UpdateCity(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		name := c.FormValue("name")

		result, err := models.UpdateCity(db, id, name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}

func DeleteCity(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		result, err := models.DeleteCity(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}
