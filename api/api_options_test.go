package api_test

import (
	"testing"
	"time"

	"github.com/chanioxaris/go-datagovgr/api"
)

var (
	testTimeFrom = time.Date(2009, 1, 3, 12, 30, 15, 0, time.UTC)
	testTimeTo   = time.Date(2021, 5, 30, 17, 15, 0, 0, time.UTC)
)

func TestWithDateFrom_Success(t *testing.T) {
	queryParams := make(map[string]string)
	api.WithDateFrom(testTimeFrom)(queryParams)

	value, ok := queryParams["date_from"]
	if !ok {
		t.Fatal(`Expected query parameters map to contain the 'date_from' key`)
	}

	if value != "2009-01-03" {
		t.Fatalf(`Expected date to be "2009-01-03", but got "%v"`, value)
	}
}

func TestWithDateTo_Success(t *testing.T) {
	queryParams := make(map[string]string)
	api.WithDateTo(testTimeTo)(queryParams)

	value, ok := queryParams["date_to"]
	if !ok {
		t.Fatal(`Expected query parameters map to contain the 'date_to' key`)
	}

	if value != "2021-05-30" {
		t.Fatalf(`Expected date to be "2021-05-30", but got "%v"`, value)
	}
}
