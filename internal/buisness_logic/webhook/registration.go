package webhook

import (
	"assignment-2/internal/database"
	"net/http"
)

func Register(request *http.Request) {
	registrationInfo := decodeWebHookRegistration(request)

	err := database.WriteToDatabase("registration", registrationInfo.Url, registrationInfo)
	if err != nil {
		return
	}
}
