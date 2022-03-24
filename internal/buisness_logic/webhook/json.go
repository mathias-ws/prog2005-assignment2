package webhook

import (
	"encoding/json"
	"log"
	"net/http"
)

// decodeWebHookRegistration decodes the registration info into the struct.
func decodeWebHookRegistration(request *http.Request) webHookRegistration {
	decoder := json.NewDecoder(request.Body)
	var registration webHookRegistration

	// Checks for errors in the decoding process.
	if err := decoder.Decode(&registration); err != nil {
		log.Println(err)
	}

	return registration
}
