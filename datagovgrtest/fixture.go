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
	API            *fixtureAPI
	URLPaths       *fixtureURLPaths
	MockData       *mockData
}

type fixtureAPI struct {
	Technology *api.Technology
	Telcos     *api.Telcos
}

type fixtureURLPaths struct {
	InternetTraffic         string
	IndicatorsAndStatistics string
}

func NewFixture(t *testing.T) *Fixture {
	t.Helper()

	gofakeit.Seed(0)
	internalClient := internalclient.New(http.DefaultClient, baseURL, apiToken)

	return &Fixture{
		Client:         datagovgr.NewClient(apiToken),
		InternalClient: internalClient,
		API:            newFixtureAPI(t, internalClient),
		URLPaths:       newFixtureURLPaths(t),
		MockData:       newMockData(t),
	}
}

func newFixtureAPI(t *testing.T, c *internalclient.Client) *fixtureAPI {
	t.Helper()
	return &fixtureAPI{
		Technology: api.NewTechnology(c),
		Telcos:     api.NewTelcos(c),
	}
}

func newFixtureURLPaths(t *testing.T) *fixtureURLPaths {
	t.Helper()
	return &fixtureURLPaths{
		InternetTraffic:         "/internet_traffic",
		IndicatorsAndStatistics: "/eett_telecom_indicators",
	}
}