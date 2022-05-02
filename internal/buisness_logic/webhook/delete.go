package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
)

func Delete(param map[string]string) (map[string]string, error) {
	err := database.DeleteDocument(constants.WebhookDbCollection, param[constants.UrlParameterWebhookId])
	if err != nil {
		return nil, err
	}

	return map[string]string{"message": "webhook deleted"}, nil
}
