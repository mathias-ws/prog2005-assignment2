package constants

// Program info
const PROGRAM_VERSION string = "v1"

// URL params

// URL_PARAM_AND for adding multiple parameters.
const URL_PARAM_AND string = "&"

// URL_PARAM_EQUALS for setting the value of a parameter.
const URL_PARAM_EQUALS string = "="

// URL_PARAM_LIMIT parameter name for setting the limit.
const URL_PARAM_LIMIT string = "limit"

// URL_COUNTRY_NAME_PARAM the urlutil parameter used by the user when searching for country.
const URL_COUNTRY_NAME_PARAM string = "country"

// URL_SCOPE_PARAMETER the urlutil parameter used by the user when searching for a given date.
const URL_SCOPE_PARAMETER string = "scope"

// URL_PARAMETER_WANTED_TIME_FORMAT is the date format used for the current time.
const URL_PARAMETER_WANTED_TIME_FORMAT string = "2006-01-02"

// CovidCasesBaseUrl covid cases api base urlutil
const CovidCasesBaseUrl string = "https://covid19-graphql.now.sh"

// CovidTrackerBaseUrl covidtracker api base urlutil
const CovidTrackerBaseUrl string = "https://covidtrackerapi.bsg.ox.ac.uk/api/"

// CovidTrackerEndpoint endpoint for doing the search.
const CovidTrackerEndpoint string = "v2/stringency/actions/"

// CovidCasesDBCollection is the collection name for the covid cases collection.
const CovidCasesDBCollection string = "covidcases"

// PolicyDBCollection is the collection name for the policy caching collection.
const PolicyDBCollection string = "stringency"
