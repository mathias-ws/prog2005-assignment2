package stubs

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/json_parsing"
	"assignment-2/internal/structs"
	"net/http"
)

func Webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		calledData := structs.WebHookPost{}

		json_parsing.Decode(r, &calledData)
		err := database.WriteDocument(constants.WebhookTestCheckVerificationCollection,
			"test", calledData)
		if err != nil {
			return
		}

		w.Header().Set("content-type", "application/json")

		w.WriteHeader(http.StatusOK)
		return
	}

	http.Error(w, "Something is wrong!", http.StatusBadRequest)
}
