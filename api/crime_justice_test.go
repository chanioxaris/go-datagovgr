package api_test

import (
	"context"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

func TestCrimeJustice_TrafficAccidents_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.TrafficAccidents,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.TrafficAccidents),
	)

	got, err := fixture.API.CrimeJustice.TrafficAccidents(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.TrafficAccidents) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.TrafficAccidents, got)
	}
}

func TestCrimeJustice_TrafficAccidents_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.TrafficAccidents,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.TrafficAccidents(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestCrimeJustice_RescueOperations_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.RescueOperations,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.RescueOperations),
	)

	got, err := fixture.API.CrimeJustice.RescueOperations(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.RescueOperations) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.RescueOperations, got)
	}
}

func TestCrimeJustice_RescueOperations_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.RescueOperations,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.RescueOperations(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestCrimeJustice_TrafficViolations_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.TrafficViolations,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.TrafficViolations),
	)

	got, err := fixture.API.CrimeJustice.TrafficViolations(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.TrafficViolations) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.TrafficViolations, got)
	}
}

func TestCrimeJustice_TrafficViolations_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.TrafficViolations,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.TrafficViolations(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestCrimeJustice_CrimeStatistics_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.CrimeStatistics,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.CrimeStatistics),
	)

	got, err := fixture.API.CrimeJustice.CrimeStatistics(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.CrimeStatistics) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.CrimeStatistics, got)
	}
}

func TestCrimeJustice_CrimeStatistics_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.CrimeStatistics,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.CrimeStatistics(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestCrimeJustice_FinancialCrimes_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.FinancialCrimes,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.FinancialCrimes),
	)

	got, err := fixture.API.CrimeJustice.FinancialCrimes(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.FinancialCrimes) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.FinancialCrimes, got)
	}
}

func TestCrimeJustice_FinancialCrimes_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.FinancialCrimes,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.FinancialCrimes(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestCrimeJustice_NumberOfLawyers_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfLawyers,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfLawyers),
	)

	got, err := fixture.API.CrimeJustice.NumberOfLawyers(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfLawyers) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfLawyers, got)
	}
}

func TestCrimeJustice_NumberOfLawyers_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfLawyers,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.NumberOfLawyers(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestCrimeJustice_NumberOfLawFirms_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfLawFirms,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfLawFirms),
	)

	got, err := fixture.API.CrimeJustice.NumberOfLawFirms(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfLawFirms) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfLawFirms, got)
	}
}

func TestCrimeJustice_NumberOfLawFirms_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfLawFirms,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.CrimeJustice.NumberOfLawFirms(ctx)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
