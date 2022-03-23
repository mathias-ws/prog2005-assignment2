package cases

import (
	"assignment-2/constants"
	"assignment-2/custom_errors"
	"assignment-2/database"
	"assignment-2/web_client"
	"encoding/json"
	"log"
	"strings"
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

// GetCovidCases takes in the url parameter and uses it to query the backend api and builds the output struct.
func GetCovidCases(urlParameters map[string]string) (CovidCasesOutput, error) {
	country := strings.Title(strings.ToLower(urlParameters[constants.URL_COUNTRY_NAME_PARAM]))

	// Checks if the country is in the cache.
	dataFromDatabase := database.GetFromDatabase(constants.CovidCasesDBCollection,
		country)

	if len(dataFromDatabase) != 0 {
		return generateOutPutStructFromMap(dataFromDatabase), nil
	}

	requestBody, err := createGraphQlRequest(country)

	if err != nil {
		return CovidCasesOutput{}, err
	}

	response, err := web_client.PostRequest(constants.CovidCasesBaseUrl, requestBody)

	if err != nil {
		return CovidCasesOutput{}, err
	}

	outputStruct := generateOutPutStruct(decodeCovidCases(response))

	// Starts a new goroutine that caches the struct.
	go func() {
		err := database.WriteToDatabase(constants.CovidCasesDBCollection, outputStruct.Country, outputStruct)
		if err != nil {
			log.Printf("Error writing to cache: %v", err)
		}
	}()

	return outputStruct, nil
}
