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

func TestHealth_COVID19VaccinationStatistics_Success(t *testing.T) {
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
				fixture.URLPaths.COVID19VaccinationStatistics,
				tt.query,
				httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.COVID19VaccinationStatistic),
			)

			got, err := fixture.API.Health.COVID19VaccinationStatistics(ctx, tt.queryParameters...)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}

			if !reflect.DeepEqual(got, fixture.MockData.COVID19VaccinationStatistic) {
				t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.COVID19VaccinationStatistic, got)
			}
		})
	}
}

func TestHealth_COVID19VaccinationStatistics_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.COVID19VaccinationStatistics,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Health.COVID19VaccinationStatistics(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestHealth_InspectionsAndViolations_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.InspectionsAndViolations,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.InspectionsAndViolations),
	)

	got, err := fixture.API.Health.InspectionsAndViolations(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.InspectionsAndViolations) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.InspectionsAndViolations, got)
	}
}

func TestHealth_InspectionsAndViolations_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.InspectionsAndViolations,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Health.InspectionsAndViolations(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestHealth_NumberOfPharmacists_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfPharmacists,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfPharmacists),
	)

	got, err := fixture.API.Health.NumberOfPharmacists(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfPharmacists) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfPharmacists, got)
	}
}

func TestHealth_NumberOfPharmacists_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfPharmacists,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Health.NumberOfPharmacists(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestHealth_NumberOfPharmacies_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfPharmacies,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfPharmacies),
	)

	got, err := fixture.API.Health.NumberOfPharmacies(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfPharmacies) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfPharmacies, got)
	}
}

func TestHealth_NumberOfPharmacies_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfPharmacies,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Health.NumberOfPharmacies(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestHealth_NumberOfDoctors_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfDoctors,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfDoctors),
	)

	got, err := fixture.API.Health.NumberOfDoctors(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfDoctors) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfDoctors, got)
	}
}

func TestHealth_NumberOfDoctors_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfDoctors,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Health.NumberOfDoctors(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestHealth_NumberOfDentists_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfDentists,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfDentists),
	)

	got, err := fixture.API.Health.NumberOfDentists(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfDentists) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfDentists, got)
	}
}

func TestHealth_NumberOfDentists_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfDentists,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Health.NumberOfDentists(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
