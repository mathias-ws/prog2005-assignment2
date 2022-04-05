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
		Url:                 "https://my.webhook/askljfhao83hiofa",
		Country:             "Norway",
		Calls:               3,
		CallsAtRegistration: 0}

	errDelCount := database.DeleteDocument(constants.CounterDbCollection, "Norway")
	assert.Nil(t, errDelCount)

	hash, err := Register(request)

	testFetch := structs.WebHookRegistration{}
	database.GetDocument(constants.WebhookDbCollection,
		"b9156dda4114476e0a499bf2182416381491011f262af679eca6520e582910a5", &testFetch)

	//Removes the testing object.
	errDel := database.DeleteDocument(constants.WebhookDbCollection,
		"b9156dda4114476e0a499bf2182416381491011f262af679eca6520e582910a5")

	assert.Nil(t, err)
	assert.Equal(t, "b9156dda4114476e0a499bf2182416381491011f262af679eca6520e582910a5", hash)
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
