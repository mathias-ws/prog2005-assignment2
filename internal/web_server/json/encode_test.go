package json

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	Encode(w, "hei")
}

func TestEncode(t *testing.T) {
	req, errReq := http.NewRequest(http.MethodGet, "/corona/v1/cases?scope=2901.12.3", nil)

	assert.Nil(t, errReq)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(testHandler)

	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	body, _ := io.ReadAll(result.Body)
	assert.Equal(t, "\"hei\"\n", string(body))
}
