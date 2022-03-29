package structs

import "time"

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

// CountryInfo represents the common name and the firestore time stamp.
type CountryInfo struct {
	Common    string    `json:"common" firestore:"name"`
	TimeStamp time.Time `firestore:"time" json:"-"`
}

// CountryNameStruct base output from the country api.
type CountryNameStruct struct {
	Name CountryInfo `json:"name"`
}

// CovidCases represents the output of the cases api.
type CovidCases struct {
	Country    string               `json:"name"`
	MostRecent CovidCasesMostRecent `json:"mostRecent"`
}

// CountryStruct represents the country in the nested structure in the input json.
type CountryStruct struct {
	Country CovidCases `json:"country"`
}

// DataStruct represents the data in the nested structure in the input json.
type DataStruct struct {
	Data CountryStruct `json:"data"`
}

// CovidCasesMostRecent represents the most recent field in the nested structure in the input json.
type CovidCasesMostRecent struct {
	Date           string  `json:"date"`
	ConfirmedCases int     `json:"confirmed"`
	Recovered      int     `json:"recovered"`
	Deaths         int     `json:"deaths"`
	GrowthRate     float64 `json:"growthRate"`
}

// CovidCasesOutput represents the flat output structure of the endpoint.
type CovidCasesOutput struct {
	Country        string    `json:"country" firestore:"country"`
	Date           string    `json:"date" firestore:"date"`
	ConfirmedCases int       `json:"confirmed" firestore:"confirmed"`
	Recovered      int       `json:"recovered" firestore:"recovered"`
	Deaths         int       `json:"deaths" firestore:"deaths"`
	GrowthRate     float64   `json:"growth_rate" firestore:"growth_rate"`
	TimeStamp      time.Time `firestore:"time" json:"-"`
}

// PolicyOutput represents the output structure of the policy endpoint.
type PolicyOutput struct {
	CountryCode string    `json:"country_code"`
	Scope       string    `json:"date_value"`
	Stringency  float64   `json:"stringency"`
	Policy      int       `json:"policies"`
	TimeStamp   time.Time `firestore:"time" json:"-"`
}

// PolicyInputFromApi represents the input data from the policy api.
type PolicyInputFromApi struct {
	StringencyData Stringency `json:"stringencyData"`
	PolicyData     []Policy   `json:"policyActions"`
}

type Stringency struct {
	DateValue        string  `json:"date_value"`
	CountryCode      string  `json:"country_code"`
	Confirmed        int     `json:"confirmed"`
	Deaths           int     `json:"deaths"`
	StringencyActual float64 `json:"stringency_actual"`
	Stringency       float64 `json:"stringency"`
}

type Policy struct {
	PolicyTypeCode          string `json:"policy_type_code"`
	PolicyValueDisplayField string `json:"policy_value_display_field"`
}
