package datagovgr

import (
	"net/http"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/internal/client"
)

const baseURL = "https://data.gov.gr/api/v1/query"

// Client holds available data.gov.gr API endpoints.
type Client struct {
	Education  *api.Education
	Society    *api.Society
	Technology *api.Technology
	Telcos     *api.Telcos
}

// NewClient creates a new Client instance. Requires an API token as input.
func NewClient(apiToken string) *Client {
	internalClient := client.New(http.DefaultClient, baseURL, apiToken)

	return &Client{
		Education:  api.NewEducation(internalClient),
		Society:    api.NewSociety(internalClient),
		Technology: api.NewTechnology(internalClient),
		Telcos:     api.NewTelcos(internalClient),
	}
}
