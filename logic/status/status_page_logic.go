package status

import (
	"assignment-2/constants"
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

// getUptime returns the current time since the startTime variable was set (the uptime)
// in hours, minutes and seconds.
func getUptime() string {
	deltaTime := int(time.Since(startTime).Seconds())

	hours := (deltaTime / (60 * 60)) % 60
	minutes := (deltaTime / 60) % 60
	seconds := deltaTime % 60

	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
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
func GetStatusInfo() status {
	// Gets the status codes.
	covidCasesApiStatusCode, errCases := getStatusCode(constants.CovidCasesBaseUrl)
	covidPolicyApiStatusCode, errPolicy := getStatusCode(constants.CovidTrackerBaseUrl)

	// If the apis are unreachable set a proper error code.
	if errCases != nil {
		covidCasesApiStatusCode = http.StatusBadGateway
	}
	if errPolicy != nil {
		covidPolicyApiStatusCode = http.StatusBadGateway
	}

	return status{
		CasesApiStatusCode:  covidCasesApiStatusCode,
		PolicyApiStatusCode: covidPolicyApiStatusCode,
		NumberOfWebhooks:    0,
		Uptime:              getUptime(),
		Version:             constants.PROGRAM_VERSION,
	}
}
