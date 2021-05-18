package types

// TrafficAccidents holds the data for the traffic accidents.
type TrafficAccidents struct {
	Year             int    `json:"year" fake:"{year}"`
	DeadlyAccidents  int    `json:"deadly_accidents" fake:"{number:0,100}"`
	Deaths           int    `json:"deaths" fake:"{number:0,100}"`
	Jurisdiction     string `json:"jurisdiction" fake:"{randomstring:[Γ.Π.Α.Δ. ΑΤΤΙΚΗΣ,Γ.Π.Α.Δ. ΚΡΗΤΗΣ,Γ.Α.Δ. ΘΕΣΣΑΛΟΝΙΚΗΣ]}"`
	OtherAccidents   int    `json:"other_accidents" fake:"{number:0,1000}"`
	OtherInjured     int    `json:"other_injured" fake:"{number:0,1000}"`
	SeriousAccidents int    `json:"serious_accidents" fake:"{number:0,100}"`
	SeriouslyInjured int    `json:"seriously_injured" fake:"{number:0,100}"`
}

// RescueOperations holds the data for the rescue operations.
type RescueOperations struct {
	Year          int    `json:"year" fake:"{year}"`
	Incident      string `json:"incident" fake:"{randomstring:[ΑΓΝΟΟΥΜΕΝΟ ΣΚΑΦΟΣ,ΕΙΣΡΟΗ ΥΔΑΤΩΝ,ΠΡΟΣΚΡΟΥΣΗ,ΚΑΤΑΓΓΕΛΙΕΣ]}"`
	Domestic      int    `json:"domestic" fake:"{number:0,50}"`
	International int    `json:"international" fake:"{number:0,10}"`
}

// TrafficViolations holds the data for the traffic violations.
type TrafficViolations struct {
	Year      int    `json:"year" fake:"{year}"`
	Violation string `json:"violation" fake:"{randomstring:[ΑΝΑΣΦΑΛΙΣΤΑ ΟΧΗΜΑΤΑ,ΘΟΡΥΒΟΙ ΓΕΝΙΚΑ,ΠΑΡΑΝΟΜΕΣ ΣΤΑΘΜΕΥΣΕΙΣ]}"`
	Count     int    `json:"count" fake:"{number:0,10000}"`
}

// CrimeStatistics holds statistics for crime.
type CrimeStatistics struct {
	Year              int    `json:"year" fake:"{year}"`
	Crime             string `json:"crime" fake:"{randomstring:[ΑΝΘΡΩΠΟΚΤΟΝΙΕΣ,ΑΠΑΤΕΣ,ΕΚΒΙΑΣΕΙΣ]}"`
	Committed         int    `json:"committed" fake:"{number:0,1000}"`
	Attempted         int    `json:"attempted" fake:"{number:0,1000}"`
	Solved            int    `json:"solved" fake:"{number:0,100}"`
	DomesticCriminals int    `json:"domestic_criminals" fake:"{number:0,1000}"`
	ForeignCriminals  int    `json:"foreign_criminals" fake:"{number:0,1000}"`
}

// FinancialCrimes holds the data for the financial crimes.
type FinancialCrimes struct {
	Year  int    `json:"year" fake:"{year}"`
	Crime string `json:"crime" fake:"{randomstring:[Λαθρεμπόριο καυσίμων,Λαθρεμπόριο αλκοόλ,Λαθρεμπόριο τσιγάρων]}"`
	Count int    `json:"count" fake:"{number:0,100}"`
}

// NumberOfLawyers holds the data for the number of lawyers.
type NumberOfLawyers struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfLawFirms holds the data for the number of law firms.
type NumberOfLawFirms struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}
