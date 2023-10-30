package handler

import (
	"net/http"

	"github.com/go-chi/render"
)

func response(w http.ResponseWriter, r *http.Request, code int, v render.M) {
	render.Status(r, code)
	render.JSON(w, r, v)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	err := "resource not found"
	response(w, r, http.StatusNotFound, render.M{"error": err})
}

func MethodNotAllowed(w http.ResponseWriter, r *http.Request) {
	err := "method not allowed"
	response(w, r, http.StatusNotFound, render.M{"error": err})
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	err := "bad request"
	response(w, r, http.StatusNotFound, render.M{"error": err})
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	err := "internal server error"
	response(w, r, http.StatusInternalServerError, render.M{"error": err})
}

func UnprocessableEntity(w http.ResponseWriter, r *http.Request, errs map[string]string) {
	response(w, r, http.StatusUnprocessableEntity, render.M{"errors": errs})
}

func OK(w http.ResponseWriter, r *http.Request, v render.M) {
	response(w, r, http.StatusOK, v)
}
