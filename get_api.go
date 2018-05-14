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
	GetRate(ctx context.Context, pair string) (*Rate, error)
	GetDepositLimit(ctx context.Context, pair string) (*DepositLimit, error)
}

// GetRate fetches currenct rate for specified currency pair
func (c *client) GetRate(ctx context.Context, pair string) (*Rate, error) {
	url := fmt.Sprintf("rate/%s", pair)
	res, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	rateResp := new(Rate)
	if err := json.Unmarshal(res, rateResp); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return rateResp, nil
}

// GetDepositLimit fetches current deposit limit for specified currenct pair
func (c *client) GetDepositLimit(ctx context.Context, pair string) (*DepositLimit, error) {
	url := fmt.Sprintf("limit/%s", pair)
	res, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	depositLimitResp := new(DepositLimit)
	if err := json.Unmarshal(res, depositLimitResp); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return depositLimitResp, nil
}
