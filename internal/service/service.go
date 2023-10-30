package service

import (
	"context"
	"fmt"
	"time"

	"github.com/romankravchuk/currency-exchanger/internal/data"
)

type CurrencyService struct {
	storage CurrencyRatesStorage
	fetcher CurrencyRatesFetcher
}

func New(fetcher CurrencyRatesFetcher, storage CurrencyRatesStorage) *CurrencyService {
	return &CurrencyService{
		fetcher: fetcher,
		storage: storage,
	}
}

func (s *CurrencyService) GetRates(ctx context.Context, query *data.RatesQuery) (*data.CurrencyRates, error) {
	rates, err := s.fetchRates(ctx, query.From, query.To)
	if err != nil {
		return nil, err
	}

	return rates, nil
}

func (s *CurrencyService) ConvertCurrency(ctx context.Context, query *data.ConvertQuery) (*data.ConvertResult, error) {
	rates, err := s.fetchRates(ctx, query.From, query.To)
	if err != nil {
		return nil, err
	}

	result := &data.ConvertResult{
		Base:    query.From,
		Results: make(map[string]float64, len(query.To)),
	}
	for v := range rates.Rates {
		result.Results[v] = rates.Rates[v] * query.Amount
	}

	return result, nil
}

func (s *CurrencyService) fetchRates(ctx context.Context, from string, to []string) (*data.CurrencyRates, error) {
	rates, found, err := s.storage.Find(ctx, from)
	if err != nil {
		return nil, fmt.Errorf("unable to find rates from storage: %w", err)
	}

	if !found {
		rates, err = s.fetcher.FetchCurrencyRates(ctx, from, to)
		if err != nil {
			return nil, fmt.Errorf("unable to fetch rates from external source: %w", err)
		}

		expires := time.Since(time.Date(
			rates.Date.Year(),
			rates.Date.Month(),
			rates.Date.Day()+1,
			0, 0, 0, 0,
			time.UTC,
		))

		err = s.storage.Store(ctx, rates, expires)
		if err != nil {
			return nil, fmt.Errorf("unable to store rates to storage: %w", err)
		}
	}

	return rates, nil
}
