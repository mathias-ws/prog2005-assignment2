package policy

import (
	"assignment-2/internal/constants"
	"time"
)

// policyOutput represents the output structure of the policy endpoint.
type policyOutput struct {
	CountryCode string    `json:"country_code"`
	Scope       string    `json:"date_value"`
	Stringency  float64   `json:"stringency"`
	Policy      int       `json:"policies"`
	TimeStamp   time.Time `firestore:"time" json:"-"`
}

// policyInputFromApi represents the input data from the policy api.
type policyInputFromApi struct {
	StringencyData stringency `json:"stringencyData"`
	PolicyData     []policy   `json:"policyActions"`
}

type stringency struct {
	DateValue        string  `json:"date_value"`
	CountryCode      string  `json:"country_code"`
	Confirmed        int     `json:"confirmed"`
	Deaths           int     `json:"deaths"`
	StringencyActual float64 `json:"stringency_actual"`
	Stringency       float64 `json:"stringency"`
}

type policy struct {
	PolicyTypeCode          string `json:"policy_type_code"`
	PolicyValueDisplayField string `json:"policy_value_display_field"`
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
