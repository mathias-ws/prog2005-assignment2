package logic

import (
	"assignment-2/constants"
	"assignment-2/model"
	"assignment-2/web_client"
	"fmt"
	"net/http"
	"time"
)

// Variable containing the start time of the web_server.
var startTime time.Time

// SetStartTime sets the startTime variable to the current time.
func SetStartTime() {
	startTime = time.Now()
}

// GetUptime returns the current time since the startTime variable was set (the uptime).
func getUptime() string {
	return fmt.Sprintf("%ds", int(time.Since(startTime).Seconds()))
}

// getStatusCode gets the status code from a webpage specified by an urlHandlingClient.
func getStatusCode(url string) (int, error) {
	statusCode, err := web_client.GetResponseFromWebPage(url)

	// Checks for errors when fetching the api.
	if err != nil {
		return 0, err
	}

	return statusCode.StatusCode, err
}

// GetStatusInfo gets the diagnosis information and turns it into a struct.
func GetStatusInfo() model.Status {
	// Gets the status codes.
	covidCasesApiStatusCode, errCases := getStatusCode(constants.CovidCasesBaseUrl)
	covidPolicyApiStatusCode, errPolicy := getStatusCode(constants.CovidTrackerEndpoint)

	// If the apis are unreachable set a proper error code.
	if errCases != nil {
		covidCasesApiStatusCode = http.StatusBadGateway
	}
	if errPolicy != nil {
		covidPolicyApiStatusCode = http.StatusBadGateway
	}

	return model.Status{
		CasesApiStatusCode:  covidCasesApiStatusCode,
		PolicyApiStatusCode: covidPolicyApiStatusCode,
		NumberOfWebhooks:    0,
		Uptime:              getUptime(),
		Version:             constants.PROGRAM_VERSION,
	}
}
