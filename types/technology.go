package types

import (
	"time"
)

// InternetTraffic holds the data for the internet traffic in Greece.
type InternetTraffic struct {
	Date            time.Time `json:"date" fake:"{date}"`
	AverageIncoming int       `json:"avg_in" fake:"{number:5000,10000}"`
	AverageOutgoing int       `json:"avg_out" fake:"{number:5000,10000}"`
	MaxIncoming     int       `json:"max_in" fake:"{number:10000,50000}"`
	MaxOutgoing     int       `json:"max_out" fake:"{number:10000,50000}"`
}
