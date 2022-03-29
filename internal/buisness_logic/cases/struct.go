package cases

import (
	"assignment-2/internal/structs"
)

// generateOutPutStruct generates the output struct so that the output from this api is a flat structure.
func generateOutPutStruct(covidCases structs.CovidCases) structs.CovidCasesOutput {
	return structs.CovidCasesOutput{
		Country:        covidCases.Country,
		Date:           covidCases.MostRecent.Date,
		ConfirmedCases: covidCases.MostRecent.ConfirmedCases,
		Recovered:      covidCases.MostRecent.Recovered,
		Deaths:         covidCases.MostRecent.Deaths,
		GrowthRate:     covidCases.MostRecent.GrowthRate,
	}
}

// generateOutPutStructFromMap turns the map retrieved from firestore back into a struct.
func generateOutPutStructFromMap(inputData map[string]interface{}) structs.CovidCasesOutput {
	return structs.CovidCasesOutput{
		Country:        inputData["Country"].(string),
		Date:           inputData["Date"].(string),
		ConfirmedCases: int(inputData["ConfirmedCases"].(int64)),
		Recovered:      int(inputData["Recovered"].(int64)),
		Deaths:         int(inputData["Deaths"].(int64)),
		GrowthRate:     inputData["GrowthRate"].(float64),
	}
}
