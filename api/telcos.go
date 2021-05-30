package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// Telcos holds required info to consume Telecommunications related endpoints.
type Telcos struct {
	client *internalclient.Client
}

// NewTelcos creates a new instance.
func NewTelcos(client *internalclient.Client) *Telcos {
	return &Telcos{client: client}
}

// IndicatorsAndStatistics retrieves indicators and statistics data for telecommunications in Greece.
func (t *Telcos) IndicatorsAndStatistics(ctx context.Context) ([]*types.IndicatorsAndStatistics, error) {
	response := make([]*types.IndicatorsAndStatistics, 0)
	if err := t.client.MakeRequestGET(ctx, "eett_telecom_indicators", &response, nil); err != nil {
		return nil, err
	}

	return response, nil
}
