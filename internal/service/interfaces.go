package service

import (
	"context"
	"time"

	"github.com/romankravchuk/currency-exchanger/internal/data"
)

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name CurrencyRatesFetcher --output ./mocks --outpkg mocks
type CurrencyRatesFetcher interface {
	FetchCurrencyRates(context.Context, string, []string) (*data.CurrencyRates, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name CurrencyRatesStorage --output ./mocks --outpkg mocks
type CurrencyRatesStorage interface {
	Store(context.Context, *data.CurrencyRates, time.Duration) error
	Find(context.Context, string) (*data.CurrencyRates, bool, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name CurrencyServicer --output ./mocks --outpkg mocks
type CurrencyServicer interface {
	GetRates(context.Context, *data.RatesQuery) (*data.CurrencyRates, error)
	ConvertCurrency(context.Context, *data.ConvertQuery) (*data.ConvertResult, error)
}
