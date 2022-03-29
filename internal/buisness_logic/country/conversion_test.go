package country

import (
	"assignment-2/internal/database"
	"assignment-2/test/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	database.InitDB()
	countryStub := httptest.NewServer(http.HandlerFunc(stubs.CountryHandler))
	defer countryStub.Close()

	SetTestUrlCountry(countryStub.URL)

	m.Run()
}

func TestGetCountryNameFromCca3(t *testing.T) {
	countryName, err := GetCountryNameFromCca3("nor")

	assert.Equal(t, "Norway", countryName)
	assert.Nil(t, err)
}
