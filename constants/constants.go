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

// URL params
const URL_PARAM_AND = "&"
const URL_PARAM_EQUALS = "="
const URL_PARAM_LIMIT = "limit"

// CLIENT_TIMEOUT web_client timeout time
const CLIENT_TIMEOUT = 30 * time.Second
