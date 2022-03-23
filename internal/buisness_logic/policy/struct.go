package policy

import "assignment-2/internal/constants"

// policyOutput represents the output structure of the policy endpoint.
type policyOutput struct {
	CountryCode string
	Scope       string
	Stringency  float64
	Policy      int
}

// policyInputFromApi represents the input data from the policy api.
type policyInputFromApi struct {
	StringencyData map[string]interface{}   `json:"stringencyData"`
	PolicyData     []map[string]interface{} `json:"policyActions"`
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
