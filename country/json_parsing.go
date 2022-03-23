package country

import (
	"encoding/json"
	"log"
	"net/http"
)

// decodeCountryInfo decodes the country info into the struct.
func decodeCountryInfo(httpResponse *http.Response) countryStruct {
	decoder := json.NewDecoder(httpResponse.Body)
	var obtainedCountry []countryStruct

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&obtainedCountry); err != nil {
		log.Printf("Error parsing json: %v", err)
	}

	return obtainedCountry[0]
}
