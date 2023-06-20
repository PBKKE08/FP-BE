package handler

import (
	"errors"
	"strings"
)

var (
	ErrInternal = errors.New("internal server error")
)

func Response(code int, msg string) map[string]any {
	m := make(map[string]any, 3)

	m["code"] = code
	m["msg"] = msg

	return m
}

func ResponseWithData(code int, msg string, data any) map[string]any {
	m := make(map[string]any, 3)

	m["code"] = code
	m["msg"] = msg

	if data != nil {
		m["data"] = data
	}

	return m
}

func isInternalErr(err error) bool {
	return strings.Contains(err.Error(), "internal server error")
}
