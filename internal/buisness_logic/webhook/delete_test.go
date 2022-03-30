package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDelete(t *testing.T) {
	testStruct := structs.WebHookRegistration{
		Country:             "Norway",
		Calls:               2,
		CallsAtRegistration: 0,
		Url:                 "https://funny.url.go.fast/very-nice",
	}

	errWrite := database.WriteDocument(constants.WebhookDbCollection, fmt.Sprintf("%v", testStruct), testStruct)

	params := map[string]string{
		constants.UrlParameterWebhookId: "0bc71804a34c3fd2ff6cc12f4423ac24562c664913f8e6a08a03b8d61c4f0e97",
	}

	errDel := Delete(params)

	testFetch := structs.WebHookRegistration{}
	database.GetDocument(constants.WebhookDbCollection,
		"0bc71804a34c3fd2ff6cc12f4423ac24562c664913f8e6a08a03b8d61c4f0e97", &testFetch)

	assert.Nil(t, errDel)
	assert.Nil(t, errWrite)

	assert.Equal(t, structs.WebHookRegistration{}, testFetch)
}
