package handler

import (
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	adminUsecase *usecase.AdminUsecase
}

func NewAdminHandler(adminUsecase *usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{adminUsecase: adminUsecase}
}

func (h *AdminHandler) Load(e *echo.Echo) {
	apiGroup := e.Group("/admin")

	apiGroup.GET("/pendaftar", h.GetAllPendaftar)
	apiGroup.PATCH("/terima", h.Terima)
	apiGroup.PATCH("/tolak", h.Tolak)
}

func (h *AdminHandler) GetAllPendaftar(c echo.Context) error {
	results := h.adminUsecase.GetPartnerYangInginMendaftar(c.Request().Context())
	return c.JSON(200, ResponseWithData(200, "OK", results))
}

func (h *AdminHandler) Tolak(c echo.Context) error {
	id := c.QueryParam("id")
	email := c.QueryParam("email")

	err := h.adminUsecase.TolakPartnerPendaftar(c.Request().Context(), id, email)
	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return c.JSON(200, Response(200, "OK"))
}

func (h *AdminHandler) Terima(c echo.Context) error {
	id := c.QueryParam("id")
	email := c.QueryParam("email")

	err := h.adminUsecase.TerimaPartnerPendaftar(c.Request().Context(), id, email)
	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return c.JSON(200, Response(200, "OK"))
}
