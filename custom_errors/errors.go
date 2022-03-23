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

// GetFailedToDecode returns error message for when the decoder/unmarshall is unable to convert the object.
func GetFailedToDecode() error {
	return errors.New("failed to decode the object")
}

// GetDatabaseError returns error message for when there is an error accessing the database.
func GetDatabaseError() error {
	return errors.New("error accessing the database")
}
