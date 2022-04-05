package handlers

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusHandlerInvalidMethod(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodPost, "/corona/v1/status", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, responseRecorder.Code)
}

func TestStatusHandler(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/status", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(StatusHandler)

	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)

	actualValue := string(body)[:len(string(body))-16]
	webhooks, errGetWebhooks := database.GetAllWebhooks(constants.WebhookDbCollection, "")
	assert.Nil(t, errGetWebhooks)

	assert.Equal(t, "{\n\t\"cases_api_status\": 200,\n\t\"policy_api_status\": 200,\n\t\"country_api_status\": 200,\n\t\"number_of_webhooks\": "+fmt.Sprint(len(webhooks))+",\n\t\"version\": \"v1\",\n\t\"uptime", actualValue)
}
