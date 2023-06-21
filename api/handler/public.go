package handler

import (
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/labstack/echo/v4"
)

type PublicHandler struct {
	publicUsecase *usecase.PublicUsecase
}

func NewPublicHandler(publicUsecase *usecase.PublicUsecase) *PublicHandler {
	return &PublicHandler{publicUsecase: publicUsecase}
}

func (h *PublicHandler) Load(e *echo.Echo) {
	apiGroup := e.Group("/public")

	apiGroup.GET("/cc", h.GetAllCityAndCategories)
}

func (h *PublicHandler) GetAllCityAndCategories(c echo.Context) error {
	return c.JSON(200, ResponseWithData(200, "OK", h.publicUsecase.GetAllCityAndCategory(c.Request().Context())))
}
