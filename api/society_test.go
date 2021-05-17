package api_test

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

func TestSociety_UnemploymentClaims_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.UnemploymentClaims,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.UnemploymentClaims),
	)

	got, err := fixture.API.Society.UnemploymentClaims()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.UnemploymentClaims) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.UnemploymentClaims, got)
	}
}

func TestSociety_UnemploymentClaims_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.UnemploymentClaims,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Society.UnemploymentClaims()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestSociety_ElectorsByAge_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.ElectorsByAge,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.ElectorsByAge),
	)

	got, err := fixture.API.Society.ElectorsByAge()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.ElectorsByAge) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.UnemploymentClaims, got)
	}
}

func TestSociety_ElectorsByAge_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.ElectorsByAge,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Society.ElectorsByAge()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}

func TestSociety_ElectorsByRegionAndGender_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.ElectorsByRegionAndGender,
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.ElectorsByRegionAndGender),
	)

	got, err := fixture.API.Society.ElectorsByRegionAndGender()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.ElectorsByRegionAndGender) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.ElectorsByRegionAndGender, got)
	}
}

func TestSociety_ElectorsByRegionAndGender_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fixture.URLPaths.ElectorsByRegionAndGender,
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.API.Society.ElectorsByRegionAndGender()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
