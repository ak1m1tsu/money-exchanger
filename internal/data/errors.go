package data

import "errors"

var (
	ErrBadRequest          = errors.New("bad request")
	ErrInternalServer      = errors.New("internal server error")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
)
