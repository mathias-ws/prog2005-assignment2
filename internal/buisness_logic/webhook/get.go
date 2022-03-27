package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
)

// GetAllRegistered returns a slice of all webhooks.
func GetAllRegistered() ([]structs.WebHookRegistration, error) {
	return database.GetAllWebhooks(webhookDbCollection, "")
}

// GetOne returns one webhook found in the database.
func GetOne(urlParam map[string]string) (structs.WebHookRegistration, error) {
	var obtainedWebhook structs.WebHookRegistration

	database.GetDocument(webhookDbCollection, urlParam[constants.UrlParameterWebhookId], &obtainedWebhook)

	if (structs.WebHookRegistration{}) == obtainedWebhook {
		return structs.WebHookRegistration{}, custom_errors.GetWebhookNotFoundError()
	}

	return obtainedWebhook, nil
}
