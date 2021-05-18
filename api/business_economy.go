package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// BusinessEconomy holds required info to consume Business and Economy related endpoints.
type BusinessEconomy struct {
	client *internalclient.Client
}

// NewBusinessEconomy creates a new instance.
func NewBusinessEconomy(client *internalclient.Client) *BusinessEconomy {
	return &BusinessEconomy{client: client}
}

// NumberOfTravelAgencies retrieves the number of active travel agencies including new entrants and retiring professionals.
func (b *BusinessEconomy) NumberOfTravelAgencies(ctx context.Context) ([]*types.NumberOfTravelAgencies, error) {
	response := make([]*types.NumberOfTravelAgencies, 0)
	if err := b.client.MakeRequestGET(ctx, "mintour_agencies", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfRealtors retrieves the number of active registered realtors including new entrants and retiring professionals.
func (b *BusinessEconomy) NumberOfRealtors(ctx context.Context) ([]*types.NumberOfRealtors, error) {
	response := make([]*types.NumberOfRealtors, 0)
	if err := b.client.MakeRequestGET(ctx, "mindev_realtors", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfEnergyInspectors retrieves the number of active energy inspectors including new entrants and retiring professionals.
func (b *BusinessEconomy) NumberOfEnergyInspectors(ctx context.Context) ([]*types.NumberOfEnergyInspectors, error) {
	response := make([]*types.NumberOfEnergyInspectors, 0)
	if err := b.client.MakeRequestGET(ctx, "minenv_inspectors", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfAuditorsAndFirms retrieves the number of active registered auditors and auditing firms including new entrants and retiring professionals.
func (b *BusinessEconomy) NumberOfAuditorsAndFirms(ctx context.Context) ([]*types.NumberOfAuditorsAndFirms, error) {
	response := make([]*types.NumberOfAuditorsAndFirms, 0)
	if err := b.client.MakeRequestGET(ctx, "elte_auditors", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// CasinoTickets retrieves the total number of tickets issued by casinos in Greece.
func (b *BusinessEconomy) CasinoTickets(ctx context.Context) ([]*types.CasinoTickets, error) {
	response := make([]*types.CasinoTickets, 0)
	if err := b.client.MakeRequestGET(ctx, "eeep_casino_tickets", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfAccountants retrieves the number of active registered accountants including new entrants and retiring professionals.
func (b *BusinessEconomy) NumberOfAccountants(ctx context.Context) ([]*types.NumberOfAccountants, error) {
	response := make([]*types.NumberOfAccountants, 0)
	if err := b.client.MakeRequestGET(ctx, "oee_accountants", &response); err != nil {
		return nil, err
	}

	return response, nil
}
