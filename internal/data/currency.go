package data

import "time"

type ConvertQuery struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	RatesQuery
}

type RatesQuery struct {
	From string   `json:"from" validate:"required,oneof=USD EUR RUB JPY"`
	To   []string `json:"to" validate:"required,min=1,dive,oneof=USD EUR RUB JPY,nefield=From"`
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
