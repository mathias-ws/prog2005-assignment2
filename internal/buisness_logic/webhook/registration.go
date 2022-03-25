package webhook

import (
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/hashing"
	"assignment-2/internal/json_parsing"
	"fmt"
	"net/http"
)

func Register(request *http.Request) (string, error) {
	var registrationInfo webHookRegistration

	json_parsing.Decode(request, &registrationInfo)

	if (webHookRegistration{}) == registrationInfo {
		return "", custom_errors.GetFailedToDecode()
	}

	err := database.WriteToDatabase(webhookDbCollection, fmt.Sprintf("%v", registrationInfo), registrationInfo)
	if err != nil {
		return "", err
	}

	return hashing.HashString(fmt.Sprintf("%v", registrationInfo))
}
