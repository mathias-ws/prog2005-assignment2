package handlers

import (
	"assignment-2/internal/buisness_logic/policy"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/web_server/json"
	"assignment-2/internal/web_server/urlutil"
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
	urlParameters, errParameters := urlutil.GetUrlParametersPolicy(r.URL)

	// Checks for errors in the encoding process.
	if errParameters != nil {
		custom_errors.HttpSearchParameters(w)
		return
	}

	policyInformation, errPolicy := policy.FindPolicyInformation(urlParameters)

	// Checks for errors in the process of getting the policy and stringency information.
	if errPolicy != nil {
		if errPolicy.Error() == custom_errors.GetFailedToDecode().Error() {
			custom_errors.HttpNoPolicy(w)
			return
		}

		custom_errors.HttpUnknownServerError(w)
		return
	}

	errEncoding := json.Encode(w, policyInformation)

	// Checks for errors in the encoding process.
	if errEncoding != nil {
		custom_errors.HttpUnknownServerError(w)
		return
	}
}
