package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// CrimeJustice holds required info to consume Crime and Justice related endpoints.
type CrimeJustice struct {
	client *internalclient.Client
}

// NewCrimeJustice creates a new instance.
func NewCrimeJustice(client *internalclient.Client) *CrimeJustice {
	return &CrimeJustice{client: client}
}

// TrafficAccidents retrieves data for road traffic accidents and casualties by Police jurisdiction.
func (c *CrimeJustice) TrafficAccidents(ctx context.Context) ([]*types.TrafficAccidents, error) {
	response := make([]*types.TrafficAccidents, 0)
	if err := c.client.MakeRequestGET(ctx, "mcp_traffic_accidents", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// RescueOperations retrieves annual statistics on rescue operations of the Hellenic Coast Guard.
func (c *CrimeJustice) RescueOperations(ctx context.Context) ([]*types.RescueOperations, error) {
	response := make([]*types.RescueOperations, 0)
	if err := c.client.MakeRequestGET(ctx, "hcg_incidents", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// TrafficViolations retrieves data for road traffic violations by type.
func (c *CrimeJustice) TrafficViolations(ctx context.Context) ([]*types.TrafficViolations, error) {
	response := make([]*types.TrafficViolations, 0)
	if err := c.client.MakeRequestGET(ctx, "mcp_traffic_violations", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CrimeStatistics retrieves data for crime statistics by type.
func (c *CrimeJustice) CrimeStatistics(ctx context.Context) ([]*types.CrimeStatistics, error) {
	response := make([]*types.CrimeStatistics, 0)
	if err := c.client.MakeRequestGET(ctx, "mcp_crime", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// FinancialCrimes retrieves number of financial crimes investigated by financial police.
func (c *CrimeJustice) FinancialCrimes(ctx context.Context) ([]*types.FinancialCrimes, error) {
	response := make([]*types.FinancialCrimes, 0)
	if err := c.client.MakeRequestGET(ctx, "mcp_financial_crimes", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfLawyers retrieves the number of active lawyer including new entrants and retiring professionals.
func (c *CrimeJustice) NumberOfLawyers(ctx context.Context) ([]*types.NumberOfLawyers, error) {
	response := make([]*types.NumberOfLawyers, 0)
	if err := c.client.MakeRequestGET(ctx, "minjust_lawyers", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfLawFirms retrieves the number of active law firms including new entrants and retiring professionals.
func (c *CrimeJustice) NumberOfLawFirms(ctx context.Context) ([]*types.NumberOfLawFirms, error) {
	response := make([]*types.NumberOfLawFirms, 0)
	if err := c.client.MakeRequestGET(ctx, "minjust_law_firms", &response); err != nil {
		return nil, err
	}

	return response, nil
}
