package policy_endpoint

// policyOutput represents the output structure of the policy_endpoint endpoint.
type policyOutput struct {
	CountryCode string
	Scope       string
	Stringency  float64
	Policy      int
}

// policyInputFromApi represents the input data from the policy_endpoint api.
type policyInputFromApi struct {
	StringencyData map[string]interface{}   `json:"stringencyData"`
	PolicyData     []map[string]interface{} `json:"policyActions"`
}
