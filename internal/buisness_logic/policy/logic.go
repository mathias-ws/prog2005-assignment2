package policy

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/web_client"
	"log"
	"strings"
)

// countPolicies counts the number of policies in place.
func countPolicies(policies []map[string]interface{}) int {
	if policies[0]["policy_type_code"] == "NONE" {
		return 0
	}

	return len(policies)
}

// getStringency checks if the stringency_actual is not null, if it is the stringency field is returned.
func getStringency(stringency map[string]interface{}) float64 {
	if stringency["stringency_actual"] != nil {
		return stringency["stringency_actual"].(float64)
	} else if stringency["stringency_actual"] != nil {
		return stringency["stringency"].(float64)
	} else {
		return defaultStringencyValue
	}
}

// generateOutputStruct generates a populated policyOutput struct.
func generateOutputStruct(inputStruct policyInputFromApi, parameters map[string]string) policyOutput {
	return policyOutput{
		CountryCode: parameters[constants.URL_COUNTRY_NAME_PARAM],
		Scope:       parameters[constants.URL_SCOPE_PARAMETER],
		Stringency:  getStringency(inputStruct.StringencyData),
		Policy:      countPolicies(inputStruct.PolicyData),
	}
}

// generateOutPutStructFromMap turns the map retrieved from firestore back into a struct.
func generateOutPutStructFromMap(inputData map[string]interface{}) policyOutput {
	return policyOutput{
		CountryCode: inputData["CountryCode"].(string),
		Scope:       inputData["Scope"].(string),
		Policy:      int(inputData["Policy"].(int64)),
		Stringency:  inputData["Stringency"].(float64),
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
	dataFromDatabase := database.GetFromDatabase(constants.PolicyDBCollection,
		urlParameters[constants.URL_COUNTRY_NAME_PARAM]+urlParameters[constants.URL_SCOPE_PARAMETER])

	if len(dataFromDatabase) != 0 {
		return generateOutPutStructFromMap(dataFromDatabase), nil
	}

	response, err := web_client.GetRequest(buildSearchUrl(urlParameters))

	if err != nil {
		return policyOutput{}, err
	}

	obtainedPolicyInformation := decodePolicyInfo(response)

	outputStruct := generateOutputStruct(obtainedPolicyInformation, urlParameters)

	// Starts a new goroutine that caches the struct.
	go func() {
		err := database.WriteToDatabase(constants.PolicyDBCollection,
			outputStruct.CountryCode+outputStruct.Scope, outputStruct)
		if err != nil {
			log.Printf("Error writing to cache: %v", err)
		}
	}()

	return outputStruct, nil
}
