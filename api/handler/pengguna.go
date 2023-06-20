package handler

import (
	"errors"

	command "github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/labstack/echo/v4"
)

type PenggunaHandler struct {
	penggunaUsecase *usecase.PenggunaUsecase
}

func NewPenggunaHandler(p *usecase.PenggunaUsecase) *PenggunaHandler {
	return &PenggunaHandler{
		penggunaUsecase: p,
	}
}

func (h *PenggunaHandler) Load(e *echo.Echo) {
	apiGroup := e.Group("/pengguna")

	apiGroup.GET("/cari-pasangan", h.CariPasangan)
	apiGroup.POST("/review", h.BeriReview)
}

func (h *PenggunaHandler) CariPasangan(c echo.Context) error {
	daerah := c.QueryParam("city")
	jenisKelamin := c.QueryParam("gender")

	result, err := h.penggunaUsecase.CariPasanganBerdasarkan(c.Request().Context(), daerah, jenisKelamin)
	if err != nil {
		return c.JSON(500, Response(500, err.Error()))
	}

	return c.JSON(200, ResponseWithData(200, "OK", result))
}

func (h *PenggunaHandler) BeriReview(c echo.Context) error {
	var req command.BeriReviewRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(422, Response(422, err.Error()))
	}

	err := h.penggunaUsecase.ReviewPartner(c.Request().Context(), req)
	if err != nil {
		if !errors.Is(err, ErrInternal) {
			return c.JSON(400, Response(400, err.Error()))
		}

		return c.JSON(500, Response(500, err.Error()))
	}

	return c.JSON(200, Response(200, "Created"))
}
