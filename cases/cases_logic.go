package cases

import (
	"assignment-2/constants"
	"assignment-2/custom_errors"
	"assignment-2/web_client"
	"encoding/json"
	"log"
)

// createGraphQlRequest Generates the body for the graphql request to the backend api.
func createGraphQlRequest(country string) ([]byte, error) {

	// Generates the graphql query.
	jsonData := map[string]string{
		"query": `query {
    country(name: "` + country + `") {
        name
        mostRecent {
            date(format: "yyyy-MM-dd")
            confirmed
            recovered
            deaths
            growthRate
        }
    }
}`,
	}

	jsonValue, err := json.Marshal(jsonData)

	if err != nil {
		log.Println(err)
		return nil, custom_errors.GetFailedToDecode()
	}

	return jsonValue, nil
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

// GetCovidCases takes in the url parameter and uses it to query the backend api and builds the output struct.
func GetCovidCases(urlParameters map[string]string) (CovidCasesOutput, error) {
	requestBody, err := createGraphQlRequest(urlParameters[constants.URL_COUNTRY_NAME_PARAM])

	if err != nil {
		return CovidCasesOutput{}, err
	}

	response, err := web_client.PostRequest(constants.CovidCasesBaseUrl, requestBody)

	if err != nil {
		return CovidCasesOutput{}, err
	}

	return generateOutPutStruct(decodeCovidCases(response)), nil
}
