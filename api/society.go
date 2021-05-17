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
func (s *Society) UnemploymentClaims(ctx context.Context) ([]*types.UnemploymentClaims, error) {
	req, err := s.client.NewRequestGET(ctx, "oaed_unemployment")
	if err != nil {
		return nil, err
	}

	response := make([]*types.UnemploymentClaims, 0)
	if err := s.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ElectorsByAge retrieves data for electors by age.
func (s *Society) ElectorsByAge(ctx context.Context) ([]*types.ElectorsByAge, error) {
	req, err := s.client.NewRequestGET(ctx, "minint_election_age")
	if err != nil {
		return nil, err
	}

	response := make([]*types.ElectorsByAge, 0)
	if err := s.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ElectorsByRegionAndGender retrieves data for electors by region and gender.
func (s *Society) ElectorsByRegionAndGender(ctx context.Context) ([]*types.ElectorsByRegionAndGender, error) {
	req, err := s.client.NewRequestGET(ctx, "minint_election_distribution")
	if err != nil {
		return nil, err
	}

	response := make([]*types.ElectorsByRegionAndGender, 0)
	if err := s.client.MakeRequest(req, &response); err != nil {
		return nil, err
	}

	return response, nil
}
