package cases

// CovidCases represents the output of the cases api.
type CovidCases struct {
	Country    string               `json:"name"`
	MostRecent CovidCasesMostRecent `json:"mostRecent"`
}

// CountryStruct represents the country in the nested structure in the input json.
type CountryStruct struct {
	Country CovidCases `json:"country"`
}

// DataStruct represents the data in the nested structure in the input json.
type DataStruct struct {
	Data CountryStruct `json:"data"`
}

// CovidCasesMostRecent represents the most recent field in the nested structure in the input json.
type CovidCasesMostRecent struct {
	Date           string  `json:"date"`
	ConfirmedCases int     `json:"confirmed"`
	Recovered      int     `json:"recovered"`
	Deaths         int     `json:"deaths"`
	GrowthRate     float64 `json:"growthRate"`
}

// CovidCasesOutput represents the flat output structure of the endpoint.
type CovidCasesOutput struct {
	Country        string  `json:"country"`
	Date           string  `json:"date"`
	ConfirmedCases int     `json:"confirmed"`
	Recovered      int     `json:"recovered"`
	Deaths         int     `json:"deaths"`
	GrowthRate     float64 `json:"growth_rate"`
}
