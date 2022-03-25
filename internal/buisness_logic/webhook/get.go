package webhook

import (
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
)

func GetAllRegistered() ([]structs.WebHookRegistration, error) {
	return database.GetAllWebhooks(webhookDbCollection)
}
