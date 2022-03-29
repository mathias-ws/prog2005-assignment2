package policy

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/structs"
)

// generateOutputStruct generates a populated policyOutput struct.
func generateOutputStruct(inputStruct structs.PolicyInputFromApi, parameters map[string]string) structs.PolicyOutput {
	return structs.PolicyOutput{
		CountryCode: parameters[constants.URL_COUNTRY_NAME_PARAM],
		Scope:       parameters[constants.URL_SCOPE_PARAMETER],
		Stringency:  getStringency(inputStruct.StringencyData),
		Policy:      countPolicies(inputStruct.PolicyData),
	}
}
