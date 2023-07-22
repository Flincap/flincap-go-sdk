package flincap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	defaultBaseURL = "https://flincap.app"
)

// FlincapClient is the client for the Flincap API.
type FlincapClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

// NewFlincapClient creates a new FlincapClient with the given authentication token.
func NewFlincapClient(token string) *FlincapClient {
	return &FlincapClient{
		BaseURL:    defaultBaseURL,
		HTTPClient: &http.Client{},
		Token:      token,
	}
}

// setHeaders sets the required headers for API requests.
func (c *FlincapClient) setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}
}

// handleError handles API request errors and returns an error message.
func (c *FlincapClient) handleError(resp *http.Response) error {
	var errResp struct {
		Message string `json:"message"`
		Code    string `json:"code"`
	}

	err := json.NewDecoder(resp.Body).Decode(&errResp)
	if err != nil {
		return fmt.Errorf("failed to decode error response: %w", err)
	}

	return fmt.Errorf("API error: %s (code: %s)", errResp.Message, errResp.Code)
}

// GetRate retrieves the rate for the specified cryptocurrency and fiat currency.
func (c *FlincapClient) GetRate(selectedCrypt, selectedFiat string) (map[string]interface{}, error) {
	url := c.BaseURL + "/v1/get-rate"
	params := fmt.Sprintf("?selectedCrypt=%s&selectedFiat=%s", selectedCrypt, selectedFiat)
	req, err := http.NewRequest("GET", url+params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.handleError(resp)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response data: %w", err)
	}

	return data, nil
}

// GetExchange retrieves the exchange data.
func (c *FlincapClient) GetExchange() (map[string]interface{}, error) {
	url := c.BaseURL + "/v1/get-exchange"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.handleError(resp)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response data: %w", err)
	}

	return data, nil
}

// CreateTransaction records a transaction.
func (c *FlincapClient) CreateTransaction(transactionData map[string]interface{}) error {
	url := c.BaseURL + "/v1/create-transaction"
	payload, err := json.Marshal(transactionData)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return c.handleError(resp)
	}

	return nil
}

// GetTransaction retrieves a transaction by ID.
func (c *FlincapClient) GetTransaction(transactionID string) (map[string]interface{}, error) {
	url := c.BaseURL + fmt.Sprintf("/v1/get-transactions/%s", transactionID)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.handleError(resp)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response data: %w", err)
	}

	return data, nil
}

// GetTransactionHistory retrieves transaction history based on transaction type and selected fiat.
func (c *FlincapClient) GetTransactionHistory(transactionType, selectedFiat string) (map[string]interface{}, error) {
	url := c.BaseURL + "/api/v1/get-transactions"
	params := fmt.Sprintf("?transactionType=%s&selectedFiat=%s", transactionType, selectedFiat)
	req, err := http.NewRequest("GET", url+params, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	c.setHeaders(req)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, c.handleError(resp)
	}

	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response data: %w", err)
	}

	return data, nil
}
