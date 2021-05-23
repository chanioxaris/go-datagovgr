package datagovgr

import (
	"errors"
	"net/http"
	"time"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/internal/client"
)

const baseURL = "https://data.gov.gr/api/v1/query"

var (
	// ErrMissingAPIToken indicates an error when a user hasn't provide a valid API token.
	ErrMissingAPIToken = errors.New("you must provide an API token")
)

var (
	defaultHTTClient = &http.Client{Timeout: time.Second * 10}
)

// Client holds available data.gov.gr API endpoints.
type Client struct {
	BusinessEconomy *api.BusinessEconomy
	CrimeJustice    *api.CrimeJustice
	Education       *api.Education
	Environment     *api.Environment
	Health          *api.Health
	Society         *api.Society
	Technology      *api.Technology
	Telcos          *api.Telcos
}

// NewClient creates a new Client instance. Requires an API token as input.
// You can also provide optional options such as a custom http client.
func NewClient(apiToken string, options ...ClientOption) (*Client, error) {
	if apiToken == "" {
		return nil, ErrMissingAPIToken
	}

	clientOptions := &clientOptions{httpClient: defaultHTTClient}

	for _, option := range options {
		if err := option(clientOptions); err != nil {
			return nil, err
		}
	}

	internalClient := client.New(clientOptions.httpClient, baseURL, apiToken)

	return &Client{
		BusinessEconomy: api.NewBusinessEconomy(internalClient),
		CrimeJustice:    api.NewCrimeJustice(internalClient),
		Education:       api.NewEducation(internalClient),
		Environment:     api.NewEnvironment(internalClient),
		Health:          api.NewHealth(internalClient),
		Society:         api.NewSociety(internalClient),
		Technology:      api.NewTechnology(internalClient),
		Telcos:          api.NewTelcos(internalClient),
	}, nil
}
