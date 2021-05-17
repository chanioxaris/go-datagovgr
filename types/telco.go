package types

// IndicatorsAndStatistics holds indicators and statistics data for Telco companies in Greece.
type IndicatorsAndStatistics struct {
	Category  string  `json:"category" fake:"{randomstring:[General,Landlines,Mobile]}"`
	Indicator string  `json:"indicator" fake:"{randomstring:[Σταθερές,Κινητές,ΕΕ,Ελλάδα]}"`
	Year      int     `json:"year" fake:"{year}"`
	Value     float64 `json:"value" fake:"{float64}"`
}
