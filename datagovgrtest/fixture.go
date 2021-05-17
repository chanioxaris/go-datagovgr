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
	Education  *api.Education
	Society    *api.Society
	Technology *api.Technology
	Telcos     *api.Telcos
}

type fixtureURLPaths struct {
	AtlasInternshipStatistics     string
	ElectorsByAge                 string
	ElectorsByRegionAndGender     string
	EudoksosRequestsAndDeliveries string
	IndicatorsAndStatistics       string
	InternetTraffic               string
	StudentsBySchool              string
	UnemploymentClaims            string
	UniversityTeachingStaff       string
}

type mockData struct {
	AtlasInternshipStatistics     []*types.AtlasInternshipStatistics
	ElectorsByAge                 []*types.ElectorsByAge
	ElectorsByRegionAndGender     []*types.ElectorsByRegionAndGender
	EudoksosRequestsAndDeliveries []*types.EudoksosRequestsAndDeliveries
	IndicatorsAndStatistics       []*types.IndicatorsAndStatistics
	InternetTraffic               []*types.InternetTraffic
	StudentsBySchool              []*types.StudentsBySchool
	UnemploymentClaims            []*types.UnemploymentClaims
	UniversityTeachingStaff       []*types.UniversityTeachingStaff
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
		Education:  api.NewEducation(c),
		Society:    api.NewSociety(c),
		Technology: api.NewTechnology(c),
		Telcos:     api.NewTelcos(c),
	}
}

func newFixtureURLPaths(t *testing.T) *fixtureURLPaths {
	t.Helper()
	return &fixtureURLPaths{
		AtlasInternshipStatistics:     "/grnet_atlas",
		ElectorsByAge:                 "/minint_election_age",
		ElectorsByRegionAndGender:     "/minint_election_distribution",
		EudoksosRequestsAndDeliveries: "/grnet_eudoxus",
		IndicatorsAndStatistics:       "/eett_telecom_indicators",
		InternetTraffic:               "/internet_traffic",
		StudentsBySchool:              "/minedu_students_school",
		UnemploymentClaims:            "/oaed_unemployment",
		UniversityTeachingStaff:       "/minedu_dep",
	}
}

func newMockData(t *testing.T) *mockData {
	t.Helper()

	var data mockData
	gofakeit.Struct(&data)
	return &data
}
