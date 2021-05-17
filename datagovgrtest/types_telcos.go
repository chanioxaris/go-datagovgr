package datagovgrtest

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chanioxaris/go-datagovgr/types"
)

func MockIndicatorsAndStatisticsSlice(total int) []*types.IndicatorsAndStatistics {
	mockData := make([]*types.IndicatorsAndStatistics, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockIndicatorsAndStatistics())
	}

	return mockData
}

func MockIndicatorsAndStatistics() *types.IndicatorsAndStatistics {
	return &types.IndicatorsAndStatistics{
		Category:  gofakeit.RandomString([]string{"General", "Landlines", "Mobile"}),
		Indicator: gofakeit.RandomString([]string{"Σταθερές", "Κινητές", "ΕΕ", "Ελλάδα"}),
		Year:      gofakeit.Year(),
		Value:     gofakeit.Float64(),
	}
}
