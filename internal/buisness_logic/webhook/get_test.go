package webhook

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/structs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllRegistered(t *testing.T) {
	webhooks, err := GetAllRegistered()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(webhooks))
}

func TestGetOne(t *testing.T) {
	expected := structs.WebHookRegistration{
		Country:             "Sweden",
		Calls:               2,
		CallsAtRegistration: 1,
		Url:                 "https://funny.url.go.fast/very-nice/swe",
	}

	params := map[string]string{
		constants.UrlParameterWebhookId: "10911bd27492a5be7c1c772c8528f6f207f7da1b35c727669235f74c93e860e2",
	}

	webhook, err := GetOne(params)

	assert.Nil(t, err)
	assert.Equal(t, expected, webhook)
}
