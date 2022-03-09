package handlers

import (
	"assignment-2/custom_errors"
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

}
