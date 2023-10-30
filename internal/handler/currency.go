package handler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/julienschmidt/httprouter"
	"github.com/romankravchuk/currency-exchanger/internal/data"
	"github.com/romankravchuk/currency-exchanger/internal/service"
	"github.com/romankravchuk/pastebin/pkg/log"
	"github.com/romankravchuk/pastebin/pkg/validator"
)

type CurrencyHandler struct {
	servicer service.CurrencyServicer
	logger   *log.Logger
}

func MountCurrencyRouter(router *httprouter.Router, logger *log.Logger) {
	handler := &CurrencyHandler{
		logger: logger,
	}

	router.POST("/rates", handler.HandleGetRates)
	router.POST("/convert", handler.HandleConvertCurrency)
}

func (h *CurrencyHandler) HandleGetRates(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	input := new(data.RatesQuery)
	if err := input.Bind(r); err != nil {
		h.logger.Error("invalid request body", err, log.FF{
			{Key: "body", Value: input},
		})

		BadRequest(w, r)

		return
	}

	v, err := validator.New()
	if err != nil {
		h.logger.Error("unable to create validator", err, log.FF{
			{Key: "body", Value: input},
		})

		InternalServerError(w, r)

		return
	}

	if !input.Validate(v) {
		h.logger.Info("invalid request body", log.FF{
			{Key: "body", Value: input},
		})

		UnprocessableEntity(w, r, v.Errors())

		return
	}

	rates, err := h.servicer.GetRates(r.Context(), input)
	if err != nil {
		h.logger.Error("unable to get rates", err, log.FF{
			{Key: "body", Value: input},
		})

		InternalServerError(w, r)

		return
	}

	OK(w, r, render.M{"rates": rates})
}

func (h *CurrencyHandler) HandleConvertCurrency(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	input := new(data.ConvertQuery)
	if err := input.Bind(r); err != nil {
		h.logger.Error("invalid request body", err, log.FF{
			{Key: "body", Value: input},
		})

		BadRequest(w, r)

		return
	}

	v, err := validator.New()
	if err != nil {
		h.logger.Error("unable to create validator", err, log.FF{
			{Key: "body", Value: input},
		})

		InternalServerError(w, r)

		return
	}

	if !input.Validate(v) {
		h.logger.Info("invalid request body", log.FF{
			{Key: "body", Value: input},
		})

		UnprocessableEntity(w, r, v.Errors())

		return
	}

	result, err := h.servicer.ConvertCurrency(r.Context(), input)
	if err != nil {
		h.logger.Error("unable to convert currency", err, log.FF{
			{Key: "body", Value: input},
		})

		InternalServerError(w, r)

		return
	}

	OK(w, r, render.M{"result": result})
}
