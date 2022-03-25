package country

import "time"

// name struct to represent the name object in the json output from the country api.
type name struct {
	Common string `json:"common"`
}

// countryStruct base output from the country api.
type countryStruct struct {
	Name      name      `json:"name"`
	TimeStamp time.Time `firestore:"time" json:"-"`
}
