package policy

import (
	"assignment-2/internal/buisness_logic/counter"
	"assignment-2/internal/buisness_logic/webhook"
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/json_parsing"
	"assignment-2/internal/web_client"
	"log"
	"strings"
)

// countPolicies counts the number of policies in place.
func countPolicies(policies []policy) int {
	if policies[0].PolicyTypeCode == "NONE" {
		return 0
	}

	return len(policies)
}

// getStringency checks if the stringency_actual is not null, if it is the stringency field is returned.
func getStringency(stringencyInfo stringency) float64 {
	if stringencyInfo.StringencyActual != 0 {
		return stringencyInfo.StringencyActual
	} else if stringencyInfo.Stringency != 0 {
		return stringencyInfo.Stringency
	} else {
		return defaultStringencyValue
	}
}

// buildSearchUrl builds the url that can be used to search the backend api.
func buildSearchUrl(parameters map[string]string) string {
	urlToSearch := strings.Builder{}

	urlToSearch.WriteString(constants.CovidTrackerBaseUrl)
	urlToSearch.WriteString(constants.CovidTrackerEndpoint)
	urlToSearch.WriteString(parameters[constants.URL_COUNTRY_NAME_PARAM])
	urlToSearch.WriteString("/")
	urlToSearch.WriteString(parameters[constants.URL_SCOPE_PARAMETER])

	return urlToSearch.String()
}

// FindPolicyInformation takes in a map of url parameter values and uses them to search the backend api
// and generates an output struct that is returned.
func FindPolicyInformation(urlParameters map[string]string) (policyOutput, error) {
	// Checks if the country is in the cache.
	var dataFromDatabase policyOutput
	database.GetDocument(constants.PolicyDBCollection,
		urlParameters[constants.URL_COUNTRY_NAME_PARAM]+urlParameters[constants.URL_SCOPE_PARAMETER], &dataFromDatabase)

	if (policyOutput{}) != dataFromDatabase {
		// Counts up the number of times the country has been searched.
		go func() {
			err := counter.CountUp(dataFromDatabase.CountryCode)

			if err != nil {
				log.Printf("Error counting up the number of searches: %v", err)
			}

			webhook.Check(dataFromDatabase.CountryCode)
		}()

		return dataFromDatabase, nil
	}

	response, err := web_client.GetRequest(buildSearchUrl(urlParameters))

	if err != nil {
		return policyOutput{}, err
	}

	var obtainedPolicyInformation policyInputFromApi

	json_parsing.Decode(response, &obtainedPolicyInformation)

	//TODO: if not all data is populated yet, it goes here... Intended behaviour?
	if (policyInputFromApi{}.StringencyData.CountryCode) == obtainedPolicyInformation.StringencyData.CountryCode {
		return policyOutput{}, custom_errors.GetFailedToDecode()
	}

	outputStruct := generateOutputStruct(obtainedPolicyInformation, urlParameters)

	// Starts a new goroutine that caches the struct.
	go func() {
		err := database.WriteDocument(constants.PolicyDBCollection,
			outputStruct.CountryCode+outputStruct.Scope, outputStruct)
		if err != nil {
			log.Printf("Error writing to cache: %v", err)
		}
	}()

	// Counts up the number of times the country has been searched.
	go func() {
		err := counter.CountUp(outputStruct.CountryCode)

		if err != nil {
			log.Printf("Error counting up the number of searches: %v", err)
		}

		webhook.Check(outputStruct.CountryCode)
	}()

	return outputStruct, nil
}
