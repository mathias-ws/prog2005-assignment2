package webhook

import (
	"assignment-2/internal/buisness_logic/counter"
	"assignment-2/internal/buisness_logic/country"
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
	"assignment-2/internal/web_client"
	"encoding/json"
	"log"
)

// Check checks all the webhooks and sees if they need to be triggered.
func Check(countryName string) {
	if len(countryName) == 3 {
		var err error
		countryName, err = country.GetCountryNameFromCca3(countryName)

		if err != nil {
			return
		}
	}

	webhooks, err := database.GetAllWebhooks(constants.WebhookDbCollection, countryName)
	if err != nil {
		return
	}

	for _, webhook := range webhooks {
		go func(webhook structs.WebHookRegistration) {
			calls, err := counter.GetNumberOfTimes(webhook.Country)
			if err != nil {
				return
			}

			if (calls-webhook.CallsAtRegistration)%webhook.Calls == 0 {
				err := trigger(webhook.Id, webhook.Country, calls, webhook.Url)
				if err != nil {
					log.Println("Unable to trigger webhook.")
					return
				}
			}
		}(webhook)
	}
}

// trigger sends the request when a webhook is triggered.
func trigger(webhookId string, country string, numberOfCalls int, url string) error {
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
