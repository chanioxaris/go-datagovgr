package testdata

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/chanioxaris/go-datagovgr/types"
)

func MockInternetTrafficSlice(total int) []*types.InternetTraffic {
	mockData := make([]*types.InternetTraffic, 0)

	for i := 0; i < total; i++ {
		mockData = append(mockData, MockInternetTraffic())
	}

	return mockData
}

func MockInternetTraffic() *types.InternetTraffic {
	return &types.InternetTraffic{
		Date:            gofakeit.Date(),
		AverageIncoming: gofakeit.Number(0, 10),
		AverageOutgoing: gofakeit.Number(2, 12),
		MaxIncoming:     gofakeit.Number(100, 1000),
		MaxOutgoing:     gofakeit.Number(200, 2000),
	}
}
