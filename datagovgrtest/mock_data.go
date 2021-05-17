package datagovgrtest

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/chanioxaris/go-datagovgr/types"
)

type mockData struct {
	InternetTraffic         []*types.InternetTraffic
	IndicatorsAndStatistics []*types.IndicatorsAndStatistics
}

func newMockData(t *testing.T) *mockData {
	t.Helper()
	return &mockData{
		InternetTraffic:         MockInternetTrafficSlice(t, gofakeit.Number(1, 7)),
		IndicatorsAndStatistics: MockIndicatorsAndStatisticsSlice(t, gofakeit.Number(1, 7)),
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
