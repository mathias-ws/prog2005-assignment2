package cases

import (
	"assignment-2/internal/buisness_logic/counter"
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/web_client"
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

// GetCovidCases takes in the url parameter and uses it to query the backend api and builds the output struct.
func GetCovidCases(urlParameters map[string]string) (CovidCasesOutput, error) {
	country := urlParameters[constants.URL_COUNTRY_NAME_PARAM]

	// Checks if the country is in the cache.
	var dataFromDatabase CovidCasesOutput

	database.GetDocument(constants.CovidCasesDBCollection,
		country, &dataFromDatabase)

	if (CovidCasesOutput{}) != dataFromDatabase {
		// Counts up the number of times the country has been searched.
		go func() {
			err := counter.CountUp(dataFromDatabase.Country)

			if err != nil {
				log.Printf("Error counting up the number of searches: %v", err)
			}
		}()

		return dataFromDatabase, nil
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
	}()

	return outputStruct, nil
}
