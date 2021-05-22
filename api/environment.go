package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// Environment holds required info to consume Environment related endpoints.
type Environment struct {
	client *internalclient.Client
}

// NewEnvironment creates a new instance.
func NewEnvironment(client *internalclient.Client) *Environment {
	return &Environment{client: client}
}

// RenewableEnergySources retrieves data for energy produced by renewable energy sources in Greece, in megawatt-hours.
func (e *Environment) RenewableEnergySources(ctx context.Context) ([]*types.RenewableEnergySources, error) {
	response := make([]*types.RenewableEnergySources, 0)
	if err := e.client.MakeRequestGET(ctx, "admie_realtimescadares", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// EnergySystemLoad retrieves data for energy system load for Greece.
func (e *Environment) EnergySystemLoad(ctx context.Context) ([]*types.EnergySystemLoad, error) {
	response := make([]*types.EnergySystemLoad, 0)
	if err := e.client.MakeRequestGET(ctx, "admie_realtimescadasystemload", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// EnergyBalance retrieves data for energy balance analysis for Greece.
func (e *Environment) EnergyBalance(ctx context.Context) ([]*types.EnergyBalance, error) {
	response := make([]*types.EnergyBalance, 0)
	if err := e.client.MakeRequestGET(ctx, "admie_dailyenergybalanceanalysis", &response); err != nil {
		return nil, err
	}

	return response, nil
}

// ElectricityConsumption retrieves data for electricity consumption in Greece.
func (e *Environment) ElectricityConsumption(ctx context.Context) ([]*types.ElectricityConsumption, error) {
	response := make([]*types.ElectricityConsumption, 0)
	if err := e.client.MakeRequestGET(ctx, "electricity_consumption", &response); err != nil {
		return nil, err
	}

	return response, nil
}
