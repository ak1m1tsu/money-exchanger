package data

import (
	"net/http"

	"github.com/romankravchuk/pastebin/pkg/validator"
)

type BindValidater interface {
	Binder
	Validater
}

type Binder interface {
	Bind(*http.Request) error
}

type Validater interface {
	Validate(v *validator.Validator) bool
}
