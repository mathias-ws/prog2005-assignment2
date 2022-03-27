package webhook

import (
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/web_client"
	"encoding/json"
	"log"
)

// Trigger sends the request when a webhook is triggered.
func Trigger(webhookId string, country string, numberOfCalls int, url string) error {
	data := generateStruct(webhookId, country, numberOfCalls)

	jsonData, err := json.Marshal(data)

	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return custom_errors.GetFailedToEncode()
	}

	response, errRequest := web_client.PostRequest(url, jsonData)

	if errRequest != nil {
		return errRequest
	}

	// Checks if the response got a valid status code.
	if response.StatusCode < 200 && response.StatusCode >= 300 {
		log.Println("Error from server receiving webhook.")
		return custom_errors.GetUnableToSendRequestError()
	}

	return nil
}
