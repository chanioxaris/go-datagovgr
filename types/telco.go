package types

// IndicatorsAndStatistics holds indicators and statistics data for Telco companies in Greece.
type IndicatorsAndStatistics struct {
	Category  string  `json:"category"`
	Indicator string  `json:"indicator"`
	Year      int     `json:"year"`
	Value     float64 `json:"value"`
}
