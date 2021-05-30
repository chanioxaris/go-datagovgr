package api

import (
	"context"

	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
)

// Technology holds required info to consume Technology related endpoints.
type Technology struct {
	client *internalclient.Client
}

// NewTechnology creates a new instance.
func NewTechnology(client *internalclient.Client) *Technology {
	return &Technology{client: client}
}

// InternetTraffic retrieves data for internet traffic in Greece by date.
func (t *Technology) InternetTraffic(ctx context.Context, params ...QueryParameter) ([]*types.InternetTraffic, error) {
	queryParams := make(map[string]string)
	for _, param := range params {
		param(queryParams)
	}

	response := make([]*types.InternetTraffic, 0)
	if err := t.client.MakeRequestGET(ctx, "internet_traffic", &response, queryParams); err != nil {
		return nil, err
	}

	return response, nil
}
