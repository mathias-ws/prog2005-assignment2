package handlers

import (
	"assignment-2/internal/buisness_logic/webhook"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/web_server/json"
	"fmt"
	"net/http"
)

// NotificationHandler checks for the http method and handles the error appropriately.
func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Println("Get")
	case http.MethodPost:
		handleRegistration(w, r)
	default:
		// Returns method not supported for unsupported rest methods.
		custom_errors.HttpUnsupportedMethod(w)
	}
}

func handleRegistration(w http.ResponseWriter, r *http.Request) {
	webhookId, err := webhook.Register(r)
	if err != nil {
		return
	}

	err = json.Encode(w, map[string]string{"webhook_id": webhookId})

	if err != nil {
		return
	}
}