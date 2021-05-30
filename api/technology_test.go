package api_test

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

func TestTechnology_InternetTraffic_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	tests := []struct {
		name            string
		queryParameters []api.QueryParameter
		query           map[string]string
	}{
		{
			name:            "Success without query parameters",
			queryParameters: nil,
			query:           nil,
		},
		{
			name: "Success with date_from query parameter",
			queryParameters: []api.QueryParameter{
				api.WithDateFrom(testTimeFrom),
			},
			query: map[string]string{
				"date_from": "2009-01-03",
			},
		},
		{
			name: "Success with date_to query parameter",
			queryParameters: []api.QueryParameter{
				api.WithDateTo(testTimeTo),
			},
			query: map[string]string{
				"date_to": "2021-05-30",
			},
		},
		{
			name: "Success with date_from and date_to query parameters",
			queryParameters: []api.QueryParameter{
				api.WithDateFrom(testTimeFrom),
				api.WithDateTo(testTimeTo),
			},
			query: map[string]string{
				"date_from": "2009-01-03",
				"date_to":   "2021-05-30",
			},
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponderWithQuery(
				http.MethodGet,
				fixture.URLPaths.InternetTraffic,
				tt.query,
				httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.InternetTraffic),
			)

			got, err := fixture.API.Technology.InternetTraffic(ctx, tt.queryParameters...)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}

			if !reflect.DeepEqual(got, fixture.MockData.InternetTraffic) {
				t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.InternetTraffic, got)
			}
		})
	}
}

func TestTechnology_InternetTraffic_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.InternetTraffic,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Technology.InternetTraffic(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
