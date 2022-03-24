package webhook

// webHookRegistration represents the structure of the webhook registration.
type webHookRegistration struct {
	Url     string `json:"url"`
	Country string `json:"country"`
	Calls   int    `json:"calls"`
}
