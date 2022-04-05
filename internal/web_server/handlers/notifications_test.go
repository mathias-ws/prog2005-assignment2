package handlers

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotificationHandlerInvalidMethod(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodPut, "/corona/v1/notification", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, responseRecorder.Code)
}

func TestNotificationHandlerGetAllWebhooks(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/notifications", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	decoder := json.NewDecoder(responseRecorder.Body)
	var webhooks []structs.WebHookRegistration

	errDecode := decoder.Decode(&webhooks)
	assert.Nil(t, errDecode)

	expectedList := []structs.WebHookRegistration{
		{
			Id:                  "10911bd27492a5be7c1c772c8528f6f207f7da1b35c727669235f74c93e860e2",
			Url:                 "https://funny.url.go.fast/very-nice/swe",
			Country:             "Sweden",
			Calls:               2,
			CallsAtRegistration: 0,
		},
		{
			Id:                  "1ba267147386413b6ca89b2963d9fefc1aa756ca591a10aa9bc9564f1eaec532",
			Url:                 "https://funny.url.go.fast/very-nice/fin",
			Country:             "Finland",
			Calls:               5,
			CallsAtRegistration: 0,
		},
		{
			Id:                  "7fbfd72f256dd5ffee45dc1df669ef9ebb76d539f1c4b775bc815d9729782a77",
			Url:                 "https://my.webhook/askljfhao83hiofa",
			Country:             "Norway",
			Calls:               3,
			CallsAtRegistration: 0,
		},
		{
			Id:                  "eb40c0dd7ba0baf262b51fe4a3e98b3518aa2259549922b8e20976fd3f4685f4",
			Url:                 "https://funny.url.go.fast/very-nice/den",
			Country:             "Denmark",
			Calls:               5,
			CallsAtRegistration: 0,
		},
	}

	assert.Equal(t, expectedList, webhooks)
}

func TestNotificationHandlerGetOneWebhook(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet,
		"/corona/v1/notifications?id=10911bd27492a5be7c1c772c8528f6f207f7da1b35c727669235f74c93e860e2", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	decoder := json.NewDecoder(responseRecorder.Body)
	var webhooks structs.WebHookRegistration

	errDecode := decoder.Decode(&webhooks)
	assert.Nil(t, errDecode)

	expected := structs.WebHookRegistration{
		Url:                 "https://funny.url.go.fast/very-nice/swe",
		Country:             "Sweden",
		Calls:               2,
		CallsAtRegistration: 0,
	}

	assert.Equal(t, expected, webhooks)
}

func TestNotificationHandlerGetInvalidParameters(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/notification?ko=kjlf", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	decoder := json.NewDecoder(responseRecorder.Body)
	var webhooks []structs.WebHookRegistration

	errDecode := decoder.Decode(&webhooks)
	assert.Nil(t, errDecode)

	expectedList := []structs.WebHookRegistration{
		{
			Id:                  "10911bd27492a5be7c1c772c8528f6f207f7da1b35c727669235f74c93e860e2",
			Url:                 "https://funny.url.go.fast/very-nice/swe",
			Country:             "Sweden",
			Calls:               2,
			CallsAtRegistration: 0,
		},
		{
			Id:                  "1ba267147386413b6ca89b2963d9fefc1aa756ca591a10aa9bc9564f1eaec532",
			Url:                 "https://funny.url.go.fast/very-nice/fin",
			Country:             "Finland",
			Calls:               5,
			CallsAtRegistration: 0,
		},
		{
			Id:                  "7fbfd72f256dd5ffee45dc1df669ef9ebb76d539f1c4b775bc815d9729782a77",
			Url:                 "https://my.webhook/askljfhao83hiofa",
			Country:             "Norway",
			Calls:               3,
			CallsAtRegistration: 0,
		},
		{
			Id:                  "eb40c0dd7ba0baf262b51fe4a3e98b3518aa2259549922b8e20976fd3f4685f4",
			Url:                 "https://funny.url.go.fast/very-nice/den",
			Country:             "Denmark",
			Calls:               5,
			CallsAtRegistration: 0,
		},
	}

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, expectedList, webhooks)
}

func TestNotificationHandlerGetInvalidID(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet,
		"/corona/v1/notifications?id=8f6f207f7da1b35c727669235f74c93e860e2", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestAddNewWebhook(t *testing.T) {
	testPost := map[string]interface{}{
		"url":     "https://webhook.site/bf1c45f9-7ca4-4867-a17b-d7a10a49cea6",
		"country": "Sweden",
		"calls":   2,
	}

	jsonData, err := json.Marshal(testPost)

	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return
	}

	req, errReq := http.NewRequest(http.MethodPost, "/corona/v1/notifications", bytes.NewBuffer(jsonData))

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	errDel := database.DeleteDocument(constants.WebhookDbCollection, "64dc957d991f64cbc8b8534ef07b0e9f9a730b31c693630256106eaa3ccdf4cc")

	result := responseRecorder.Result()
	body, _ := io.ReadAll(result.Body)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Contains(t, string(body), "64dc957d991f64cbc8b8534ef07b0e9f9a730b31c693630256106eaa3ccdf4cc")
	assert.Nil(t, errDel)
}

func TestAddNewWebhookWrongBody(t *testing.T) {
	testPost := map[string]interface{}{
		"country": "Sweden",
	}

	jsonData, err := json.Marshal(testPost)

	if err != nil {
		log.Printf("Error marshalling data: %v", err)
		return
	}

	req, errReq := http.NewRequest(http.MethodPost, "/corona/v1/notifications", bytes.NewBuffer(jsonData))

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(NotificationHandler)

	handler.ServeHTTP(responseRecorder, req)

	errDel := database.DeleteDocument(constants.WebhookDbCollection, "64dc957d991f64cbc8b8534ef07b0e9f9a730b31c693630256106eaa3ccdf4cc")

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
	assert.Nil(t, errDel)
}
