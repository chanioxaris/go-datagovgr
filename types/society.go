package types

// UnemploymentClaims holds the data for unemployment claims based on OAED.
type UnemploymentClaims struct {
	Unemployed int    `json:"unemployed" fake:"{number:5000,10000}"`
	Benefits   int    `json:"benefits" fake:"{number:1000,5000}"`
	AsOfDate   string `json:"asofdate" fake:"{year}-{month}-{day}" format:"2006-01-02"`
}

// ElectorsByAge holds the data for electors by age.
type ElectorsByAge struct {
	Year              int    `json:"year" fake:"{year}"`
	AgeGroup          string `json:"age_group" fake:"{randomstring:[18-23,24-29,30-35,36-41]}"`
	Count             int    `json:"count" fake:"{number:10000,100000}"`
	ElectionType      string `json:"election_type" fake:"{randomstring:[βουλευτικές,τοπικής αυτοδιοίκηςη]}"`
	ElectoralDistrict string `json:"electoral_district" fake:"{randomstring:[ΑΙΤΩΛ/ΝΙΑΣ,ΑΙΤΩΛ/ΝΙΑΣ,ΑΡΚΑΔΙΑΣ,ΑΡΤΗΣ]}"`
}

// ElectorsByRegionAndGender holds the data for electors by region and gender.
type ElectorsByRegionAndGender struct {
	Year              int    `json:"year" fake:"{year}"`
	VotersMale        int    `json:"voters_male" fake:"{number:10000,100000}"`
	VotersFemale      int    `json:"voters_female" fake:"{number:10000,100000}"`
	ElectionType      string `json:"election_type" fake:"{randomstring:[βουλευτικές,τοπικής αυτοδιοίκηςη]}"`
	ElectoralDistrict string `json:"electoral_district" fake:"{randomstring:[ΑΙΤΩΛ/ΝΙΑΣ,ΑΡΚΑΔΙΑΣ,ΑΡΤΗΣ]}"`
	Municipality      string `json:"municipality" fake:"{randomstring:[,ΔΕΛΤΑ,ΟΡΕΣΤΙΔΟΣ,ΠΑΤΡΕΩΝ]}"`
}
