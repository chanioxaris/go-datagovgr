package types

import (
	"time"
)

// COVID19VaccinationStatistics holds the data for COVID-19 vaccination statistics.
type COVID19VaccinationStatistics struct {
	Area                 string    `json:"area" fake:"{randomstring:[ΑΡΓΟΛΙΔΑΣ,ΜΥΚΟΝΟΥ,ΚΟΡΙΝΘΙΑΣ]}"`
	AreaID               int       `json:"areaid" fake:"{number:0,1000}"`
	DailyDose1           int       `json:"dailydose1" fake:"{number:0,1000}"`
	DailyDose2           int       `json:"dailydose2" fake:"{number:0,1000}"`
	DayDiff              int       `json:"daydiff" fake:"{number:0,100}"`
	DayTotal             int       `json:"daytotal" fake:"{number:0,2000}"`
	ReferenceDate        time.Time `json:"referencedate" fake:"{date}"`
	TotalDistinctPersons int       `json:"totaldistinctpersons" fake:"{number:0,100000}"`
	TotalDose1           int       `json:"totaldose1" fake:"{number:0,50000}"`
	TotalDose2           int       `json:"totaldose2" fake:"{number:0,50000}"`
	TotalVaccinations    int       `json:"totalvaccinations" fake:"{number:0,100000}"`
}

// InspectionsAndViolations holds the data for inspections and violations.
type InspectionsAndViolations struct {
	Year                   int     `json:"year" fake:"{year}"`
	Inspections            int     `json:"inspections" fake:"{number:0,10000}"`
	Violations             int     `json:"violations" fake:"{number:0,10000}"`
	ViolatingOrganizations float64 `json:"violating_organizations" fake:"{number:0,1000}"`
	Penalties              float64 `json:"penalties" fake:"{float64}"`
}

// NumberOfPharmacists holds the data for the number of pharmacists.
type NumberOfPharmacists struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfPharmacies holds the data for the number of pharmacies.
type NumberOfPharmacies struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfDoctors holds the data for the number of doctors.
type NumberOfDoctors struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}

// NumberOfDentists holds the data for the number of dentists.
type NumberOfDentists struct {
	Year     int    `json:"year" fake:"{year}"`
	Quarter  string `json:"quarter" fake:"{randomstring:[Q1,Q2,Q3,Q4]}"`
	Active   int    `json:"active" fake:"{number:0,10000}"`
	Entrants int    `json:"entrants" fake:"{number:0,100}"`
	Exits    int    `json:"exits" fake:"{number:0,100}"`
}
