package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/hashing"
	"assignment-2/internal/structs"
	"assignment-2/test/stubs"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	webhookClient := httptest.NewServer(http.HandlerFunc(stubs.Webhook))
	constants.SetTestUrlWebhookClient(webhookClient.URL)

	testStruct := structs.WebHookRegistration{
		Country:             "Germany",
		Calls:               1,
		CallsAtRegistration: 0,
		Url:                 constants.WebhookClientUrl,
	}

	//Create webhook
	errWrite := database.WriteDocument(constants.WebhookDbCollection, fmt.Sprintf("%v", testStruct), testStruct)

	hashedDocName, _ := hashing.HashString(fmt.Sprintf("%v", testStruct))
	expectedValue := structs.WebHookPost{
		Id:      hashedDocName,
		Country: "Germany",
		Calls:   1,
	}

	// Count up so that it gets triggered.
	database.IncrementCounter(constants.CounterDbCollection, "Germany")

	Check("Germany")

	// Get the value from the db.
	actualValue := structs.WebHookPost{}
	// A wait to fix concurrency issues.
	time.Sleep(time.Second)
	database.GetDocument(constants.WebhookTestCheckVerificationCollection,
		"test", &actualValue)

	// Delete temp structures
	errDel := database.DeleteDocument(constants.WebhookDbCollection,
		fmt.Sprintf("%v", testStruct))
	errDelCount := database.DeleteDocument(constants.CounterDbCollection,
		"Germany")
	errDelClient := database.DeleteDocument(constants.WebhookTestCheckVerificationCollection,
		"test")

	assert.Nil(t, errWrite)
	assert.Nil(t, errDel)
	assert.Nil(t, errDelCount)
	assert.Nil(t, errDelClient)
	assert.Equal(t, expectedValue, actualValue)
}

func Prep() error {
	webhookClient := httptest.NewServer(http.HandlerFunc(stubs.Webhook))
	constants.SetTestUrlWebhookClient(webhookClient.URL)

	testStruct := structs.WebHookRegistration{
		Country:             "Germany",
		Calls:               2,
		CallsAtRegistration: 1,
		Url:                 constants.WebhookClientUrl,
	}

	errWrite := database.WriteDocument(constants.WebhookDbCollection, fmt.Sprintf("%v", testStruct), testStruct)

	if errWrite != nil {
		return errWrite
	}
	return nil
}
