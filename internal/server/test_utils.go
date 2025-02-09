package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func ServerHTTP(r *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	return w
}

func AssertError(t *testing.T, rr *httptest.ResponseRecorder, wantCode int, wantMessage string) {
	t.Helper()

	assert.Equal(t, wantCode, rr.Code, "wrong status code")

	var data map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	assert.Equal(t, wantMessage, data["error"], "wrong error message")
}

func Assert(t *testing.T, rr *httptest.ResponseRecorder, wantCode int, expectedOutput map[string]any) {
	t.Helper()

	assert.Equal(t, wantCode, rr.Code, "wrong status code")

	var data map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &data)
	if err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	assert.Equal(t, expectedOutput, data, "wrong response")
}
