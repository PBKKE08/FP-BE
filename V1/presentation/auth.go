package presentation

import (
	"github.com/PBKKE08/FP-BE/V1/business/guest"
	"github.com/labstack/echo/v4"
)

type Map map[string]any

var (
	BadInputMap = Map{
		"code":    422,
		"message": "Bad input",
	}
)

func Login(service *guest.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var in guest.LoginInput

		if err := c.Bind(&in); err != nil {
			return c.JSON(422, BadInputMap)
		}

		out := service.Login(in)
		resp := struct {
			Message    string `json:"message"`
			StatusCode int    `json:"code"`
			Token      string `json:"_token"`
		}{
			Message:    out.Msg,
			StatusCode: out.StatusCode,
			Token:      out.Token,
		}

		return c.JSON(out.StatusCode, resp)
	}
}

func Register(service *guest.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var in guest.RegisterInput

		if err := c.Bind(&in); err != nil {
			return c.JSON(422, BadInputMap)
		}

		out := service.Register(in)

		return c.JSON(out.StatusCode, out)
	}
}
