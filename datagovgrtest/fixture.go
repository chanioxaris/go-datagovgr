package datagovgrtest

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/datagovgr"
	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
)

const (
	baseURL  = "https://test.com"
	apiToken = "test-token"
)

type Fixture struct {
	Client         *datagovgr.Client
	InternalClient *internalclient.Client
	Technology     *api.Technology
	Telcos         *api.Telcos
	MockData       *MockData
}

func NewFixture(t *testing.T) *Fixture {
	t.Helper()

	gofakeit.Seed(0)
	internalClient := internalclient.New(http.DefaultClient, baseURL, apiToken)

	return &Fixture{
		Client:         datagovgr.NewClient(apiToken),
		InternalClient: internalClient,
		Technology:     api.NewTechnology(internalClient),
		Telcos:         api.NewTelcos(internalClient),
		MockData:       NewMockData(t),
	}
}
