package handler

import (
	"errors"
	"github.com/labstack/echo/v4"
	"strings"
)

var (
	ErrInternal = errors.New("internal server error")
)

func Response(code int, msg string) echo.Map {
	m := echo.Map{
		"code": code,
		"msg":  msg,
	}

	return m
}

func ResponseWithData(code int, msg string, data any) echo.Map {
	m := echo.Map{
		"code": code,
		"msg":  msg,
		"data": data,
	}

	return m
}

func isInternalErr(err error) bool {
	return strings.Contains(err.Error(), "internal server error")
}
