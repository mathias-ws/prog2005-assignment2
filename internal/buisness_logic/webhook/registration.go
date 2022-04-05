package webhook

import (
	"assignment-2/internal/buisness_logic/counter"
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/hashing"
	"assignment-2/internal/json_parsing"
	"assignment-2/internal/structs"
	"fmt"
	"net/http"
)

func Register(request *http.Request) (string, error) {
	var registrationInfo structs.WebHookRegistration

	json_parsing.Decode(request, &registrationInfo)

	if (structs.WebHookRegistration{}) == registrationInfo {
		return "", custom_errors.GetFailedToDecode()
	}

	if registrationInfo.Calls == 0 || registrationInfo.Country == "" || registrationInfo.Url == "" {
		return "", custom_errors.GetMissingJsonFieldsError()
	}

	numberOfCounts, _ := counter.GetNumberOfTimes(registrationInfo.Country)

	registrationInfo.CallsAtRegistration = numberOfCounts

	err := database.WriteDocument(constants.WebhookDbCollection, fmt.Sprintf("%v", registrationInfo), registrationInfo)
	if err != nil {
		return "", err
	}

	return hashing.HashString(fmt.Sprintf("%v", registrationInfo))
}
