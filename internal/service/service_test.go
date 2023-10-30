package service

import (
	"context"
	"testing"
	"time"

	"github.com/romankravchuk/currency-exchanger/internal/data"
	"github.com/romankravchuk/currency-exchanger/internal/service/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func getMocks(t *testing.T) (*CurrencyService, *mocks.CurrencyRatesFetcher, *mocks.CurrencyRatesStorage) {
	t.Helper()

	fetcher := mocks.NewCurrencyRatesFetcher(t)
	storage := mocks.NewCurrencyRatesStorage(t)

	return New(fetcher, storage), fetcher, storage
}

func TestCurrencyService_fetchRates(t *testing.T) {
	t.Parallel()

	t.Run("fetch rates from storage", func(t *testing.T) {
		t.Parallel()

		var (
			ctx                 = context.TODO()
			service, _, storage = getMocks(t)
			query               = &data.RatesQuery{
				From: "USD",
				To:   []string{"RUB"},
			}
			rates = &data.CurrencyRates{
				Date: time.Now(),
				Base: "USD",
				Rates: map[string]float64{
					"RUB": 100.0,
				},
			}
		)

		storage.On("Find", ctx, query.From).
			Once().
			Return(rates, true, nil)

		result, err := service.fetchRates(ctx, query.From, query.To)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotEmpty(t, result)
	})

	t.Run("fecth rates from external service", func(t *testing.T) {
		t.Parallel()

		var (
			now                       = time.Now()
			ctx                       = context.TODO()
			service, fetcher, storage = getMocks(t)
			query                     = &data.RatesQuery{
				From: "USD",
				To:   []string{"RUB"},
			}
			rates = &data.CurrencyRates{
				Date: now,
				Base: "USD",
				Rates: map[string]float64{
					"RUB": 100.0,
				},
			}
		)

		storage.On("Find", ctx, query.From).
			Once().
			Return(nil, false, nil)
		fetcher.On("FetchCurrencyRates", ctx, query.From, query.To).
			Once().
			Return(rates, nil)
		storage.On("Store", ctx, rates, mock.AnythingOfType("time.Duration")).
			Once().
			Return(nil)

		result, err := service.fetchRates(ctx, query.From, query.To)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotEmpty(t, result)
	})
}
