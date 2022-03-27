package country

import "time"

// countryInfo represents the common name and the firestore time stamp.
type countryInfo struct {
	Common    string    `json:"common" firestore:"name"`
	TimeStamp time.Time `firestore:"time" json:"-"`
}

// countryStruct base output from the country api.
type countryStruct struct {
	Name countryInfo `json:"name"`
}
