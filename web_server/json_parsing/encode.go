package json_parsing

import (
	"encoding/json"
	"log"
	"net/http"
)

// Encode encodes some data into json and displays it on the website. The input can be anything.
func Encode(w http.ResponseWriter, valueToEncode interface{}) error {
	w.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(w)

	// Ensures that the json has a pretty format when viewed.
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(valueToEncode); err != nil {
		log.Println("Unable to encode data: ", err)
		return err
	}

	return nil
}
