package datagovgr_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgr"
	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
)

func TestNewClient_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	c, err := datagovgr.NewClient(fixture.APIToken, datagovgr.WithHTTPClient(http.DefaultClient))
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if c.BusinessEconomy == nil {
		t.Fatal(`Expected "BusinessEconomy" field not to be nil`)
	}

	if c.CrimeJustice == nil {
		t.Fatal(`Expected "CrimeJustice" field not to be nil`)
	}

	if c.Education == nil {
		t.Fatal(`Expected "Education" field not to be nil`)
	}

	if c.Environment == nil {
		t.Fatal(`Expected "Environment" field not to be nil`)
	}

	if c.Health == nil {
		t.Fatal(`Expected "Health" field not to be nil`)
	}

	if c.Society == nil {
		t.Fatal(`Expected "Society" field not to be nil`)
	}

	if c.Technology == nil {
		t.Fatal(`Expected "Technology" field not to be nil`)
	}

	if c.Telcos == nil {
		t.Fatal(`Expected "Telcos" field not to be nil`)
	}
}

func TestNewClient_Error(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	tests := []struct {
		name        string
		apiToken    string
		option      datagovgr.ClientOption
		expectedErr error
	}{
		{
			name:        "Empty API token",
			apiToken:    "",
			expectedErr: datagovgr.ErrMissingAPIToken,
		},
		{
			name:        "Nil http client",
			apiToken:    fixture.APIToken,
			option:      datagovgr.WithHTTPClient(nil),
			expectedErr: datagovgr.ErrNilHTTPClient,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := datagovgr.NewClient(tt.apiToken, tt.option)
			if err == nil {
				t.Fatal("Expected error, but got nil")
			}

			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf(`Expected error "%v", but got "%v"`, tt.expectedErr, err)
			}
		})
	}
}
