package handlers

import (
	"assignment-2/internal/custom_errors"
	"net/http"
)

// DefaultHandler checks for the http method and handles the error appropriately.
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestDefault(w)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

// handleGetRequestDefault handles the get requests for the endpoint. Gives a 404 when the path doesn't exist.
// Is the default handler of the project.
func handleGetRequestDefault(w http.ResponseWriter) {
	custom_errors.HttpNotFound(w)
}
