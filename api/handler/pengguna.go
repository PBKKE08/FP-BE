package handler

import (
	"errors"
	command "github.com/PBKKE08/FP-BE/api/command/beri_review"
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
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

	privateGroup := e.Group("/penggunapriv")
	privateGroup.Use(echojwt.WithConfig(jwtConfig))

	privateGroup.GET("/history-transaksi", h.LihatTransaksi)
	privateGroup.GET("/history-transaksi/:order_id", h.LihatDetailTransaksi)
	privateGroup.GET("/review", h.BeriReview)
}

func (h *PenggunaHandler) CariPasangan(c echo.Context) error {
	daerah := c.QueryParam("city")
	jenisKelamin := c.QueryParam("gender")
	kebutuhan := c.QueryParam("category")

	result, err := h.penggunaUsecase.CariPasanganBerdasarkan(c.Request().Context(), daerah, jenisKelamin, kebutuhan)
	if err != nil {
		return c.JSON(500, Response(500, err.Error()))
	}

	return c.JSON(200, ResponseWithData(200, "OK", result))
}

func (h *PenggunaHandler) BeriReview(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*customClaims)

	var req command.BeriReviewRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(422, Response(422, err.Error()))
	}

	req.PenggunaID = claims.ID

	err := h.penggunaUsecase.ReviewPartner(c.Request().Context(), req)
	if err != nil {
		if !errors.Is(err, ErrInternal) {
			return c.JSON(400, Response(400, err.Error()))
		}

		return c.JSON(500, Response(500, err.Error()))
	}

	return c.JSON(200, Response(200, "Review success"))
}

func (h *PenggunaHandler) LihatTransaksi(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*customClaims)

	results, err := h.penggunaUsecase.LihatRiwayaTransaksi(c.Request().Context(), claims.ID)

	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return c.JSON(200, ResponseWithData(200, "OK", results))
}

func (h *PenggunaHandler) LihatDetailTransaksi(c echo.Context) error {
	id := c.Param("order_id")

	results, err := h.penggunaUsecase.LihatDetailTransaksi(c.Request().Context(), id)

	if err != nil {
		return c.JSON(400, Response(400, err.Error()))
	}

	return c.JSON(200, ResponseWithData(200, "OK", results))
}
