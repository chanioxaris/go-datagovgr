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

func TestBusinessEconomy_NumberOfTravelAgencies_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfTravelAgencies,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfTravelAgencies),
	)

	got, err := fixture.API.BusinessEconomy.NumberOfTravelAgencies(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfTravelAgencies) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfTravelAgencies, got)
	}
}

func TestBusinessEconomy_NumberOfTravelAgencies_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfTravelAgencies,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.BusinessEconomy.NumberOfTravelAgencies(ctx)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestBusinessEconomy_NumberOfRealtors_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfRealtors,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfRealtors),
	)

	got, err := fixture.API.BusinessEconomy.NumberOfRealtors(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfRealtors) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfRealtors, got)
	}
}

func TestBusinessEconomy_NumberOfRealtors_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfRealtors,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.BusinessEconomy.NumberOfRealtors(ctx)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestBusinessEconomy_NumberOfEnergyInspectors_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfEnergyInspectors,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfEnergyInspectors),
	)

	got, err := fixture.API.BusinessEconomy.NumberOfEnergyInspectors(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfEnergyInspectors) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfEnergyInspectors, got)
	}
}

func TestBusinessEconomy_NumberOfEnergyInspectors_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfEnergyInspectors,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.BusinessEconomy.NumberOfEnergyInspectors(ctx)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestBusinessEconomy_NumberOfAuditorsAndFirms_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfAuditorsAndFirms,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfAuditorsAndFirms),
	)

	got, err := fixture.API.BusinessEconomy.NumberOfAuditorsAndFirms(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfAuditorsAndFirms) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfAuditorsAndFirms, got)
	}
}

func TestBusinessEconomy_NumberOfAuditorsAndFirms_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfAuditorsAndFirms,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.BusinessEconomy.NumberOfAuditorsAndFirms(ctx)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestBusinessEconomy_CasinoTickets_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.CasinoTickets,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.CasinoTickets),
	)

	got, err := fixture.API.BusinessEconomy.CasinoTickets(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.CasinoTickets) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.CasinoTickets, got)
	}
}

func TestBusinessEconomy_CasinoTickets_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.CasinoTickets,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.BusinessEconomy.CasinoTickets(ctx)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestBusinessEconomy_NumberOfAccountants_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfAccountants,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.NumberOfAccountants),
	)

	got, err := fixture.API.BusinessEconomy.NumberOfAccountants(ctx)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.NumberOfAccountants) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.NumberOfAccountants, got)
	}
}

func TestBusinessEconomy_NumberOfAccountants_Error(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.NumberOfAccountants,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.BusinessEconomy.NumberOfAccountants(ctx)
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
