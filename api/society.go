package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// Society holds required info to consume Society related endpoints.
type Society struct {
	client *internalclient.Client
}

// NewSociety creates a new instance.
func NewSociety(client *internalclient.Client) *Society {
	return &Society{client: client}
}

// UnemploymentClaims retrieves data for unemployment claims based on OAED.
func (s *Society) UnemploymentClaims(ctx context.Context, params ...QueryParameter) ([]*types.UnemploymentClaims, error) {
	queryParams := make(map[string]string)
	for _, param := range params {
		param(queryParams)
	}

	response := make([]*types.UnemploymentClaims, 0)
	if err := s.client.MakeRequestGET(ctx, "oaed_unemployment", &response, queryParams); err != nil {
		return nil, err
	}

	return response, nil
}

// ElectorsByAge retrieves data for electors by age.
func (s *Society) ElectorsByAge(ctx context.Context) ([]*types.ElectorsByAge, error) {
	response := make([]*types.ElectorsByAge, 0)
	if err := s.client.MakeRequestGET(ctx, "minint_election_age", &response, nil); err != nil {
		return nil, err
	}

	return response, nil
}

// ElectorsByRegionAndGender retrieves data for electors by region and gender.
func (s *Society) ElectorsByRegionAndGender(ctx context.Context) ([]*types.ElectorsByRegionAndGender, error) {
	response := make([]*types.ElectorsByRegionAndGender, 0)
	if err := s.client.MakeRequestGET(ctx, "minint_election_distribution", &response, nil); err != nil {
		return nil, err
	}

	return response, nil
}
