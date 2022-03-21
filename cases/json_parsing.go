package cases

import (
	"encoding/json"
	"log"
	"net/http"
)

// decodeCovidCases decodes the covid cases info into the CovidCases struct.
func decodeCovidCases(httpResponse *http.Response) CovidCases {
	decoder := json.NewDecoder(httpResponse.Body)
	var covidCases Data

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&covidCases); err != nil {
		log.Println(err)
	}

	return covidCases.Data.Country
}
