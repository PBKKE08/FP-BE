package common

import (
	"github.com/rs/zerolog/log"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"code,omitempty"`
	Msg        string `json:"message,omitempty"`
	Token      string `json:"-"`
}

func (res *Response) SetError(statusCode int, err error) {
	res.Msg = err.Error()
	res.StatusCode = statusCode
	log.Error().Msg(err.Error())
}

func (res *Response) Set(statusCode int, msg string) {
	res.Msg = msg
	res.StatusCode = statusCode
}

func (res *Response) SetOK() {
	res.StatusCode = http.StatusOK
	res.Msg = "OK"
}
