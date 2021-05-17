package types

// UnemploymentClaims holds the data for unemployment claims based on OAED.
type UnemploymentClaims struct {
	Unemployed int    `json:"unemployed"`
	Benefits   int    `json:"benefits"`
	AsOfDate   string `json:"asofdate"`
}

// ElectorsByAge holds the data for electors by age.
type ElectorsByAge struct {
	Year              int    `json:"year"`
	AgeGroup          string `json:"age_group"`
	Count             int    `json:"count"`
	ElectionType      string `json:"election_type"`
	ElectoralDistrict string `json:"electoral_district"`
}

// ElectorsByRegionAndGender holds the data for electors by region and gender.
type ElectorsByRegionAndGender struct {
	Year              int    `json:"year"`
	VotersMale        int    `json:"voters_male"`
	VotersFemale      int    `json:"voters_female"`
	ElectionType      string `json:"election_type"`
	ElectoralDistrict string `json:"electoral_district"`
	Municipality      string `json:"municipality"`
}
