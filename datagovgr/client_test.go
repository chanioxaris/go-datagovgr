package datagovgr_test

import (
	"errors"
	"testing"

	"github.com/chanioxaris/go-datagovgr/datagovgr"
	"github.com/chanioxaris/go-datagovgr/datagovgrtest"
)

func TestNewClient_Success(t *testing.T) {
	fixture := datagovgrtest.NewFixture(t)

	c, err := datagovgr.NewClient(fixture.APIToken)
	if err != nil {
		t.Fatalf("Unexpected error %v", err)
	}

	if c.Education == nil {
		t.Fatal(`Expected "Education" field not to be nil`)
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
	tests := []struct {
		name        string
		apiToken    string
		expectedErr error
	}{
		{
			name:        "Empty API token",
			apiToken:    "",
			expectedErr: datagovgr.ErrMissingAPIToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := datagovgr.NewClient(tt.apiToken)
			if err == nil {
				t.Fatalf("Expected error, but got nil")
			}

			if !errors.Is(err, tt.expectedErr) {
				t.Fatalf(`Expected error "%v", but got "%v"`, tt.expectedErr, err)
			}
		})
	}
}
