package handlers

import (
	"assignment-2/internal/buisness_logic/webhook"
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/web_server/json"
	"assignment-2/internal/web_server/urlutil"
	"net/http"
)

// NotificationHandler checks for the http method and handles the error appropriately.
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleGetWebhooks(w, r)
	case http.MethodPost:
		handleRegistration(w, r)
	case http.MethodDelete:
		handleDeleteRequest(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	webhookId, err := webhook.Register(r)
	if err != nil {
		if err.Error() == custom_errors.GetMissingJsonFieldsError().Error() {
			custom_errors.HttpWrongJsonInfo(w)
			return
		}
		custom_errors.HttpUnknownServerError(w)
		return
	}

	err = json.Encode(w, map[string]string{"webhook_id": webhookId})

	if err != nil {
		return
	}
}

// handleGetWebhooks handles the get request for the notification endpoint. If the id parameter is present, it gets
// one webhook if no parameter is present it gets all webhooks.
func handleGetWebhooks(w http.ResponseWriter, r *http.Request) {
	urlParams, err := urlutil.GetWebhookParameter(r.URL)

	if err != nil {
		custom_errors.HttpSearchParameters(w)
		return
	} else if !(urlParams[constants.UrlParameterWebhookId] == "") {
		// Return
		obtainedWebhook, err := webhook.GetOne(urlParams)
		if err != nil {
			return
		}
		err = json.Encode(w, obtainedWebhook)
		if err != nil {
			return
		}
	} else {
		webhooks, _ := webhook.GetAllRegistered()
		err := json.Encode(w, webhooks)
		if err != nil {
			return
		}
	}
}

func handleDeleteRequest(w http.ResponseWriter, r *http.Request) {
	urlParams, err := urlutil.GetWebhookParameter(r.URL)

	if err != nil {
		custom_errors.HttpSearchParameters(w)
		return
	} else if len(urlParams) == 0 {
		custom_errors.HttpSearchParameters(w)
		return
	}

	message, err := webhook.Delete(urlParams)

	if err != nil {
		return
	}

	errEncode := json.Encode(w, message)
	if errEncode != nil {
		return
	}
}
