package json_parsing

import (
	"encoding/json"
	"log"
	"net/http"
)

// Decode decodes the input body into the struct.
func Decode(input interface{}, dataStruct interface{}) {
	var decoder *json.Decoder

	// Switch case that finds the type of the input object.
	switch input.(type) {
	case *http.Response:
		decoder = json.NewDecoder(input.(*http.Response).Body)
	case *http.Request:
		decoder = json.NewDecoder(input.(*http.Request).Body)
	default:
		log.Println("Unsupported object to decode.")
		return
	}

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&dataStruct); err != nil {
		log.Println(err)
	}
}
