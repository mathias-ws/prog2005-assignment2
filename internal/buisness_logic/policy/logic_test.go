package policy

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/custom_errors"
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

	policyStub := httptest.NewServer(http.HandlerFunc(stubs.PolicyHandler))

	constants.SetTestUrlPolicy(policyStub.URL + "/")

	constants.SetTestDB()

	m.Run()
}

func TestCountPolicies(t *testing.T) {
	testSlice := []structs.Policy{
		{
			PolicyTypeCode:          "C1",
			PolicyValueDisplayField: "sdolak",
		},
		{
			PolicyTypeCode:          "C2",
			PolicyValueDisplayField: "sdolak",
		},
		{
			PolicyTypeCode:          "C3",
			PolicyValueDisplayField: "sdolak",
		},
	}

	assert.Equal(t, 3, countPolicies(testSlice))
}

func TestCountPoliciesNone(t *testing.T) {
	testSlice := []structs.Policy{
		{
			PolicyTypeCode:          "NONE",
			PolicyValueDisplayField: "sdolak",
		},
	}

	assert.Equal(t, 0, countPolicies(testSlice))
}

func TestCountPoliciesNil(t *testing.T) {
	assert.Equal(t, 0, countPolicies(nil))
}

func TestGetStringencyActual(t *testing.T) {
	testData := structs.Stringency{
		Stringency:       12.2,
		DateValue:        "2022-02-10",
		StringencyActual: 12.3,
		Confirmed:        100,
		Deaths:           2,
		CountryCode:      "nor",
	}

	assert.Equal(t, 12.3, getStringency(testData))
}

func TestGetStringency(t *testing.T) {
	testData := structs.Stringency{
		Stringency:       12.2,
		DateValue:        "2022-02-10",
		StringencyActual: 0,
		Confirmed:        100,
		Deaths:           2,
		CountryCode:      "nor",
	}

	assert.Equal(t, 12.2, getStringency(testData))
}

func TestGetStringencyNon(t *testing.T) {
	testData := structs.Stringency{
		Stringency:       0,
		DateValue:        "2022-02-10",
		StringencyActual: 0,
		Confirmed:        100,
		Deaths:           2,
		CountryCode:      "nor",
	}

	assert.Equal(t, float64(-1), getStringency(testData))
}

func TestBuildSearchUrl(t *testing.T) {
	params := map[string]string{
		constants.URL_COUNTRY_NAME_PARAM: "nor",
		constants.URL_SCOPE_PARAMETER:    "2022-02-10",
	}

	assert.Equal(t, constants.CovidTrackerBaseUrl+"/v2/stringency/actions/nor/2022-02-10",
		buildSearchUrl(params))
}

func TestFindPolicyInformation(t *testing.T) {
	params := map[string]string{
		constants.URL_COUNTRY_NAME_PARAM: "nor",
		constants.URL_SCOPE_PARAMETER:    "2022-10-20",
	}

	outputExpect := structs.PolicyOutput{
		CountryCode: "nor",
		Policy:      2,
		TimeStamp:   time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
		Stringency:  13.89,
		Scope:       "2022-10-20",
	}

	output, err := FindPolicyInformation(params)
	errDelCount := database.DeleteDocument(constants.CounterDbCollection, "Norway")
	errDel := database.DeleteDocument(constants.PolicyDBCollection, output.CountryCode+output.Scope)

	assert.Nil(t, err)
	assert.Equal(t, outputExpect, output)
	assert.Nil(t, errDel)
	assert.Nil(t, errDelCount)
}

func TestFindPolicyInformationCached(t *testing.T) {
	params := map[string]string{
		constants.URL_COUNTRY_NAME_PARAM: "swe",
		constants.URL_SCOPE_PARAMETER:    "2022-10-05",
	}

	outputExpect := structs.PolicyOutput{
		CountryCode: "swe",
		Policy:      2,
		TimeStamp:   time.Date(2022, time.July, 1, 10, 44, 33, 0, time.UTC),
		Stringency:  13.89,
		Scope:       "2022-10-05",
	}

	output, err := FindPolicyInformation(params)
	errDel := database.DeleteDocument(constants.CounterDbCollection, "Sweden")

	assert.Nil(t, err)
	assert.Equal(t, outputExpect, output)
	assert.Nil(t, errDel)
}

func TestFindPolicyInformationNoParams(t *testing.T) {
	params := map[string]string{
		constants.URL_COUNTRY_NAME_PARAM: "hjk",
		constants.URL_SCOPE_PARAMETER:    "2022-10-20",
	}

	constants.SetTestUrlPolicy("http://akljdsfh")

	_, err := FindPolicyInformation(params)

	assert.Equal(t, custom_errors.GetUnableToReachBackendApisError(), err)
}
