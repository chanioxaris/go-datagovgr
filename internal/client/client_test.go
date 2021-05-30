package client_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
	"github.com/jarcoal/httpmock"
)

type testPayload struct {
	Author string `json:"author"`
}

func TestClient_MakeRequestGET_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedPayload := testPayload{Author: "chanioxaris"}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s/%s", fixture.BaseURL, fixture.TestPath),
		httpmock.NewJsonResponderOrPanic(http.StatusOK, expectedPayload),
	)

	payload := testPayload{}
	err := fixture.InternalClient.MakeRequestGET(ctx, fixture.TestPath, &payload, nil)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if !reflect.DeepEqual(payload, expectedPayload) {
		t.Fatalf("Expected payload %+v, but got %+v", expectedPayload, payload)
	}
}

func TestClient_MakeRequestGET_QueryParameters_Success(t *testing.T) {
	ctx := context.Background()
	fixture := datagovgrtest.NewFixture(t)
	expectedPayload := testPayload{Author: "chanioxaris"}

	tests := []struct {
		name  string
		query map[string]string
	}{
		{
			name: "HTTP request with date_from query parameter",
			query: map[string]string{
				"date_from": "2009-01-03",
			},
		},
		{
			name: "HTTP request with date_to query parameter",
			query: map[string]string{
				"date_to": "2009-01-03",
			},
		},
		{
			name: "HTTP request with date_from and date_to query parameters",
			query: map[string]string{
				"date_from": "2009-01-03",
				"date_to":   "2009-01-04",
			},
		},
	}

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpmock.RegisterResponderWithQuery(
				http.MethodGet,
				fmt.Sprintf("%s/%s", fixture.BaseURL, fixture.TestPath),
				tt.query,
				httpmock.NewJsonResponderOrPanic(http.StatusOK, expectedPayload),
			)

			payload := testPayload{}
			err := fixture.InternalClient.MakeRequestGET(ctx, fixture.TestPath, &payload, tt.query)
			if err != nil {
				t.Fatalf("Unexpected error %v", err)
			}

			if !reflect.DeepEqual(payload, expectedPayload) {
				t.Fatalf("Expected payload %+v, but got %+v", expectedPayload, payload)
			}
		})
	}
}

func TestClient_MakeRequestGET_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	tests := []struct {
		name          string
		path          string
		statusCode    int
		body          string
		ctx           context.Context
		expectedError string
	}{
		{
			name:          "Nil context",
			path:          "nil-context",
			statusCode:    0,
			body:          "",
			ctx:           nil,
			expectedError: "nil Context",
		},
		{
			name:          "Unexpected status code",
			path:          "unexpected-status-code",
			statusCode:    http.StatusInternalServerError,
			body:          "",
			ctx:           context.Background(),
			expectedError: fmt.Sprintf("unexpected status code: %v", http.StatusInternalServerError),
		},
		{
			name:          "Invalid response body",
			path:          "invalid-response-body",
			statusCode:    http.StatusOK,
			body:          `{"author": 13}`,
			ctx:           context.Background(),
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

			payload := testPayload{}
			err := fixture.InternalClient.MakeRequestGET(tt.ctx, tt.path, &payload, nil)
			if err == nil {
				t.Fatal("Expected error, but got nil")
			}

			if !strings.Contains(err.Error(), tt.expectedError) {
				t.Fatalf(`Expected error to contain "%v", but got "%v"`, tt.expectedError, err)
			}
		})
	}
}

func TestClient_MakeRequestGET_Error_Timeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*500)
	defer cancel()

	fixture := datagovgrtest.NewFixture(t)
	expectedError := "context deadline exceeded"

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder(
		http.MethodGet,
		fmt.Sprintf("%s/%s", fixture.BaseURL, fixture.TestPath),
		func(req *http.Request) (*http.Response, error) {
			time.Sleep(time.Second * 1)
			return httpmock.NewJsonResponse(http.StatusOK, testPayload{Author: "chanioxaris"})
		},
	)

	payload := testPayload{}
	err := fixture.InternalClient.MakeRequestGET(ctx, fixture.TestPath, &payload, nil)
	if err == nil {
		t.Fatal("Expected error, but got nil")
	}

	if !strings.Contains(err.Error(), expectedError) {
		t.Fatalf(`Expected error to contain "%v", but got "%v"`, expectedError, err)
	}
}
