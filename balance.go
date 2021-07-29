package gocs

import (
	"context"
	"net/http"
	"time"
)

// this interfaces with the endpoints for getting the account balance
type BalanceService interface {
	Get(context.Context) (*Balance, *Response, error)
}

// handles communication with the balance-related methods of the API
type BalanceServiceOp struct {
	client *Client
}

var _ BalanceServcie = &BalanceServiceOp{}

// represents a Cloud Sigma account balance
// TODO: this struct needs to know the actual json keys for the returned data...`account_balance` is just a placeholder
type Balance struct {
	TotalBalance string `json:"account_balance"`
	Currency     string `json:"Currency"`
}

func (r Balance) String() string {
	return Stringify(r)
}

// get the balance info
func (s *BalanceServiceOp) Get(ctx context.Context) (*Balance, *Response, error) {
	path := "/api/2.0/balance"

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(Balance)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}
