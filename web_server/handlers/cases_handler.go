package handlers

import (
	"assignment-2/cases"
	"assignment-2/custom_errors"
	"assignment-2/web_server/json_parsing"
	"assignment-2/web_server/urlHandlingServer"
	"net/http"
)

// CovidCasesHandler checks for the http method and handles the error appropriately.
func CovidCasesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestCases(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

// handleGetRequestCases handles the get request for the covid cases endpoint.
func handleGetRequestCases(w http.ResponseWriter, r *http.Request) {
	urlParameters, errParameters := urlHandlingServer.GetUrlParametersPolicy(r.URL)

	// Checks for errors finding the url parameters.
	if errParameters != nil {
		custom_errors.HttpSearchParameters(w)
		return
	}

	covidCases, errCases := cases.GetCovidCases(urlParameters)

	// Checks for errors in the process of getting the covid cases information.
	if errCases != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}

	errEncoding := json_parsing.Encode(w, covidCases)

	// Checks for errors in the encoding process.
	if errEncoding != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
