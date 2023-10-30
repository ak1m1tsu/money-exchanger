package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/romankravchuk/currency-exchanger/internal/data"
)

const apiEndpoint = "https://api.currencyfreaks.com/v2.0/rates/latest"

type CurrencyFreaksClient struct {
	endpoint string
	apiKey   string
}

func New(apiKey string) *CurrencyFreaksClient {
	return &CurrencyFreaksClient{
		apiKey:   apiKey,
		endpoint: apiEndpoint,
	}
}

func (c *CurrencyFreaksClient) FetchCurrencyRates(ctx context.Context, from string, to []string) (*data.CurrencyRates, error) {
	endpoint := fmt.Sprintf("%s?apikey=%s&base=%s&symbols=%s", c.endpoint, c.apiKey, from, strings.Join(to, ","))

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result data.CurrencyRates
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
