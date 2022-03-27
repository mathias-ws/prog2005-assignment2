package urlutil

import (
	"assignment-2/internal/custom_errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

func TestGetWebhookParameterNegative(t *testing.T) {
	testUrl1 := url.URL{RawQuery: "country=Norway"}
	parameters1, err1 := GetWebhookParameter(&testUrl1)

	assert.Empty(t, parameters1)
	assert.Nil(t, err1)

	testUrl2 := url.URL{RawQuery: "id=Norway"}
	parameters2, err2 := GetWebhookParameter(&testUrl2)

	assert.Empty(t, parameters2)
	assert.Equal(t, custom_errors.GetParameterError(), err2)

	testUrl3 := url.URL{RawQuery: "id=a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470&country=Norway"}
	parameters3, err3 := GetWebhookParameter(&testUrl3)

	assert.NotContains(t, parameters3["country"], "Norway")
	assert.Contains(t, "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470", parameters3["id"])
	assert.Nil(t, err3)

	testUrl4 := url.URL{RawQuery: ""}
	parameters4, err4 := GetWebhookParameter(&testUrl4)

	assert.Empty(t, parameters4)
	assert.Nil(t, err4)
}

func TestGetWebhookParameterPositive(t *testing.T) {
	testUrl1 := url.URL{RawQuery: "id=a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470"}
	parameters1, err1 := GetWebhookParameter(&testUrl1)

	assert.Contains(t, "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470", parameters1["id"])
	assert.Nil(t, err1)
}
