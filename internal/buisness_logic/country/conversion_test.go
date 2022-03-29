package country

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/test/stubs"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	database.InitDB("../../../auth.json")
	defer database.CloseFirestore()

	countryStub := httptest.NewServer(http.HandlerFunc(stubs.CountryHandler))
	defer countryStub.Close()

	constants.SetTestDB()
	constants.SetTestUrlCountry(countryStub.URL)

	m.Run()
}

func TestGetCountryNameFromCca3(t *testing.T) {
	countryName, err := GetCountryNameFromCca3("nor")

	errDel := database.DeleteDocument(constants.CountryDbCollection, "nor")

	assert.Equal(t, "Norway", countryName)
	assert.Nil(t, err)
	assert.Nil(t, errDel)
}

func TestGetCountryNameFromCca3FromCache(t *testing.T) {
	countryName, err := GetCountryNameFromCca3("swe")

	assert.Equal(t, "Sweden", countryName)
	assert.Nil(t, err)
}
