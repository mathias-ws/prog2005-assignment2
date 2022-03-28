package urlutil

import (
	"assignment-2/internal/custom_errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestGetWebhookParameterWrongParameter(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=Norway"}
	parameters, err := GetWebhookParameter(&testUrl)

	assert.Empty(t, parameters)
	assert.Nil(t, err)
}

func TestGetWebhookParameterWrongData(t *testing.T) {
	testUrl := url.URL{RawQuery: "id=Norway"}
	parameters, err := GetWebhookParameter(&testUrl)

	assert.Empty(t, parameters)
	assert.Equal(t, custom_errors.GetParameterError(), err)
}

func TestGetWebhookParameterExtraParam(t *testing.T) {
	testUrl := url.URL{RawQuery: "id=a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470&country=Norway"}
	parameters, err := GetWebhookParameter(&testUrl)

	assert.NotContains(t, parameters["country"], "Norway")
	assert.Contains(t, "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470", parameters["id"])
	assert.Nil(t, err)
}

func TestGetWebhookParameterValid(t *testing.T) {
	testUrl := url.URL{RawQuery: "id=a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470"}
	parameters, err := GetWebhookParameter(&testUrl)

	assert.Contains(t, "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470", parameters["id"])
	assert.Nil(t, err)
}

func TestGetWebhookParameterNoQuery(t *testing.T) {
	testUrl := url.URL{RawQuery: ""}
	parameters, err := GetWebhookParameter(&testUrl)

	assert.Empty(t, parameters)
	assert.Nil(t, err)
}

func TestGetWebhookParameterPositive(t *testing.T) {
	testUrl1 := url.URL{RawQuery: "id=a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470"}
	parameters1, err1 := GetWebhookParameter(&testUrl1)

	assert.Contains(t, "a2c4d3e5e592a8b75da5d9b3a27ad846a40338ffe2ed00771179e63991619470", parameters1["id"])
	assert.Nil(t, err1)
}

func TestGetUrlParametersPolicyNoParams(t *testing.T) {
	testUrl := url.URL{RawQuery: ""}
	parameters, err := GetUrlParametersPolicy(&testUrl)

	assert.Empty(t, parameters)
	assert.Equal(t, custom_errors.GetParameterError(), err)
}

func TestGetUrlParametersPolicyNoCountry(t *testing.T) {
	testUrl := url.URL{RawQuery: "scope=2022-02-29"}
	parameters, err := GetUrlParametersPolicy(&testUrl)

	assert.Empty(t, parameters)
	assert.Equal(t, custom_errors.GetParameterError(), err)
}

func TestGetUrlParametersPolicyInvalidDateFormat(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=Norway&scope=2022-2-10"}
	parameters, err := GetUrlParametersPolicy(&testUrl)

	assert.Equal(t, map[string]string(nil), parameters)
	assert.Equal(t, custom_errors.GetParameterError(), err)
}

func TestGetUrlParametersPolicyValid(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=Norway&scope=2022-02-10"}
	parameters, err := GetUrlParametersPolicy(&testUrl)

	assert.Equal(t, map[string]string{
		"country": "Norway",
		"scope":   "2022-02-10",
	}, parameters)

	assert.Nil(t, err)
}

func TestGetUrlParametersPolicyWithoutScope(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=Norway"}
	parameters, err := GetUrlParametersPolicy(&testUrl)

	assert.Equal(t, map[string]string{
		"country": "Norway",
		"scope":   time.Now().Format("2006-01-02"),
	}, parameters)

	assert.Nil(t, err)
}

func TestGetUrlParametersCasesCheckUppercase(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=norway"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string{"country": "Norway"}, parameters)
	assert.Nil(t, err)
}

func TestGetUrlParametersCasesus(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=us"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string{"country": "US"}, parameters)
	assert.Nil(t, err)
}

func TestGetUrlParametersCasesUs(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=Us"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string{"country": "US"}, parameters)
	assert.Nil(t, err)
}

func TestGetUrlParametersCasesUsa(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=Usa"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string{"country": "US"}, parameters)
	assert.Nil(t, err)
}

func TestGetUrlParametersCasesUSA(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=USA"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string{"country": "US"}, parameters)
	assert.Nil(t, err)
}

func TestGetUrlParametersCasesCca3Code(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=nor"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string{"country": "Norway"}, parameters)
	assert.Nil(t, err)
}

func TestGetUrlParametersCasesNumber(t *testing.T) {
	testUrl := url.URL{RawQuery: "country=nor2"}
	parameters, err := GetUrlParametersCases(&testUrl)

	assert.Equal(t, map[string]string(nil), parameters)
	assert.Equal(t, custom_errors.GetParameterError(), err)
}
