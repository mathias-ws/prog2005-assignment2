package cases

// CovidCases represents the output of the cases api.
type CovidCases struct {
	Country    string               `json:"name"`
	MostRecent CovidCasesMostRecent `json:"mostRecent"`
}

type CovidCases2 struct {
	Country CovidCases `json:"country"`
}

type Data struct {
	Data CovidCases2 `json:"data"`
}

type CovidCasesMostRecent struct {
	Date           string  `json:"date"`
	ConfirmedCases int     `json:"confirmed"`
	Recovered      int     `json:"recovered"`
	Deaths         int     `json:"deaths"`
	GrowthRate     float64 `json:"growthRate"`
}

type CovidCasesOutput struct {
	Country        string  `json:"country"`
	Date           string  `json:"date"`
	ConfirmedCases int     `json:"confirmed"`
	Recovered      int     `json:"recovered"`
	Deaths         int     `json:"deaths"`
	GrowthRate     float64 `json:"growth_rate"`
}
