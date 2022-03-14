package handlers

import (
	"assignment-2/custom_errors"
	"assignment-2/json_parser"
	"assignment-2/web_client"
	"net/http"
)

// PolicyHandler checks for the http method and handles the error appropriately.
func PolicyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestPolicy(w)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

func handleGetRequestPolicy(w http.ResponseWriter) {
	response, _ := web_client.GetResponseFromWebPage(
		"https://covidtrackerapi.bsg.ox.ac.uk/api/v2/stringency/actions/nor/2022-01-05")
	list := json_parser.DecodePolicyInfo(response)
	err := json_parser.Encode(w, list)

	// Checks for errors in the encoding process.
	if err != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
