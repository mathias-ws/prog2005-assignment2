package handlers

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCovidCasesHandlerInvalidMethod(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodPost, "/corona/v1/cases", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CovidCasesHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, responseRecorder.Code)
}

func TestCovidCasesHandler(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/cases?country=Norway", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CovidCasesHandler)

	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()

	errDel := database.DeleteDocument(constants.CovidCasesDBCollection, "Norway")
	errDelCount := database.DeleteDocument(constants.CounterDbCollection, "Norway")

	body, _ := io.ReadAll(result.Body)

	assert.Nil(t, errDel)
	assert.Nil(t, errDelCount)
	assert.Equal(t, http.StatusOK, result.StatusCode)
	assert.Equal(t, "{\n\t\"country\": \"Norway\",\n\t\"date\": \"2022-03-28\",\n\t\"confirmed\": 1399714,\n\t\"recovered\": 0,\n\t\"deaths\": 2339,\n\t\"growth_rate\": 0.0014853631627073677\n}\n", string(body))
}

func TestCovidCasesHandlerNoQuery(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/cases", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CovidCasesHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCovidCasesHandlerWrongQuery(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/cases?scope=2901.12.3", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CovidCasesHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestCovidCasesHandlerNoInternet(t *testing.T) {
	errDel := database.DeleteDocument(constants.CovidCasesDBCollection, "Norway")
	assert.Nil(t, errDel)

	constants.SetTestUrlCases("aslkfj")

	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/cases?country=norway", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CovidCasesHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadGateway, responseRecorder.Code)
}
