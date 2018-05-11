package shapeshift

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// GetAPIClient is an interface for a get method api.
type GetAPIClient interface {
	GetRate(ctx context.Context, pair *CurrencyPair) (*RateResponse, error)
}

// CurrencyPair represents a model for currency pair
type CurrencyPair struct {
	Name string `json:"pair,omitempty"`
}

// RateResponse represents a model for rate api response
type RateResponse struct {
	Pair string `json:"pair,omitempty"`
	Rate string `json:"rate"`
}

// GetRate fetches currenct rate for specified currency pair
func (c *client) GetRate(ctx context.Context, pair *CurrencyPair) (*RateResponse, error) {
	url := fmt.Sprintf("rate/%s", pair.Name)
	res, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	rateResp := new(RateResponse)
	if err := json.Unmarshal(res, rateResp); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return rateResp, nil
}
