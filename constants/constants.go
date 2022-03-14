package constants

import "time"

// Program info
const PORT string = "80"
const PROGRAM_VERSION string = "v1"

// URL paths
const CASES_LOCATION string = "/corona/v1/cases"
const POLICY_LOCATION string = "/corona/v1/policy"
const STATUS_LOCATION string = "/corona/v1/status"
const WEBHOOK_LOCATION string = "/corona/v1/notifications/"
const DEFAULT_LOCATION string = "/"

// Covid cases api
const COVID_CASES_BASE_URL string = "https://covid19-graphql.now.sh"

// Covid tracker api
const COVID_TRACKER_BASE_URL string = "https://covidtrackerapi.bsg.ox.ac.uk/api/"

// COVID_TRACKER_ENDPOINT endpoint for doing the search.
const COVID_TRACKER_ENDPOINT string = "v2/stringency/actions/"

// CLIENT_TIMEOUT web_client timeout time
const CLIENT_TIMEOUT = 30 * time.Second

// URL params

// URL_PARAM_AND for adding multiple parameters.
const URL_PARAM_AND string = "&"

// URL_PARAM_EQUALS for setting the value of a parameter.
const URL_PARAM_EQUALS string = "="

// URL_PARAM_LIMIT parameter name for setting the limit.
const URL_PARAM_LIMIT string = "limit"

// URL_COUNTRY_NAME_PARAM the url parameter used by the user when searching for country.
const URL_COUNTRY_NAME_PARAM string = "country"

// URL_SCOPE_PARAMETER the url parameter used by the user when searching for a given date.
const URL_SCOPE_PARAMETER string = "scope"

// REGEX_CHECK_VALID_DATE regex expression for checking if string is valid date format.
// Inspired by: https://www.golangprograms.com/regular-expression-to-validate-the-date-format-in-dd-mm-yyyy.html
const REGEX_CHECK_VALID_DATE string = "((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])"

// REGEX_CHECK_VALID_STRING regex expression that checks that a string contains small or capital letters.
const REGEX_CHECK_VALID_STRING string = "^[a-zA-Z]+$"

// URL_PARAMETER_WANTED_TIME_FORMAT is the date format used for the current time.
const URL_PARAMETER_WANTED_TIME_FORMAT string = "2006-01-02"
