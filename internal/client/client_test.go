package client_test

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

type testPayload struct {
	Author string `json:"author"`
}

func TestClient_NewRequestGET_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)
	expectedHeaderValue := fmt.Sprintf("Token %s", fixture.APIToken)
	expectedURLHost := strings.TrimPrefix(fixture.BaseURL, "https://")

	got, err := fixture.InternalClient.NewRequestGET("test-path")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if got.Method != http.MethodGet {
		t.Fatalf(`Expected method "GET", but got "%v"`, got.Method)
	}

	if got.Header.Get("Authorization") == "" {
		t.Fatal("Expected Authorization header to have been set")
	}

	if headerValue := got.Header.Get("Authorization"); headerValue != expectedHeaderValue {
		t.Fatalf(`Expected Authorization header value "%s", but got "%v"`, expectedHeaderValue, headerValue)
	}

	if urlHost := got.URL.Host; urlHost != expectedURLHost {
		t.Fatalf(`Expected url host "%s", but got "%v"`, expectedURLHost, urlHost)
	}

	if urlPath := got.URL.Path; urlPath != "/test-path" {
		t.Fatalf(`Expected url path "/test-path", but got "%v"`, urlPath)
	}
}

func TestClient_MakeRequest_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)
	expectedPayload := testPayload{Author: "chanioxaris"}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s/test-path", fixture.BaseURL),
		httpmock.NewJsonResponderOrPanic(http.StatusOK, expectedPayload),
	)

	req, _ := fixture.InternalClient.NewRequestGET("test-path")

	payload := testPayload{}
	if err := fixture.InternalClient.MakeRequest(req, &payload); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(payload, expectedPayload) {
		t.Fatalf("Expected payload %+v, but got %+v", expectedPayload, payload)
	}
}

func TestClient_MakeRequest_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	tests := []struct {
		name          string
		path          string
		statusCode    int
		body          string
		expectedError string
	}{
		{
			name:          "Unexpected status code",
			path:          "unexpected-status-code",
			statusCode:    http.StatusInternalServerError,
			body:          "",
			expectedError: fmt.Sprintf("unexpected status code: %v", http.StatusInternalServerError),
		},
		{
			name:          "Invalid response body",
			path:          "invalid-response-body",
			statusCode:    http.StatusOK,
			body:          `{"author": 13}`,
			expectedError: "json: cannot unmarshal",
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponder(
				http.MethodGet,
				fmt.Sprintf("%s/%s", fixture.BaseURL, tt.path),
				httpmock.NewStringResponder(tt.statusCode, tt.body),
			)

			req, _ := fixture.InternalClient.NewRequestGET(tt.path)

			payload := testPayload{}
			err := fixture.InternalClient.MakeRequest(req, &payload)
			if err == nil {
				t.Fatal("Expected error, but got nil")
			}

			if !strings.Contains(err.Error(), tt.expectedError) {
				t.Fatalf(`Expected error to contain "%v", but got "%v"`, tt.expectedError, err)
			}
		})
	}
}
