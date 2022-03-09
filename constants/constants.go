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

// API urls

// URL params
const URL_PARAM_AND = "&"
const URL_PARAM_EQUALS = "="
const URL_PARAM_LIMIT = "limit"

// Web web_client timeout time
const CLIENT_TIMEOUT = 30 * time.Second
