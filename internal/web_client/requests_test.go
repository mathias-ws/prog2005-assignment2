package web_client

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
	"assignment-2/test/stubs"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	casesStub := httptest.NewServer(http.HandlerFunc(stubs.CasesHandler))
	defer casesStub.Close()

	constants.SetTestUrlCases(casesStub.URL)

	m.Run()
}

func TestGetStatusCode(t *testing.T) {
	code, err := GetStatusCode(constants.CovidCasesBaseUrl)

	assert.Nil(t, err)
	assert.Equal(t, 200, code)
}

func TestGetStatusCodeWrongUrl(t *testing.T) {
	code, err := GetStatusCode("http://alsfkhj")

	assert.Equal(t, custom_errors.GetUnableToReachBackendApisError(), err)
	assert.Equal(t, 0, code)
}

func TestCreateRequestNoBody(t *testing.T) {
	request, err := createRequest("test.com", http.MethodGet, nil)

	assert.Nil(t, err)
	assert.Equal(t, http.MethodGet, request.Method)
	assert.Equal(t, "test.com", request.URL.Path)
	assert.Equal(t, http.NoBody, request.Body)
}

func TestCreateRequestBody(t *testing.T) {
	body := map[string]interface{}{
		"url":     "https://my.webhook/askljfhao83hiofa",
		"country": "Norway",
		"calls":   3,
	}

	jsonData, _ := json.Marshal(body)

	request, err := createRequest("test.com", http.MethodGet, jsonData)

	readBody, _ := io.ReadAll(request.Body)

	assert.Nil(t, err)
	assert.Equal(t, http.MethodGet, request.Method)
	assert.Equal(t, "test.com", request.URL.Path)
	assert.Equal(t, "{\"calls\":3,\"country\":\"Norway\",\"url\":\"https://my.webhook/askljfhao83hiofa\"}", string(readBody))
}
