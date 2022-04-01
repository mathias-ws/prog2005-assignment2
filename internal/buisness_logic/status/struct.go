package status

import "assignment-2/internal/constants"

// status output struct for the status endpoint.
type status struct {
	CasesApiStatusCode   int    `json:"cases_api_status"`
	PolicyApiStatusCode  int    `json:"policy_api_status"`
	CountryApiStatusCode int    `json:"country_api_status"`
	NumberOfWebhooks     int    `json:"number_of_webhooks"`
	Version              string `json:"version"`
	Uptime               string `json:"uptime"`
}

// generateOutputStruct generates the output struct for the status endpoint.
func generateOutputStruct(covidCasesApiStatusCode int, covidPolicyApiStatusCode int,
	countryApiStatusCode int, numberOfWebHooks int, uptime string) status {
	return status{
		CasesApiStatusCode:   covidCasesApiStatusCode,
		PolicyApiStatusCode:  covidPolicyApiStatusCode,
		CountryApiStatusCode: countryApiStatusCode,
		NumberOfWebhooks:     numberOfWebHooks,
		Uptime:               uptime,
		Version:              constants.PROGRAM_VERSION,
	}
}
