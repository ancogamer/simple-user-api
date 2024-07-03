package models

import (
	"strconv"
	"strings"
)

type Error struct {
	HTTPCode int    `json:"code"`
	Msg      string `json:"msg"`
}

func NewErr(code int, msg string) (err error) {
	return &Error{
		HTTPCode: code,
		Msg:      msg,
	}
}

func (e *Error) Error() string {
	return strings.Join([]string{strconv.Itoa(e.HTTPCode), ":", e.Msg}, "")
}

func UnWrapperError(err error) (out *Error) {
	out, _ = err.(*Error)

	return
}
