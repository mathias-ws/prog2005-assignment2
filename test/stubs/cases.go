package stubs

import (
	"assignment-2/internal/structs"
	"encoding/json"
	"net/http"
)

// CasesHandler returns the json output mimicking the cases api.
func CasesHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := structs.DataStruct{
		Data: structs.CountryStruct{
			Country: structs.CovidCases{
				Country: "Norway",
				MostRecent: structs.CovidCasesMostRecent{
					Date:           "2022-03-28",
					ConfirmedCases: 1399714,
					Deaths:         2339,
					GrowthRate:     0.0014853631627073677,
					Recovered:      0,
				},
			},
		},
	}

	encoder := json.NewEncoder(w)

	err := encoder.Encode(jsonData)
	if err != nil {
		return
	}
}
