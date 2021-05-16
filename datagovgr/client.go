package datagovgr

import (
	"net/http"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/internal/client"
)

const baseURL = "https://data.gov.gr/api/v1/query"

// Client holds available data.gov.gr API endpoints.
type Client struct {
	Technology *api.Technology
}

// NewClient creates a new Client instance. Requires an API token as input.
func NewClient(apiToken string) *Client {
	internalClient := client.New(http.DefaultClient, baseURL, apiToken)

	return &Client{
		Technology: api.NewTechnology(internalClient),
	}
}
