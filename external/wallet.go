package external

import (
	"bytes"
	"context"
	"e-wallet-transaction/helpers"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type UpdateBalance struct {
	Reference string  `json:"reference"`
	Amount    float64 `json:"amount"`
}

type UpdateBalanceResponse struct {
	Message string `json:"message"`
	Data    struct {
		Balance float64 `json:"balance"`
	} `json:"data"`
}

func (e *External) CreditBalance(ctx context.Context, token string, req UpdateBalance) (*UpdateBalanceResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	url := helpers.GetEnv("WALLET_HOST", "localhost") + helpers.GetEnv("WALLET_ENDPOINT_CREDIT", "/")
	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create wallet http request")
	}

	httpReq.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to create wallet")
	}

	result := &UpdateBalanceResponse{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	defer resp.Body.Close()
	return result, nil
}

func (e *External) DebitBalance(ctx context.Context, token string, req UpdateBalance) (*UpdateBalanceResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal payload")
	}

	url := helpers.GetEnv("WALLET_HOST", "localhost") + helpers.GetEnv("WALLET_ENDPOINT_DEBIT", "/")
	httpReq, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create wallet http request")
	}

	fmt.Println(token)
	httpReq.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect wallet service")
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to debit balance")
	}

	result := &UpdateBalanceResponse{}
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	defer resp.Body.Close()
	return result, nil
}
