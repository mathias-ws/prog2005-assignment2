package json_parsing

import (
	"assignment-2/internal/structs"
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestDecodeRequest(t *testing.T) {
	body := map[string]interface{}{
		"url":     "https://my.webhook/askljfhao83hiofa",
		"country": "Norway",
		"calls":   3,
	}

	jsonData, _ := json.Marshal(body)
	request, _ := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonData))

	actualValue := structs.WebHookRegistration{}

	Decode(request, &actualValue)

	expected := structs.WebHookRegistration{
		Id:                  "",
		Url:                 "https://my.webhook/askljfhao83hiofa",
		Country:             "Norway",
		Calls:               3,
		CallsAtRegistration: 0,
	}

	assert.Equal(t, expected, actualValue)
}

func TestDecodeResponse(t *testing.T) {
	body := map[string]interface{}{
		"url":     "https://my.webhook/askljfhao83hiofa",
		"country": "Norway",
		"calls":   3,
	}

	jsonData, _ := json.Marshal(body)

	response := http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewBuffer(jsonData)),
	}

	actualValue := structs.WebHookRegistration{}

	Decode(&response, &actualValue)

	expected := structs.WebHookRegistration{
		Id:                  "",
		Url:                 "https://my.webhook/askljfhao83hiofa",
		Country:             "Norway",
		Calls:               3,
		CallsAtRegistration: 0,
	}

	assert.Equal(t, expected, actualValue)
}

func TestDecodeNilInput(t *testing.T) {
	actualValue := structs.WebHookRegistration{}

	Decode(nil, &actualValue)

	assert.Equal(t, structs.WebHookRegistration{}, structs.WebHookRegistration{})
}
