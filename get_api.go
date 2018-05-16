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
	GetMarketInfo(ctx context.Context, pair string) (*MarketInfo, error)
	GetTransactionList(ctx context.Context, limit int) (*[]Transaction, error)
	GetTimeRemaining(ctx context.Context, address string) (*TimeRemaining, error)
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

// GetMarketInfo fetches current market info for specified currency pair
func (c *client) GetMarketInfo(ctx context.Context, pair string) (*MarketInfo, error) {
	url := fmt.Sprintf("marketinfo/%s", pair)
	res, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	marketInfoResp := new(MarketInfo)
	if err := json.Unmarshal(res, marketInfoResp); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return marketInfoResp, nil
}

// GetTransactionList fetches recent n executed transactions
// limit is an optional maximum number of transactions to return.
// If limit is not specified this will return 5 transactions.
// Also, limit MUST be a number between 1 and 50 (inclusive).
func (c *client) GetTransactionList(ctx context.Context, limit int) (*[]Transaction, error) {
	url := fmt.Sprintf("recenttx/%d", limit)
	res, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	txResp := new([]Transaction)
	if err := json.Unmarshal(res, txResp); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return txResp, nil
}

// GetTimeRemaining returns how many seconds are left before the transaction expires.
// address is the deposit address to look up.
func (c *client) GetTimeRemaining(ctx context.Context, address string) (*TimeRemaining, error) {
	url := fmt.Sprintf("timeremaining/%s", address)
	res, err := c.do(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	timeRemResp := new(TimeRemaining)
	if err := json.Unmarshal(res, timeRemResp); err != nil {
		return nil, errors.Wrap(err, "")
	}
	return timeRemResp, nil
}
