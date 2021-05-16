package api_test

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/chanioxaris/go-datagovgr/testdata"
	"github.com/jarcoal/httpmock"
)

func TestTechnology_InternetTraffic_Success(t *testing.T) {
	mockData := testdata.MockInternetTrafficSlice(5)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"/internet_traffic",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, mockData),
	)

	c := client.New(http.DefaultClient, "https://test.com", "test-token")
	technology := api.NewTechnology(c)

	got, err := technology.InternetTraffic()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, mockData) {
		t.Fatalf("Expected data %+v, but got %+v", mockData, got)
	}
}

func TestTechnology_InternetTraffic_Error(t *testing.T) {
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"/internet_traffic",
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	c := client.New(http.DefaultClient, "https://test.com", "test-token")
	technology := api.NewTechnology(c)

	_, err := technology.InternetTraffic()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
