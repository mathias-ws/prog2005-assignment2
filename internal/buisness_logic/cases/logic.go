package cases

import (
	"assignment-2/internal/buisness_logic/counter"
	"assignment-2/internal/buisness_logic/webhook"
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/json_parsing"
	"assignment-2/internal/structs"
	"assignment-2/internal/web_client"
	"encoding/json"
	"log"
	"time"
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
		return nil, custom_errors.GetFailedToEncode()
	}

	return jsonValue, nil
}

// GetCovidCases takes in the url parameter and uses it to query the backend api and builds the output struct.
func GetCovidCases(urlParameters map[string]string) (structs.CovidCasesOutput, error) {
	country := urlParameters[constants.URL_COUNTRY_NAME_PARAM]

	// Checks if the country is in the cache.
	var dataFromDatabase structs.CovidCasesOutput

	database.GetDocument(constants.CovidCasesDBCollection,
		country, &dataFromDatabase)

	if (structs.CovidCasesOutput{}) != dataFromDatabase {
		// Cache only for twelve hours.
		if !(time.Since(dataFromDatabase.TimeStamp).Hours() > (time.Hour * 12).Hours()) {
			// Counts up the number of times the country has been searched.
			go func() {
				err := counter.CountUp(dataFromDatabase.Country)

				if err != nil {
					log.Printf("Error counting up the number of searches: %v", err)
				}

				webhook.Check(dataFromDatabase.Country)
			}()

			return dataFromDatabase, nil
		}
	}

	requestBody, err := createGraphQlRequest(country)

	if err != nil {
		return structs.CovidCasesOutput{}, err
	}

	response, err := web_client.PostRequest(constants.CovidCasesBaseUrl, requestBody)

	if err != nil {
		return structs.CovidCasesOutput{}, err
	}

	covidInfo := structs.DataStruct{}
	json_parsing.Decode(response, &covidInfo)

	outputStruct := generateOutPutStruct(covidInfo.Data.Country)

	// Starts a new goroutine that caches the struct.
	go func() {
		err := database.WriteDocument(constants.CovidCasesDBCollection, outputStruct.Country, outputStruct)
		if err != nil {
			log.Printf("Error writing to cache: %v", err)
		}
	}()

	// Counts up the number of times the country has been searched.
	go func() {
		err := counter.CountUp(outputStruct.Country)

		if err != nil {
			log.Printf("Error counting up the number of searches: %v", err)
		}

		webhook.Check(outputStruct.Country)
	}()

	return outputStruct, nil
}

// GetStatusCode returns the status code of the cases api.
func GetStatusCode() (int, error) {
	request, err := createGraphQlRequest("Norway")
	if err != nil {
		return 0, err
	}

	response, err := web_client.PostRequest(constants.CovidCasesBaseUrl, request)

	if err != nil {
		return 0, err
	}

	return response.StatusCode, nil
}
