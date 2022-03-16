package policy_endpoint

import (
	"assignment-2/constants"
	"assignment-2/custom_errors"
	"assignment-2/web_client"
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
	} else {
		return stringency["stringency"].(float64)
	}
}

// stringencyIsValid checks if there is any stringency information in the struct.
func stringencyIsValid(stringency map[string]interface{}) bool {
	if stringency["msg"] != nil {
		return false
	} else {
		return true
	}
}

// generateOutputStruct generates a policyOutput struct if the stringency information is valid.
func generateOutputStruct(inputStruct policyInputFromApi) (policyOutput, error) {
	if !stringencyIsValid(inputStruct.StringencyData) {
		return policyOutput{}, custom_errors.GetNoContentStringencyFoundError()
	}

	return policyOutput{
		CountryCode: inputStruct.StringencyData["country_code"].(string),
		Scope:       inputStruct.StringencyData["date_value"].(string),
		Stringency:  getStringency(inputStruct.StringencyData),
		Policy:      countPolicies(inputStruct.PolicyData),
	}, nil
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
	response, err := web_client.GetResponseFromWebPage(buildSearchUrl(urlParameters))

	if err != nil {
		return policyOutput{}, err
	}

	obtainedPolicyInformation := decodePolicyInfo(response)

	return generateOutputStruct(obtainedPolicyInformation)
}
