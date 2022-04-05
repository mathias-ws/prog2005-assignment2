package handlers

import (
	"assignment-2/internal/constants"
	"assignment-2/internal/database"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestPolicyHandlerInvalidMethod(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodPost, "/corona/v1/policy", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PolicyHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, responseRecorder.Code)
}

func TestPolicyHandler(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/policy?country=nor&scope=2022-03-04", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PolicyHandler)

	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	body, _ := io.ReadAll(result.Body)

	time.Sleep(time.Millisecond * 200)
	errDel := database.DeleteDocument(constants.PolicyDBCollection, "nor2022-03-04")
	errDelCount := database.DeleteDocument(constants.CounterDbCollection, "Norway")
	assert.Nil(t, errDel)
	assert.Nil(t, errDelCount)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "{\n\t\"country_code\": \"nor\",\n\t\"date_value\": \"2022-03-04\",\n\t\"stringency\": 13.89,\n\t\"policies\": 2\n}\n", string(body))
}

func TestPolicyHandlerWithoutScope(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/policy?country=nor", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PolicyHandler)

	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	body, _ := io.ReadAll(result.Body)

	currentDate := time.Now().Format("2006-01-02")
	time.Sleep(time.Millisecond * 200)

	errDel := database.DeleteDocument(constants.PolicyDBCollection, "nor"+currentDate)
	errDelCount := database.DeleteDocument(constants.CounterDbCollection, "Norway")
	assert.Nil(t, errDel)
	assert.Nil(t, errDelCount)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, fmt.Sprintf("{\n\t\"country_code\": \"nor\",\n\t\"date_value\": \"%v\",\n\t\"stringency\": 13.89,\n\t\"policies\": 2\n}\n", currentDate), string(body))
}

func TestPolicyHandlerNoParams(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/policy", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PolicyHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestPolicyHandlerWrongParams(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/policy?foaspj=alskfn", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PolicyHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusBadRequest, responseRecorder.Code)
}

func TestPolicyHandlerNoAccessToBackendApi(t *testing.T) {
	constants.SetTestUrlPolicy("http://aslkfh")

	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/policy?country=nor&scope=2022-03-04", nil)
	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(PolicyHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusInternalServerError, responseRecorder.Code)
}
