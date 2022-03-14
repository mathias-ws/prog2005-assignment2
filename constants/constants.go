package constants

import "time"

// Program info
const PORT = "80"
const PROGRAM_VERSION = "v1"

// URL paths
const CASES_LOCATION = "/corona/v1/cases"
const POLICY_LOCATION = "/corona/v1/policy"
const STATUS_LOCATION = "/corona/v1/status"
const WEBHOOK_LOCATION = "/corona/v1/notifications/"
const DEFAULT_LOCATION = "/"

// Covid cases api
const COVID_CASES_BASE_URL = "https://covid19-graphql.now.sh"

// Covid tracker api
const COVID_TRACKER_BASE_URL = "https://covidtrackerapi.bsg.ox.ac.uk/api"

// CLIENT_TIMEOUT web_client timeout time
const CLIENT_TIMEOUT = 30 * time.Second

// URL params

// URL_PARAM_AND for adding multiple parameters.
const URL_PARAM_AND = "&"

// URL_PARAM_EQUALS for setting the value of a parameter.
const URL_PARAM_EQUALS = "="

// URL_PARAM_LIMIT parameter name for setting the limit.
const URL_PARAM_LIMIT = "limit"

// URL_COUNTRY_NAME_PARAM the url parameter used by the user when searching for country.
const URL_COUNTRY_NAME_PARAM = "country"

// URL_SCOPE_PARAMETER the url parameter used by the user when searching for a given date.
const URL_SCOPE_PARAMETER = "scope"

// REGEX_CHECK_VALID_DATE regex expression for checking if string is valid date format.
// Inspired by: https://www.golangprograms.com/regular-expression-to-validate-the-date-format-in-dd-mm-yyyy.html
const REGEX_CHECK_VALID_DATE = "((19|20)\\d\\d)-(0?[1-9]|1[012])-(0?[1-9]|[12][0-9]|3[01])"

// REGEX_CHECK_VALID_STRING regex expression that checks that a string contains small or capital letters.
const REGEX_CHECK_VALID_STRING = "^[a-zA-Z]+$"

// URL_PARAMETER_WANTED_TIME_FORMAT is the date format used for the current time.
const URL_PARAMETER_WANTED_TIME_FORMAT = "2006-01-02"
