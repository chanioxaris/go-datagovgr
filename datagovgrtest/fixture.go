package datagovgrtest

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/datagovgr"
	internalclient "github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/types"
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
	Society    *api.Society
	Technology *api.Technology
	Telcos     *api.Telcos
}

type fixtureURLPaths struct {
	ElectorsByAge             string
	ElectorsByRegionAndGender string
	IndicatorsAndStatistics   string
	InternetTraffic           string
	UnemploymentClaims        string
}

type mockData struct {
	ElectorsByAge             []*types.ElectorsByAge
	ElectorsByRegionAndGender []*types.ElectorsByRegionAndGender
	IndicatorsAndStatistics   []*types.IndicatorsAndStatistics
	InternetTraffic           []*types.InternetTraffic
	UnemploymentClaims        []*types.UnemploymentClaims
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
		Society:    api.NewSociety(c),
		Technology: api.NewTechnology(c),
		Telcos:     api.NewTelcos(c),
	}
}

func newFixtureURLPaths(t *testing.T) *fixtureURLPaths {
	t.Helper()
	return &fixtureURLPaths{
		ElectorsByAge:             "/minint_election_age",
		ElectorsByRegionAndGender: "/minint_election_distribution",
		IndicatorsAndStatistics:   "/eett_telecom_indicators",
		InternetTraffic:           "/internet_traffic",
		UnemploymentClaims:        "/oaed_unemployment",
	}
}

func newMockData(t *testing.T) *mockData {
	t.Helper()

	var data mockData
	gofakeit.Struct(&data)
	return &data
}
