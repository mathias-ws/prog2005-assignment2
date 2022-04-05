package status

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
	defer countryStub.Close()

	casesStub := httptest.NewServer(http.HandlerFunc(stubs.CasesHandler))
	defer casesStub.Close()

	policyStub := httptest.NewServer(http.HandlerFunc(stubs.PolicyHandler))
	defer policyStub.Close()

	constants.SetTestDB()
	constants.SetTestUrlCases(casesStub.URL)
	constants.SetTestUrlCountry(countryStub.URL)
	constants.SetTestUrlPolicy(policyStub.URL)

	m.Run()
}

func TestGetUptime(t *testing.T) {
	assert.Equal(t, "0h 0m 0s", getUptime())
}

func TestGetNumberOfWebhooks(t *testing.T) {
	time.Sleep(time.Second * 2)
	webhooks, err := getNumberOfWebhooks()

	allWebhooks, errGetWebhooks := database.GetAllWebhooks(constants.WebhookDbCollection, "")
	assert.Nil(t, errGetWebhooks)

	assert.Nil(t, err)
	assert.Equal(t, len(allWebhooks), webhooks)
}

func TestGetStatusInfo(t *testing.T) {
	webhooks, errGetWebhooks := database.GetAllWebhooks(constants.WebhookDbCollection, "")
	assert.Nil(t, errGetWebhooks)

	expect := status{
		CountryApiStatusCode: 200,
		CasesApiStatusCode:   200,
		PolicyApiStatusCode:  200,
		NumberOfWebhooks:     len(webhooks),
		Version:              constants.PROGRAM_VERSION,
		Uptime:               getUptime(),
	}

	actual := GetStatusInfo()

	assert.Equal(t, expect, actual)
}

func TestGetStatusInfoNoInternet(t *testing.T) {
	time.Sleep(time.Second * 2)
	constants.SetTestUrlCases("http://oiasdnf")
	constants.SetTestUrlCountry("http://oiasdnf")
	constants.SetTestUrlPolicy("http://oiasdnf")

	webhooks, errGetWebhooks := database.GetAllWebhooks(constants.WebhookDbCollection, "")
	assert.Nil(t, errGetWebhooks)

	expect := status{
		CountryApiStatusCode: 502,
		CasesApiStatusCode:   502,
		PolicyApiStatusCode:  502,
		NumberOfWebhooks:     len(webhooks),
		Version:              constants.PROGRAM_VERSION,
		Uptime:               getUptime(),
	}

	actual := GetStatusInfo()

	assert.Equal(t, expect, actual)
}
