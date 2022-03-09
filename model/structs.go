package model

type Status struct {
	CasesApiStatusCode  int    `json:"cases_api_status"`
	PolicyApiStatusCode int    `json:"policy_api_status"`
	NumberOfWebhooks    int    `json:"number_of_webhooks"`
	Version             string `json:"version"`
	Uptime              string `json:"uptime"`
}
