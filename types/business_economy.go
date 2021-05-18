package types

// NumberOfTravelAgencies holds the data for the number of travel agencies.
type NumberOfTravelAgencies struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfRealtors holds the data for the number of realtors.
type NumberOfRealtors struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfEnergyInspectors holds the data for the number of energy inspectors.
type NumberOfEnergyInspectors struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfAuditorsAndFirms holds the data for the number of auditors and firms.
type NumberOfAuditorsAndFirms struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
	Firms    int    `json:"firms" fake:"{number:0,100}"`
}

// CasinoTickets holds the data for casino tickets.
type CasinoTickets struct {
	Year    int `json:"year" fake:"{year}"`
	Tickets int `json:"tickets" fake:"{number:0,1000000}"`
}

// NumberOfAccountants holds the data for the number of accountants.
type NumberOfAccountants struct {
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}
