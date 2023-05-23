package auth

import (
	"github.com/labstack/echo/v4"
)

type Service struct {
}

func (s *Service) Login(c echo.Context) error {
	return nil
}

func (s *Service) Logout(c echo.Context) error {
	return nil
}

func (s *Service) Register(c echo.Context) error {
	return nil
}

func (s *Service) VerifyEmail(c echo.Context) error {
	return nil
}
