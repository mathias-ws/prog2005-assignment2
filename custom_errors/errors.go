package custom_errors

import "errors"

// Structure inspired by this stackoverflow thread:
//https://stackoverflow.com/questions/38361919/how-do-i-cleanly-separate-user-facing-errors-from-internal-errors-in-golang

// GetUnableToReachBackendApisError returns the error message for the web_client side error from the apis.
func GetUnableToReachBackendApisError() error {
	return errors.New("error sending request or getting response from the api")
}

// GetNoContentStringencyFoundError returns the error message for when the given stringency information was not found.
func GetNoContentStringencyFoundError() error {
	return errors.New("no stringency information with the given search parameters were not found")
}

//GetParameterError returns error message for when the user has not provided the mandatory parameters.
func GetParameterError() error {
	return errors.New("missing parameters or wrong parameters")
}

//GetInvalidLimitError returns error message for when the user has not used a proper limit.
func GetInvalidLimitError() error {
	return errors.New("invalid limit")
}
