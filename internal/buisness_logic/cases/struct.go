package cases

import "time"

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
	Country        string    `json:"country" firestore:"country"`
	Date           string    `json:"date" firestore:"date"`
	ConfirmedCases int       `json:"confirmed" firestore:"confirmed"`
	Recovered      int       `json:"recovered" firestore:"recovered"`
	Deaths         int       `json:"deaths" firestore:"deaths"`
	GrowthRate     float64   `json:"growth_rate" firestore:"growth_rate"`
	TimeStamp      time.Time `firestore:"time" json:"-"`
}

// generateOutPutStruct generates the output struct so that the output from this api is a flat structure.
func generateOutPutStruct(covidCases CovidCases) CovidCasesOutput {
	return CovidCasesOutput{
		Country:        covidCases.Country,
		Date:           covidCases.MostRecent.Date,
		ConfirmedCases: covidCases.MostRecent.ConfirmedCases,
		Recovered:      covidCases.MostRecent.Recovered,
		Deaths:         covidCases.MostRecent.Deaths,
		GrowthRate:     covidCases.MostRecent.GrowthRate,
	}
}

// generateOutPutStructFromMap turns the map retrieved from firestore back into a struct.
func generateOutPutStructFromMap(inputData map[string]interface{}) CovidCasesOutput {
	return CovidCasesOutput{
		Country:        inputData["Country"].(string),
		Date:           inputData["Date"].(string),
		ConfirmedCases: int(inputData["ConfirmedCases"].(int64)),
		Recovered:      int(inputData["Recovered"].(int64)),
		Deaths:         int(inputData["Deaths"].(int64)),
		GrowthRate:     inputData["GrowthRate"].(float64),
	}
}
