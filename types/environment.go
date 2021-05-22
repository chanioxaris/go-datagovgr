package types

import (
	"time"
)

// RenewableEnergySources holds the data for the renewable energy sources.
type RenewableEnergySources struct {
	Date      time.Time `json:"date" fake:"{date}"`
	EnergyMWH int       `json:"energy_mwh" fake:"{number:0,1000}"`
}

// EnergySystemLoad holds the data for the energy system load.
type EnergySystemLoad struct {
	Date      time.Time `json:"date" fake:"{date}"`
	EnergyMWH int       `json:"energy_mwh" fake:"{number:0,10000}"`
}

// EnergyBalance holds the data for the energy balance.
type EnergyBalance struct {
	Date       time.Time `json:"date" fake:"{date}"`
	EnergyMWH  int       `json:"energy_mwh" fake:"{number:0,10000}"`
	Fuel       string    `json:"fuel" fake:"{randomstring:[ΛΙΓΝΙΤΗΣ,ΑΕΡΙΟ,ΑΙΟΛΙΚΑ]}"`
	Percentage float64   `json:"percentage" fake:"{float64}"`
}

// ElectricityConsumption holds the data for the electricity consumption.
type ElectricityConsumption struct {
	Date      time.Time `json:"date" fake:"{date}"`
	EnergyMWH int       `json:"energy_mwh" fake:"{number:0,10000}"`
	Area      string    `json:"area" fake:"{randomstring:[ΚΡΗΤΗ,ΡΟΔΟΣ,ΛΕΣΒΟΣ,ΚΑΡΠΑΘΟΣ]}"`
}
