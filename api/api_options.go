package api

import (
	"time"
)

const (
	dateLayout = "2006-01-02"
)

// QueryParameter describes a functional query parameter for the api functions.
type QueryParameter func(map[string]string)

// WithDateFrom allows to use the optional query parameter 'date_from'.
func WithDateFrom(from time.Time) QueryParameter {
	return func(params map[string]string) {
		params["date_from"] = from.Format(dateLayout)
	}
}

// WithDateTo allows to use the optional query parameter 'date_to'.
func WithDateTo(to time.Time) QueryParameter {
	return func(params map[string]string) {
		params["date_to"] = to.Format(dateLayout)
	}
}
