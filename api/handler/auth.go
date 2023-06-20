package handler

import (
	"github.com/PBKKE08/FP-BE/api/command/buat_user"
	"github.com/PBKKE08/FP-BE/api/usecase"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase *usecase.AuthUsecase
}

func NewAuthHandler(authUsecase *usecase.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) Load(e *echo.Echo) {
	apiGroup := e.Group("/auth")

	apiGroup.POST("/register", h.Register)
}

func (h *AuthHandler) Register(c echo.Context) error {
	var r buat_user.BuatUserRequest

	if err := c.Bind(&r); err != nil {
		return c.JSON(422, Response(422, err.Error()))
	}

	if err := h.authUsecase.Register(c.Request().Context(), r); err != nil {
		if isInternalErr(err) {
			return c.JSON(500, Response(500, err.Error()))
		}

		return c.JSON(400, Response(400, err.Error()))
	}

	return c.JSON(201, Response(201, "Created"))
}
