package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// Health holds required info to consume Health related endpoints.
type Health struct {
	client *internalclient.Client
}

// NewHealth creates a new instance.
func NewHealth(client *internalclient.Client) *Health {
	return &Health{client: client}
}

// COVID19VaccinationStatistics retrieves general public COVID-19 vaccination statistics by region and day.
func (h *Health) COVID19VaccinationStatistics(ctx context.Context) ([]*types.COVID19VaccinationStatistics, error) {
	response := make([]*types.COVID19VaccinationStatistics, 0)
	if err := h.client.MakeRequestGET(ctx, "mdg_emvolio", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// InspectionsAndViolations retrieves annual statistics of food companies inspections and violations.
func (h *Health) InspectionsAndViolations(ctx context.Context) ([]*types.InspectionsAndViolations, error) {
	response := make([]*types.InspectionsAndViolations, 0)
	if err := h.client.MakeRequestGET(ctx, "efet_inspections", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfPharmacists retrieves the number of active pharmacists including new entrants and retiring professionals.
func (h *Health) NumberOfPharmacists(ctx context.Context) ([]*types.NumberOfPharmacists, error) {
	response := make([]*types.NumberOfPharmacists, 0)
	if err := h.client.MakeRequestGET(ctx, "minhealth_pharmacists", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfPharmacies retrieves the number of active pharmacies including new entrants and retiring professionals.
func (h *Health) NumberOfPharmacies(ctx context.Context) ([]*types.NumberOfPharmacies, error) {
	response := make([]*types.NumberOfPharmacies, 0)
	if err := h.client.MakeRequestGET(ctx, "minhealth_pharmacies", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfDoctors retrieves the number of active doctors including new entrants and retiring professionals.
func (h *Health) NumberOfDoctors(ctx context.Context) ([]*types.NumberOfDoctors, error) {
	response := make([]*types.NumberOfDoctors, 0)
	if err := h.client.MakeRequestGET(ctx, "minhealth_doctors", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// NumberOfDentists retrieves the number of active dentists including new entrants and retiring professionals.
func (h *Health) NumberOfDentists(ctx context.Context) ([]*types.NumberOfDentists, error) {
	response := make([]*types.NumberOfDentists, 0)
	if err := h.client.MakeRequestGET(ctx, "minhealth_dentists", &response); err != nil {
		return nil, err
	}

	return response, nil
}
