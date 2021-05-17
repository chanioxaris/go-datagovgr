package api_test

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

func TestTechnology_InternetTraffic_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"/internet_traffic",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, fixture.MockData.InternetTraffic),
	)

	got, err := fixture.Technology.InternetTraffic()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, fixture.MockData.InternetTraffic) {
		t.Fatalf("Expected data %+v, but got %+v", fixture.MockData.InternetTraffic, got)
	}
}

func TestTechnology_InternetTraffic_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"/internet_traffic",
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	_, err := fixture.Technology.InternetTraffic()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
