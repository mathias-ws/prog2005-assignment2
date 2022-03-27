package structs

// WebHookRegistration represents the structure of the webhook registration.
type WebHookRegistration struct {
	Id                  string `json:"id,omitempty" firestore:"-"`
	Url                 string `json:"url"`
	Country             string `json:"country"`
	Calls               int    `json:"calls"`
	CallsAtRegistration int    `json:"-"`
}

// CountryCounter represents the counter of a country.
type CountryCounter struct {
	Count int `firebase:"count"`
}
