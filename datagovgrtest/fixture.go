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
	testPath = "test-path"
)

// Fixture holds the necessary data to make testing easier.
type Fixture struct {
	BaseURL        string
	APIToken       string
	TestPath       string
	Client         *datagovgr.Client
	InternalClient *internalclient.Client
	API            *fixtureAPI
	URLPaths       *fixtureURLPaths
	MockData       *mockData
}

type fixtureAPI struct {
	Education  *api.Education
	Health     *api.Health
	Society    *api.Society
	Technology *api.Technology
	Telcos     *api.Telcos
}

type fixtureURLPaths struct {
	AtlasInternshipStatistics     string
	COVID19VaccinationStatistics  string
	ElectorsByAge                 string
	ElectorsByRegionAndGender     string
	EudoksosRequestsAndDeliveries string
	IndicatorsAndStatistics       string
	InspectionsAndViolations      string
	InternetTraffic               string
	NumberOfDentists              string
	NumberOfDoctors               string
	NumberOfPharmacies            string
	NumberOfPharmacists           string
	StudentsBySchool              string
	UnemploymentClaims            string
	UniversityTeachingStaff       string
}

type mockData struct {
	AtlasInternshipStatistics     []*types.AtlasInternshipStatistics
	COVID19VaccinationStatistic   []*types.COVID19VaccinationStatistics
	ElectorsByAge                 []*types.ElectorsByAge
	ElectorsByRegionAndGender     []*types.ElectorsByRegionAndGender
	EudoksosRequestsAndDeliveries []*types.EudoksosRequestsAndDeliveries
	IndicatorsAndStatistics       []*types.IndicatorsAndStatistics
	InspectionsAndViolations      []*types.InspectionsAndViolations
	InternetTraffic               []*types.InternetTraffic
	NumberOfDentists              []*types.NumberOfDentists
	NumberOfDoctors               []*types.NumberOfDoctors
	NumberOfPharmacies            []*types.NumberOfPharmacies
	NumberOfPharmacists           []*types.NumberOfPharmacists
	StudentsBySchool              []*types.StudentsBySchool
	UnemploymentClaims            []*types.UnemploymentClaims
	UniversityTeachingStaff       []*types.UniversityTeachingStaff
}

// NewFixture creates a new Fixture instance.
func NewFixture(t *testing.T) *Fixture {
	t.Helper()

	gofakeit.Seed(0)
	client, _ := datagovgr.NewClient(apiToken)
	internalClient := internalclient.New(http.DefaultClient, baseURL, apiToken)

	return &Fixture{
		BaseURL:        baseURL,
		APIToken:       apiToken,
		TestPath:       testPath,
		Client:         client,
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
		Health:     api.NewHealth(c),
		Society:    api.NewSociety(c),
		Technology: api.NewTechnology(c),
		Telcos:     api.NewTelcos(c),
	}
}

func newFixtureURLPaths(t *testing.T) *fixtureURLPaths {
	t.Helper()
	return &fixtureURLPaths{
		AtlasInternshipStatistics:     "/grnet_atlas",
		COVID19VaccinationStatistics:  "/mdg_emvolio",
		ElectorsByAge:                 "/minint_election_age",
		ElectorsByRegionAndGender:     "/minint_election_distribution",
		EudoksosRequestsAndDeliveries: "/grnet_eudoxus",
		IndicatorsAndStatistics:       "/eett_telecom_indicators",
		InspectionsAndViolations:      "/efet_inspections",
		InternetTraffic:               "/internet_traffic",
		NumberOfDentists:              "/minhealth_dentists",
		NumberOfDoctors:               "/minhealth_doctors",
		NumberOfPharmacies:            "/minhealth_pharmacies",
		NumberOfPharmacists:           "/minhealth_pharmacists",
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
