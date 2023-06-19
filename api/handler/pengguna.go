package handler

import (
	usecase "github.com/PBKKE08/FP-BE/api/usecase/pengguna"
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

func (p *PenggunaHandler) CariPasangan(c echo.Context) error {
	return nil
}

func (p *PenggunaHandler) BeriReview(c echo.Context) error {
	return nil
}