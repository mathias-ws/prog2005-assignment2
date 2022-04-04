package handlers

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"assignment-2/test/stubs"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(m *testing.M) {
	database.InitDB("../../../auth.json")

	casesStub := httptest.NewServer(http.HandlerFunc(stubs.CasesHandler))
	defer casesStub.Close()

	webhookStub := httptest.NewServer(http.HandlerFunc(stubs.Webhook))
	defer webhookStub.Close()

	policyStub := httptest.NewServer(http.HandlerFunc(stubs.PolicyHandler))
	defer policyStub.Close()

	countryStub := httptest.NewServer(http.HandlerFunc(stubs.CountryHandler))
	defer countryStub.Close()

	constants.SetTestDB()
	constants.SetTestUrlCases(casesStub.URL)
	constants.SetTestUrlCases(webhookStub.URL)
	constants.SetTestUrlPolicy(policyStub.URL)
	constants.SetTestUrlCountry(countryStub.URL)

	m.Run()
}
