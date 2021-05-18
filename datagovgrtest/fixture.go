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
	BusinessEconomy *api.BusinessEconomy
	Education       *api.Education
	Health          *api.Health
	Society         *api.Society
	Technology      *api.Technology
	Telcos          *api.Telcos
}

type fixtureURLPaths struct {
	AtlasInternshipStatistics     string
	CasinoTickets                 string
	COVID19VaccinationStatistics  string
	ElectorsByAge                 string
	ElectorsByRegionAndGender     string
	EudoksosRequestsAndDeliveries string
	IndicatorsAndStatistics       string
	InspectionsAndViolations      string
	InternetTraffic               string
	NumberOfAccountants           string
	NumberOfAuditorsAndFirms      string
	NumberOfDentists              string
	NumberOfDoctors               string
	NumberOfEnergyInspectors      string
	NumberOfPharmacies            string
	NumberOfPharmacists           string
	NumberOfRealtors              string
	NumberOfTravelAgencies        string
	StudentsBySchool              string
	UnemploymentClaims            string
	UniversityTeachingStaff       string
}

type mockData struct {
	AtlasInternshipStatistics     []*types.AtlasInternshipStatistics
	CasinoTickets                 []*types.CasinoTickets
	COVID19VaccinationStatistic   []*types.COVID19VaccinationStatistics
	ElectorsByAge                 []*types.ElectorsByAge
	ElectorsByRegionAndGender     []*types.ElectorsByRegionAndGender
	EudoksosRequestsAndDeliveries []*types.EudoksosRequestsAndDeliveries
	IndicatorsAndStatistics       []*types.IndicatorsAndStatistics
	InspectionsAndViolations      []*types.InspectionsAndViolations
	InternetTraffic               []*types.InternetTraffic
	NumberOfAccountants           []*types.NumberOfAccountants
	NumberOfAuditorsAndFirms      []*types.NumberOfAuditorsAndFirms
	NumberOfDentists              []*types.NumberOfDentists
	NumberOfDoctors               []*types.NumberOfDoctors
	NumberOfEnergyInspectors      []*types.NumberOfEnergyInspectors
	NumberOfPharmacies            []*types.NumberOfPharmacies
	NumberOfPharmacists           []*types.NumberOfPharmacists
	NumberOfRealtors              []*types.NumberOfRealtors
	NumberOfTravelAgencies        []*types.NumberOfTravelAgencies
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
		BusinessEconomy: api.NewBusinessEconomy(c),
		Education:       api.NewEducation(c),
		Health:          api.NewHealth(c),
		Society:         api.NewSociety(c),
		Technology:      api.NewTechnology(c),
		Telcos:          api.NewTelcos(c),
	}
}

func newFixtureURLPaths(t *testing.T) *fixtureURLPaths {
	t.Helper()
	return &fixtureURLPaths{
		AtlasInternshipStatistics:     "/grnet_atlas",
		CasinoTickets:                 "/eeep_casino_tickets",
		COVID19VaccinationStatistics:  "/mdg_emvolio",
		ElectorsByAge:                 "/minint_election_age",
		ElectorsByRegionAndGender:     "/minint_election_distribution",
		EudoksosRequestsAndDeliveries: "/grnet_eudoxus",
		IndicatorsAndStatistics:       "/eett_telecom_indicators",
		InspectionsAndViolations:      "/efet_inspections",
		InternetTraffic:               "/internet_traffic",
		NumberOfAccountants:           "/oee_accountants",
		NumberOfAuditorsAndFirms:      "/elte_auditors",
		NumberOfDentists:              "/minhealth_dentists",
		NumberOfDoctors:               "/minhealth_doctors",
		NumberOfEnergyInspectors:      "/minenv_inspectors",
		NumberOfPharmacies:            "/minhealth_pharmacies",
		NumberOfPharmacists:           "/minhealth_pharmacists",
		NumberOfRealtors:              "/mindev_realtors",
		NumberOfTravelAgencies:        "/mintour_agencies",
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
