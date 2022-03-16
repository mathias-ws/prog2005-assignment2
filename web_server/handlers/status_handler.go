package handlers

import (
	"assignment-2/custom_errors"
	"assignment-2/status_endpoint"
	"assignment-2/web_server/json_parsing"
	"net/http"
)

// StatusHandler checks for the http method and handles the error appropriately.
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestStatus(w)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

func handleGetRequestStatus(w http.ResponseWriter) {
	err := json_parsing.Encode(w, status_endpoint.GetStatusInfo())

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
