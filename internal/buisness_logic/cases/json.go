package cases

import (
	"assignment-2/internal/structs"
	"encoding/json"
	"log"
	"net/http"
)

// decodeCovidCases decodes the covid cases info into the CovidCases struct.
func decodeCovidCases(httpResponse *http.Response) structs.CovidCases {
	decoder := json.NewDecoder(httpResponse.Body)
	var covidCases structs.DataStruct

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&covidCases); err != nil {
		log.Println(err)
	}

	return covidCases.Data.Country
}
