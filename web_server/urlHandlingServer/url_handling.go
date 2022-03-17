package urlHandlingServer

import (
	"assignment-2/constants"
	"assignment-2/custom_errors"
	"assignment-2/utilities"
	"net/url"
	"strconv"
	"time"
)

// GetLimit returns the limit specified by the user.
func GetLimit(url *url.URL) (int, error) {
	obtainedQuery := url.Query()

	// Checks if the limit parameter exists.
	if obtainedQuery.Has(constants.URL_PARAM_LIMIT) {
		limit, err := strconv.Atoi(obtainedQuery[constants.URL_PARAM_LIMIT][0])

		// Checks that the value is valid (bigger than zero).
		if !(limit > 0) || err != nil {
			return 0, custom_errors.GetInvalidLimitError()
		}

		return limit, nil
	}

	return 0, nil
}

// GetUrlParametersPolicy checks that the given url contains the wanted parameters for the policy endpoint
// with valid and correct parameters. If no scope set, the current date is set.
func GetUrlParametersPolicy(url *url.URL) (map[string]string, error) {
	obtainedQuery := url.Query()
	parametersToReturn := map[string]string{}

	// Checks if the url has the country parameter
	if obtainedQuery.Has(constants.URL_COUNTRY_NAME_PARAM) {
		parametersToReturn[constants.URL_COUNTRY_NAME_PARAM] = obtainedQuery[constants.URL_COUNTRY_NAME_PARAM][0]

		if !utilities.CheckIfStringIsNotEmpty(parametersToReturn[constants.URL_COUNTRY_NAME_PARAM]) {
			return nil, custom_errors.GetParameterError()
		}
	} else {
		return nil, custom_errors.GetParameterError()
	}

	// Checks if url has the scope parameter
	if obtainedQuery.Has(constants.URL_SCOPE_PARAMETER) {
		parametersToReturn[constants.URL_SCOPE_PARAMETER] = obtainedQuery[constants.URL_SCOPE_PARAMETER][0]

		if !utilities.CheckIfValidDateFormat(parametersToReturn[constants.URL_SCOPE_PARAMETER]) {
			return nil, custom_errors.GetParameterError()
		}
	} else {
		parametersToReturn[constants.URL_SCOPE_PARAMETER] = time.Now().Format(constants.URL_PARAMETER_WANTED_TIME_FORMAT)
	}

	return parametersToReturn, nil
}
