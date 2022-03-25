package structs

// WebHookRegistration represents the structure of the webhook registration.
type WebHookRegistration struct {
	Id      string `json:"id,omitempty"`
	Url     string `json:"url"`
	Country string `json:"country"`
	Calls   int    `json:"calls"`
}
