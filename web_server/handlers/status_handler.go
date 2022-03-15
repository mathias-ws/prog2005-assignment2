package handlers

import (
	"assignment-2/custom_errors"
	"assignment-2/json_parser"
	"assignment-2/logic/status"
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
	err := json_parser.Encode(w, status.GetStatusInfo())

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
