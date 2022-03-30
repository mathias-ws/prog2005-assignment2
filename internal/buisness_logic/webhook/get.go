package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
)

// GetAllRegistered returns a slice of all webhooks.
func GetAllRegistered() ([]structs.WebHookRegistration, error) {
	return database.GetAllWebhooks(constants.WebhookDbCollection, "")
}

// GetOne returns one webhook found in the database.
func GetOne(urlParam map[string]string) (structs.WebHookRegistration, error) {
	var obtainedWebhook structs.WebHookRegistration

	database.GetDocument(constants.WebhookDbCollection, urlParam[constants.UrlParameterWebhookId], &obtainedWebhook)

	if (structs.WebHookRegistration{}) == obtainedWebhook {
		return structs.WebHookRegistration{}, custom_errors.GetWebhookNotFoundError()
	}

	return obtainedWebhook, nil
}
