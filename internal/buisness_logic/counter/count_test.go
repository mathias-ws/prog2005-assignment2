package counter

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
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

	countryStub := httptest.NewServer(http.HandlerFunc(stubs.CountryHandler))

	constants.SetTestUrlCountry(countryStub.URL)

	constants.SetTestDB()

	m.Run()
}

func TestGetCachedCca3(t *testing.T) {
	currentCount, errGet := GetNumberOfTimes("swe")

	assert.Equal(t, 1, currentCount)
	assert.Nil(t, errGet)
}

func TestCountCreateNew(t *testing.T) {
	errCountUp := CountUp("Norway")

	currentCount, errGet := GetNumberOfTimes("Norway")

	// To make sure that the caching is done
	time.Sleep(time.Second)
	errDel := database.DeleteDocument(constants.CounterDbCollection, "Norway")

	assert.Equal(t, 1, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
}

func TestCountCreateNewCca3(t *testing.T) {
	errCountUp := CountUp("nor")

	currentCount, errGet := GetNumberOfTimes("Norway")

	// To make sure that the caching is done
	time.Sleep(time.Second)
	errDel := database.DeleteDocument(constants.CounterDbCollection, "Norway")

	assert.Equal(t, 1, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
}

func TestCountCached(t *testing.T) {
	errCountUp := CountUp("Sweden")

	// To make sure that the caching is done
	time.Sleep(time.Second)

	currentCount, errGet := GetNumberOfTimes("Sweden")

	// Replaces it in the db so that it always is 1.
	errDel := database.DeleteDocument(constants.CounterDbCollection, "Sweden")
	errReplace := CountUp("Sweden")

	assert.Equal(t, 2, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
	assert.Nil(t, errReplace)
}

func TestCountCachedCca3(t *testing.T) {
	errCountUp := CountUp("swe")

	currentCount, errGet := GetNumberOfTimes("Sweden")

	// To make sure that the caching is done
	time.Sleep(time.Second)

	// Replaces it in the db so that it always is 1.
	errDel := database.DeleteDocument(constants.CounterDbCollection, "Sweden")
	errReplace := CountUp("Sweden")

	assert.Equal(t, 2, currentCount)
	assert.Nil(t, errCountUp)
	assert.Nil(t, errGet)
	assert.Nil(t, errDel)
	assert.Nil(t, errReplace)
}
