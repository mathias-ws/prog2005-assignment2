package counter

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

	constants.SetTestUrlCountry(countryStub.URL)

	constants.SetTestDB()

	m.Run()
}

func TestGetCachedCca3(t *testing.T) {
	currentCount, errGet := GetNumberOfTimes("fra")

	errDel := database.DeleteDocument(constants.CountryDbCollection, "France")

	assert.Equal(t, 1, currentCount)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
}

func TestCountCreateNew(t *testing.T) {
	errCountUp := CountUp("Ireland")

	currentCount, errGet := GetNumberOfTimes("Ireland")

	errDel := database.DeleteDocument(constants.CounterDbCollection, "Ireland")

	assert.Equal(t, 1, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
}

func TestCountCreateNewCca3(t *testing.T) {
	errCountUp := CountUp("lva")

	currentCount, errGet := GetNumberOfTimes("Latvia")

	errDelCountry := database.DeleteDocument(constants.CountryDbCollection, "Latvia")
	errDel := database.DeleteDocument(constants.CounterDbCollection, "Latvia")

	assert.Equal(t, 1, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
	assert.Nil(t, errDelCountry)
}

func TestCountCached(t *testing.T) {
	errCountUp := CountUp("Netherlands")

	currentCount, errGet := GetNumberOfTimes("Netherlands")

	errDel := database.DeleteDocument(constants.CounterDbCollection, "Netherlands")
	errReplace := CountUp("Netherlands")

	assert.Equal(t, 2, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
	assert.Nil(t, errReplace)
}

func TestCountCachedCca3(t *testing.T) {
	errCountUp := CountUp("nld")

	currentCount, errGet := GetNumberOfTimes("Netherlands")

	errDelCountry := database.DeleteDocument(constants.CountryDbCollection, "Netherlands")

	errDel := database.DeleteDocument(constants.CounterDbCollection, "Netherlands")
	errReplace := CountUp("Netherlands")

	assert.Equal(t, 2, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
	assert.Nil(t, errReplace)
	assert.Nil(t, errDelCountry)
}
