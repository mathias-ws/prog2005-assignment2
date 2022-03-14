package model

type Status struct {
	CasesApiStatusCode  int    `json:"cases_api_status"`
	PolicyApiStatusCode int    `json:"policy_api_status"`
	NumberOfWebhooks    int    `json:"number_of_webhooks"`
	Version             string `json:"version"`
	Uptime              string `json:"uptime"`
}

// PolicyOutput represents the output structure of the policy endpoint.
type PolicyOutput struct {
	CountryCode string
	Scope       string
	Stringency  float32
	Policy      int
}

// PolicyInputFromApi represents the input data from the policy api.
type PolicyInputFromApi struct {
	StringencyData map[string]interface{}   `json:"stringencyData"`
	PolicyData     []map[string]interface{} `json:"policyActions"`
}
