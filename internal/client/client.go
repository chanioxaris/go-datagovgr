package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Client holds the necessary data to perform http requests to data.gov.gr API.
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiToken   string
}

// New creates a new instance of client.
func New(httpClient *http.Client, baseURL, apiToken string) *Client {
	return &Client{
		httpClient: httpClient,
		baseURL:    baseURL,
		apiToken:   apiToken,
	}
}

// MakeRequestGET makes a GET http request to the API and populates the provided payload.
func (c *Client) MakeRequestGET(ctx context.Context, path string, payload interface{}, queryParams map[string]string) error {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil, queryParams)
	if err != nil {
		return err
	}

	return c.makeRequest(req, payload)
}

// newRequest creates a new http request with the provided parameters.
func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader, queryParams map[string]string) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s/%s", c.baseURL, path), body)
	if err != nil {
		return nil, err
	}

	// Add Authorization header to request.
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.apiToken))

	// Add any query parameters to request.
	q := req.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// makeRequest makes an http request and populates the provided payload with the http response body after decoding it.
// Also verifies if the request was successful by checking the response status code.
func (c *Client) makeRequest(req *http.Request, payload interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %v", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(payload); err != nil {
		return err
	}

	return nil
}
