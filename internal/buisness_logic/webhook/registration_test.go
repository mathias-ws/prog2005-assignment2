package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestRegister(t *testing.T) {
	body := map[string]interface{}{
		"url":     "https://my.webhook/askljfhao83hiofa",
		"country": "Norway",
		"calls":   3,
	}

	jsonData, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonData))

	resultingStruct := structs.WebHookRegistration{
		Id:                  "",
		Url:                 "https://my.webhook/askljfhao83hiofa",
		Country:             "Norway",
		Calls:               3,
		CallsAtRegistration: 1}

	hash, err := Register(request)

	testFetch := structs.WebHookRegistration{}
	database.GetDocument(constants.WebhookDbCollection,
		"7fbfd72f256dd5ffee45dc1df669ef9ebb76d539f1c4b775bc815d9729782a77", &testFetch)

	//Removes the testing object.
	errDel := database.DeleteDocument(constants.WebhookDbCollection,
		"7fbfd72f256dd5ffee45dc1df669ef9ebb76d539f1c4b775bc815d9729782a77")

	assert.Nil(t, err)
	assert.Equal(t, "7fbfd72f256dd5ffee45dc1df669ef9ebb76d539f1c4b775bc815d9729782a77", hash)
	assert.Equal(t, resultingStruct, testFetch)
	assert.Nil(t, errDel)
}

func TestRegisterWrongBody(t *testing.T) {
	body := map[string]interface{}{}

	jsonData, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonData))

	_, err := Register(request)

	assert.Equal(t, custom_errors.GetFailedToDecode(), err)
}
