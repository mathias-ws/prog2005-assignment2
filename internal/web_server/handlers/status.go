package handlers

import (
	"assignment-2/internal/buisness_logic/status"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/web_server/json"
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
	err := json.Encode(w, status.GetStatusInfo())

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
