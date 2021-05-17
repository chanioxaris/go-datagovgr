package datagovgrtest

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/chanioxaris/go-datagovgr/types"
)

type mockData struct {
	ElectorsByAge             []*types.ElectorsByAge
	ElectorsByRegionAndGender []*types.ElectorsByRegionAndGender
	IndicatorsAndStatistics   []*types.IndicatorsAndStatistics
	InternetTraffic           []*types.InternetTraffic
	UnemploymentClaims        []*types.UnemploymentClaims
}

func newMockData(t *testing.T) *mockData {
	t.Helper()
	return &mockData{
		ElectorsByAge:             MockElectorsByAgeSlice(t, gofakeit.Number(1, 7)),
		ElectorsByRegionAndGender: MockElectorsByRegionAndGenderSlice(t, gofakeit.Number(1, 7)),
		IndicatorsAndStatistics:   MockIndicatorsAndStatisticsSlice(t, gofakeit.Number(1, 7)),
		InternetTraffic:           MockInternetTrafficSlice(t, gofakeit.Number(1, 7)),
		UnemploymentClaims:        MockUnemploymentClaimsSlice(t, gofakeit.Number(1, 7)),
	}
}

func MockElectorsByRegionAndGenderSlice(t *testing.T, total int) []*types.ElectorsByRegionAndGender {
	t.Helper()
	mockData := make([]*types.ElectorsByRegionAndGender, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockElectorsByRegionAndGender(t))
	}

	return mockData
}

func MockElectorsByRegionAndGender(t *testing.T) *types.ElectorsByRegionAndGender {
	t.Helper()
	return &types.ElectorsByRegionAndGender{
		Year:              gofakeit.Year(),
		VotersMale:        gofakeit.Number(10000, 100000),
		VotersFemale:      gofakeit.Number(10000, 100000),
		ElectionType:      gofakeit.RandomString([]string{"βουλευτικές", "τοπικής αυτοδιοίκηςη"}),
		ElectoralDistrict: gofakeit.RandomString([]string{"ΑΙΤΩΛ/ΝΙΑΣ", "ΑΙΤΩΛ/ΝΙΑΣ", "ΑΡΚΑΔΙΑΣ", "ΑΡΤΗΣ"}),
		Municipality:      gofakeit.RandomString([]string{"", "ΔΕΛΤΑ", "ΟΡΕΣΤΙΔΟΣ", "ΠΑΤΡΕΩΝ"}),
	}
}

func MockElectorsByAgeSlice(t *testing.T, total int) []*types.ElectorsByAge {
	t.Helper()
	mockData := make([]*types.ElectorsByAge, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockElectorsByAge(t))
	}

	return mockData
}

func MockElectorsByAge(t *testing.T) *types.ElectorsByAge {
	t.Helper()
	return &types.ElectorsByAge{
		Year:              gofakeit.Year(),
		AgeGroup:          gofakeit.RandomString([]string{"18-23", "24-29", "30-35", "36-41"}),
		Count:             gofakeit.Number(10000, 100000),
		ElectionType:      gofakeit.RandomString([]string{"βουλευτικές", "τοπικής αυτοδιοίκηςη"}),
		ElectoralDistrict: gofakeit.RandomString([]string{"ΑΙΤΩΛ/ΝΙΑΣ", "ΑΙΤΩΛ/ΝΙΑΣ", "ΑΡΚΑΔΙΑΣ", "ΑΡΤΗΣ"}),
	}
}

func MockIndicatorsAndStatisticsSlice(t *testing.T, total int) []*types.IndicatorsAndStatistics {
	t.Helper()
	mockData := make([]*types.IndicatorsAndStatistics, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockIndicatorsAndStatistics(t))
	}

	return mockData
}

func MockIndicatorsAndStatistics(t *testing.T) *types.IndicatorsAndStatistics {
	t.Helper()
	return &types.IndicatorsAndStatistics{
		Category:  gofakeit.RandomString([]string{"General", "Landlines", "Mobile"}),
		Indicator: gofakeit.RandomString([]string{"Σταθερές", "Κινητές", "ΕΕ", "Ελλάδα"}),
		Year:      gofakeit.Year(),
		Value:     gofakeit.Float64(),
	}
}

func MockInternetTrafficSlice(t *testing.T, total int) []*types.InternetTraffic {
	t.Helper()
	mockData := make([]*types.InternetTraffic, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockInternetTraffic(t))
	}

	return mockData
}

func MockInternetTraffic(t *testing.T) *types.InternetTraffic {
	t.Helper()
	return &types.InternetTraffic{
		Date:            gofakeit.Date(),
		AverageIncoming: gofakeit.Number(0, 10),
		AverageOutgoing: gofakeit.Number(2, 12),
		MaxIncoming:     gofakeit.Number(100, 1000),
		MaxOutgoing:     gofakeit.Number(200, 2000),
	}
}

func MockUnemploymentClaimsSlice(t *testing.T, total int) []*types.UnemploymentClaims {
	t.Helper()
	mockData := make([]*types.UnemploymentClaims, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockUnemploymentClaims(t))
	}

	return mockData
}

func MockUnemploymentClaims(t *testing.T) *types.UnemploymentClaims {
	t.Helper()
	return &types.UnemploymentClaims{
		Unemployed: gofakeit.Number(1000, 10000),
		Benefits:   gofakeit.Number(100, 2000),
		AsOfDate:   gofakeit.Date().String(),
	}
}
