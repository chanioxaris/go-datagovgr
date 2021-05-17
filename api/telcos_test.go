package api_test

import (
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/api"
	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/jarcoal/httpmock"
)

func TestTelcos_IndicatorsAndStatistics_Success(t *testing.T) {
	mockData := datagovgrtest.MockIndicatorsAndStatisticsSlice(3)

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"/eett_telecom_indicators",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, mockData),
	)

	c := client.New(http.DefaultClient, "https://test.com", "test-token")
	telcos := api.NewTelcos(c)

	got, err := telcos.IndicatorsAndStatistics()
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(got, mockData) {
		t.Fatalf("Expected data %+v, but got %+v", mockData, got)
	}
}

func TestTelcos_IndicatorsAndStatistics_Error(t *testing.T) {
	expectedError := "unexpected status code"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"/eett_telecom_indicators",
		httpmock.NewJsonResponderOrPanic(http.StatusInternalServerError, nil),
	)

	c := client.New(http.DefaultClient, "https://test.com", "test-token")
	telcos := api.NewTelcos(c)

	_, err := telcos.IndicatorsAndStatistics()
	if err == nil {
		t.Fatalf("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
