package handlers

import (
	"assignment-2/custom_errors"
	"assignment-2/policy_endpoint"
	"assignment-2/web_server/json_parsing"
	"assignment-2/web_server/urlHandlingServer"
	"net/http"
)

// PolicyHandler checks for the http method and handles the error appropriately.
func PolicyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetRequestPolicy(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

func handleGetRequestPolicy(w http.ResponseWriter, r *http.Request) {
	urlParameters, errParameters := urlHandlingServer.GetUrlParametersPolicy(r.URL)

	// Checks for errors in the encoding process.
	if errParameters != nil {
		custom_errors.HttpSearchParameters(w)
		return
	}

	policyInformation, errPolicy := policy_endpoint.FindPolicyInformation(urlParameters)

	// Checks for errors in the process of getting the policy and stringency information.
	if errPolicy != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}

	errEncoding := json_parsing.Encode(w, policyInformation)

	// Checks for errors in the encoding process.
	if errEncoding != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
