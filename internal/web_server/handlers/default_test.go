package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDefaultHandlerInvalidMethod(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodPost, "/corona/v1/statuasdlfkjs", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DefaultHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusMethodNotAllowed, responseRecorder.Code)
}

func TestDefaultHandler(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/asfkø/statuasdlfkjs", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DefaultHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestDefaultHandlerNoPath(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DefaultHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}

func TestDefaultHandlerTypo(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/notification", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DefaultHandler)

	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
}
