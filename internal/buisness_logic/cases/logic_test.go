package cases

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/internal/structs"
	"assignment-2/test/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	database.InitDB("../../../auth.json")
	defer database.CloseFirestore()
	countryStub := httptest.NewServer(http.HandlerFunc(stubs.CasesHandler))
	defer countryStub.Close()

	constants.SetTestDB()
	constants.SetTestUrlCases(countryStub.URL)

	m.Run()
}

func TestCreateGraphQlRequestNorway(t *testing.T) {
	jsonData, err := createGraphQlRequest("Norway")

	expected := "{\"query\":\"query {\\n    country(name: \\\"Norway\\\") {\\n        name\\n        " +
		"mostRecent {\\n            date(format: \\\"yyyy-MM-dd\\\")\\n            confirmed\\n            " +
		"recovered\\n            deaths\\n            growthRate\\n        }\\n    }\\n}\"}"

	assert.Equal(t, expected, string(jsonData))
	assert.Nil(t, err)
}

func TestCreateGraphQlRequestSweden(t *testing.T) {
	jsonData, err := createGraphQlRequest("Sweden")

	expected := "{\"query\":\"query {\\n    country(name: \\\"Sweden\\\") {\\n        name\\n        " +
		"mostRecent {\\n            date(format: \\\"yyyy-MM-dd\\\")\\n            confirmed\\n            " +
		"recovered\\n            deaths\\n            growthRate\\n        }\\n    }\\n}\"}"

	assert.Equal(t, expected, string(jsonData))
	assert.Nil(t, err)
}

func TestGetCovidCasesValid(t *testing.T) {
	urlParams := map[string]string{
		constants.URL_COUNTRY_NAME_PARAM: "Norway",
	}

	result, err := GetCovidCases(urlParams)

	expected := structs.CovidCasesOutput{
		Country:        "Norway",
		Date:           "2022-03-28",
		ConfirmedCases: 1399714,
		Recovered:      0,
		Deaths:         2339,
		GrowthRate:     0.0014853631627073677,
		TimeStamp:      time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC)}

	// To make sure that the caching is done
	time.Sleep(time.Millisecond * 500)
	errCache := database.DeleteDocument(constants.CovidCasesDBCollection, "Norway")
	errCount := database.DeleteDocument(constants.CounterDbCollection, "Norway")

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
	assert.Nil(t, errCache)
	assert.Nil(t, errCount)
}

func TestGetCovidCasesFromCache(t *testing.T) {
	urlParams := map[string]string{
		constants.URL_COUNTRY_NAME_PARAM: "Sweden",
	}

	result, err := GetCovidCases(urlParams)

	expected := structs.CovidCasesOutput{
		Country:        "Sweden",
		Date:           "2022-03-28",
		ConfirmedCases: 1399714,
		Recovered:      0,
		Deaths:         2339,
		GrowthRate:     0.0014853631627073677,
		TimeStamp:      time.Date(2022, time.March, 29, 17, 47, 44, 297000000, time.UTC)}

	// To make sure that the caching is done
	time.Sleep(time.Millisecond * 500)
	errCount := database.DeleteDocument(constants.CounterDbCollection, "Sweden")

	assert.Equal(t, expected, result)
	assert.Nil(t, err)
	assert.Nil(t, errCount)
}
