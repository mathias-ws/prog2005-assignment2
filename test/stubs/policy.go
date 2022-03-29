package stubs

import (
	"assignment-2/internal/structs"
	"encoding/json"
	"net/http"
)

// PolicyHandler returns the json output mimicking the policy api.
func PolicyHandler(w http.ResponseWriter, r *http.Request) {
	jsonData := structs.PolicyInputFromApi{
		StringencyData: structs.Stringency{
			Deaths:           1664,
			Stringency:       13.89,
			CountryCode:      "NOR",
			StringencyActual: 13.89,
			Confirmed:        1299549,
			DateValue:        "2022-03-04",
		}, PolicyData: []structs.Policy{
			{
				PolicyTypeCode:          "C1",
				PolicyValueDisplayField: "No measures",
			},
			{
				PolicyTypeCode:          "C2",
				PolicyValueDisplayField: "No measures",
			},
		},
	}

	encoder := json.NewEncoder(w)

	err := encoder.Encode(jsonData)
	if err != nil {
		return
	}
}
