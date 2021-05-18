package datagovgr

import (
	"errors"
	"net/http"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/internal/client"
)

const baseURL = "https://data.gov.gr/api/v1/query"

var (
	// ErrMissingAPIToken indicates an error when a user hasn't provide a valid API token.
	ErrMissingAPIToken = errors.New("you must provide an API token")
)

// Client holds available data.gov.gr API endpoints.
type Client struct {
	BusinessEconomy *api.BusinessEconomy
	Education       *api.Education
	Health          *api.Health
	Society         *api.Society
	Technology      *api.Technology
	Telcos          *api.Telcos
}

// NewClient creates a new Client instance. Requires an API token as input.
func NewClient(apiToken string) (*Client, error) {
	if apiToken == "" {
		return nil, ErrMissingAPIToken
	}

	internalClient := client.New(http.DefaultClient, baseURL, apiToken)

	return &Client{
		BusinessEconomy: api.NewBusinessEconomy(internalClient),
		Education:       api.NewEducation(internalClient),
		Health:          api.NewHealth(internalClient),
		Society:         api.NewSociety(internalClient),
		Technology:      api.NewTechnology(internalClient),
		Telcos:          api.NewTelcos(internalClient),
	}, nil
}
