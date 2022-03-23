package status

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/web_client"
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

// GetStatusInfo gets the diagnosis information and turns it into a struct.
func GetStatusInfo() status {
	// Gets the status codes.
	covidCasesApiStatusCode, errCases := web_client.GetStatusCode(constants.CovidCasesBaseUrl)
	covidPolicyApiStatusCode, errPolicy := web_client.GetStatusCode(constants.CovidTrackerBaseUrl)

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
