package client

import (
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

// NewRequestGET creates a new http.Request with GET method for provided path and applies required headers.
func (c *Client) NewRequestGET(path string) (*http.Request, error) {
	headers := map[string]string{"Authorization": fmt.Sprintf("Token %s", c.apiToken)}
	return c.newRequest(http.MethodGet, path, nil, headers)
}

func (c *Client) MakeRequest(req *http.Request, payload interface{}) error {
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

func (c *Client) newRequest(method, path string, body io.Reader, headers map[string]string) (*http.Request, error) {
	req, err := http.NewRequest(method, fmt.Sprintf("%s/%s", c.baseURL, path), body)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	return req, nil
}
