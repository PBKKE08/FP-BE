package handler

import (
	"github.com/PBKKE08/FP-BE/infra/authentication"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

var jwtConfig = echojwt.Config{
	NewClaimsFunc: func(c echo.Context) jwt.Claims {
		return new(authentication.CustomClaims)
	},
	SigningKey: []byte("123123123"),
}

type customClaims = authentication.CustomClaims
