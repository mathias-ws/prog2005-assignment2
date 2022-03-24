package webhook

import (
	"assignment-2/internal/database"
	"fmt"
	"net/http"
)

func Register(request *http.Request) {
	registrationInfo := decodeWebHookRegistration(request)

	err := database.WriteToDatabase(webhookDbCollection, fmt.Sprintf("%v", registrationInfo), registrationInfo)
	if err != nil {
		return
	}
}
