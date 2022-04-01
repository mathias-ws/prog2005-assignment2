package status

import (
	"assignment-2/internal/buisness_logic/cases"
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/web_client"
	"fmt"
	"net/http"
	"time"
)

// Variable containing the start time of the web_server.
var startTime = time.Now()

// getUptime returns the current time since the startTime variable was set (the uptime)
// in hours, minutes and seconds.
func getUptime() string {
	deltaTime := int(time.Since(startTime).Seconds())

	hours := (deltaTime / (60 * 60)) % 60
	minutes := (deltaTime / 60) % 60
	seconds := deltaTime % 60

	return fmt.Sprintf("%dh %dm %ds", hours, minutes, seconds)
}

// getNumberOfWebhooks gets all the webhooks and returns the length of the slice.
func getNumberOfWebhooks() (int, error) {
	webhooks, err := database.GetAllWebhooks(constants.WebhookDbCollection, "")

	if err != nil {
		return -1, err
	}

	return len(webhooks), err
}

// GetStatusInfo gets the diagnosis information and turns it into a struct.
func GetStatusInfo() status {
	// Gets the status codes.
	covidCasesApiStatusCode, errCases := cases.GetStatusCode()
	covidPolicyApiStatusCode, errPolicy := web_client.GetStatusCode(constants.CovidTrackerBaseUrl)
	countryApiStatusCode, errCountry := web_client.GetStatusCode(constants.CountryApiUrl)
	numberOfWebhooks, _ := getNumberOfWebhooks()

	// If the apis are unreachable set a proper error code.
	if errCases != nil {
		covidCasesApiStatusCode = http.StatusBadGateway
	}
	if errPolicy != nil {
		covidPolicyApiStatusCode = http.StatusBadGateway
	}
	if errCountry != nil {
		countryApiStatusCode = http.StatusBadGateway
	}

	return generateOutputStruct(covidCasesApiStatusCode, covidPolicyApiStatusCode, countryApiStatusCode,
		numberOfWebhooks, getUptime())
}
