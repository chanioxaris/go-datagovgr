package types

import (
	"time"
)

// InternetTraffic holds the data for the internet traffic in Greece.
type InternetTraffic struct {
	Date            time.Time `json:"date"`
	AverageIncoming int       `json:"avg_in"`
	AverageOutgoing int       `json:"avg_out"`
	MaxIncoming     int       `json:"max_in"`
	MaxOutgoing     int       `json:"max_out"`
}
