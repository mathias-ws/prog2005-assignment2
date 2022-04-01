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

// UrlParameterWebhookId is the url parm for the webhook id.
const UrlParameterWebhookId string = "id"

// CovidTrackerEndpoint endpoint for doing the search.
const CovidTrackerEndpoint string = "v2/stringency/actions/"

const WebhookTestCheckVerificationCollection string = "webhookclient"

var (
	// CovidCasesBaseUrl covid cases api base url
	CovidCasesBaseUrl = "https://covid19-graphql.now.sh"

	// CovidTrackerBaseUrl covidtracker api base url
	CovidTrackerBaseUrl = "https://covidtrackerapi.bsg.ox.ac.uk/api/"

	// CovidCasesDBCollection is the collection name for the covid cases collection.
	CovidCasesDBCollection = "covidcases"

	// PolicyDBCollection is the collection name for the policy caching collection.
	PolicyDBCollection = "stringency"

	CountryApiUrl       = "https://restcountries.com/v3.1/alpha/"
	CountryDbCollection = "cca3tocountry"

	// CounterDbCollection the name of the database collection.
	CounterDbCollection = "counter"

	WebhookDbCollection = "webhook"

	WebhookClientUrl string = ""
)

// SetTestUrlCases sets the test url for the cases api.
func SetTestUrlCases(testUrl string) {
	CovidCasesBaseUrl = testUrl
}

// SetTestUrlPolicy sets the test url for the policy api.
func SetTestUrlPolicy(testUrl string) {
	CovidTrackerBaseUrl = testUrl
}

// SetTestUrlCountry sets the test url for the country api.
func SetTestUrlCountry(testUrl string) {
	CountryApiUrl = testUrl
}

// SetTestUrlWebhookClient sets the test url for the webhook client.
func SetTestUrlWebhookClient(testUrl string) {
	WebhookClientUrl = testUrl
}

//SetTestDB sets the db collections to test mode.
func SetTestDB() {
	PolicyDBCollection = "teststringency"
	CovidCasesDBCollection = "testcovidcases"
	CountryDbCollection = "testcca3tocountry"
	CounterDbCollection = "testcounter"
	WebhookDbCollection = "testwebhook"
}
