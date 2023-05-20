package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/PBKKE08/FP-BE/echo-rest/models"
	"github.com/labstack/echo/v4"
)

func ReadAllCategories(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := models.ReadAllCategories(db)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}

func CreateCategory(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.FormValue("name")
		result, err := models.CreateCategory(db, name)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}
}

func UpdateCategory(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		name := c.FormValue("name")

		result, err := models.UpdateCategory(db, id, name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})

		}
		return c.JSON(http.StatusOK, result)
	}
}

func DeleteCategory(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.FormValue("id"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		result, err := models.DeleteCategory(db, id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}

}
