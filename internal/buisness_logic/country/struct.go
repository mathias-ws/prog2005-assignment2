package country

import "time"

// countryStruct base output from the country api.
type countryStruct struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
	TimeStamp time.Time `firestore:"time" json:"-"`
}
