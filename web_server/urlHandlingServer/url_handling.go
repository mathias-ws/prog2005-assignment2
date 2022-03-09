package urlHandlingServer

import (
	"assignment-2/constants"
	"assignment-2/custom_errors"
	"net/url"
	"strconv"
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
