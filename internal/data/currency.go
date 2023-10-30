package data

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/romankravchuk/pastebin/pkg/validator"
)

type ConvertQuery struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	RatesQuery
}

func (cq *ConvertQuery) Bind(r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(cq)
}

func (cq *ConvertQuery) Validate(v *validator.Validator) bool {
	return v.Valid(cq)
}

type RatesQuery struct {
	From string   `json:"from" validate:"required,oneof=USD EUR RUB JPY"`
	To   []string `json:"to" validate:"required,min=1,dive,oneof=USD EUR RUB JPY,nefield=From"`
}

func (rq *RatesQuery) Bind(r *http.Request) error {
	return json.NewDecoder(r.Body).Decode(rq)
}

func (rq *RatesQuery) Validate(v *validator.Validator) bool {
	return v.Valid(rq)
}

type CurrencyRates struct {
	Date  time.Time          `json:"date"`
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type ConvertResult struct {
	Base    string             `json:"base"`
	Results map[string]float64 `json:"results"`
}
