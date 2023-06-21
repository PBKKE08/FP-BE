package handler

import (
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/labstack/echo/v4"
)

type PartnerHandler struct {
	partnerUsecase *usecase.PartnerUsecase
}

func NewPartnerHandler(partnerUsecase *usecase.PartnerUsecase) *PartnerHandler {
	return &PartnerHandler{partnerUsecase: partnerUsecase}
}

func (h *PartnerHandler) Load(e *echo.Echo) {
	apiGroup := e.Group("/partner")

	apiGroup.GET("/:id", h.GetPartnerDetail)
}

func (h *PartnerHandler) GetPartnerDetail(c echo.Context) error {
	id := c.Param("id")

	result, err := h.partnerUsecase.GetPartnerDetail(c.Request().Context(), id)
	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return c.JSON(200, ResponseWithData(200, "OK", result))
}
