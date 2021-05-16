package client_test

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/chanioxaris/go-datagovgr/internal/client"
	"github.com/jarcoal/httpmock"
)

type testPayload struct {
	Author string `json:"author"`
}

func TestClient_NewRequestGET_Success(t *testing.T) {
	c := client.New(http.DefaultClient, "https://test.com", "test-token")

	got, err := c.NewRequestGET("test-path")
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if got.Method != http.MethodGet {
		t.Fatalf(`Expected method "GET", but got "%v"`, got.Method)
	}

	if got.Header.Get("Authorization") == "" {
		t.Fatal("Expected Authorization header to have been set")
	}

	if headerValue := got.Header.Get("Authorization"); headerValue != "Token test-token" {
		t.Fatalf(`Expected Authorization header value "Token test-token", but got "%v"`, headerValue)
	}

	if urlHost := got.URL.Host; urlHost != "test.com" {
		t.Fatalf(`Expected url host "test.com", but got "%v"`, urlHost)
	}

	if urlPath := got.URL.Path; urlPath != "/test-path" {
		t.Fatalf(`Expected url path "/test-path", but got "%v"`, urlPath)
	}
}

func TestClient_MakeRequest_Success(t *testing.T) {
	expectedPayload := testPayload{Author: "chanioxaris"}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		"https://test.com/test-path",
		httpmock.NewJsonResponderOrPanic(http.StatusOK, expectedPayload),
	)

	c := client.New(http.DefaultClient, "https://test.com", "test-token")

	req, _ := c.NewRequestGET("test-path")

	payload := testPayload{}
	if err := c.MakeRequest(req, &payload); err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(payload, expectedPayload) {
		t.Fatalf("Expected payload %+v, but got %+v", expectedPayload, payload)
	}
}

func TestClient_MakeRequest_Error(t *testing.T) {
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
				fmt.Sprintf("https://test.com/%s", tt.path),
				httpmock.NewStringResponder(tt.statusCode, tt.body),
			)

			c := client.New(http.DefaultClient, "https://test.com", "test-token")
			req, _ := c.NewRequestGET(tt.path)

			payload := testPayload{}
			err := c.MakeRequest(req, &payload)
			if err == nil {
				t.Fatal("Expected error, but got nil")
			}

			if !strings.Contains(err.Error(), tt.expectedError) {
				t.Fatalf(`Expected error to contain "%v", but got "%v"`, tt.expectedError, err)
			}
		})
	}
}
