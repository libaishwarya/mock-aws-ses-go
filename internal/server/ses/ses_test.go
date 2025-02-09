package ses_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/libaishwarya/mock-aws-ses-go/internal/server"
	"github.com/libaishwarya/mock-aws-ses-go/internal/server/ses"
	"github.com/libaishwarya/mock-aws-ses-go/internal/store/inmemory"
)

func TestSendEmail_Validation(t *testing.T) {
	router := startServer()

	tests := []struct {
		Name           string
		Body           map[string]any
		ExpectedStatus int
		ExpectedError  string
	}{
		{
			Name: "missing source",
			Body: map[string]any{
				"source":      "invalid email",
				"destination": "test@gmail.com",
				"message": map[string]any{
					"subject": map[string]any{"data": "test"},
					"body": map[string]any{
						"html": "test",
					},
				},
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedError:  "validation failed: source: Invalid email format",
		},
		{
			Name: "missing destination",
			Body: map[string]any{
				"source":      "test@gmail.com",
				"destination": "invalid email",
				"message": map[string]any{
					"subject": map[string]any{"data": "test"},
					"body": map[string]any{
						"html": "test",
					},
				},
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedError:  "validation failed: destination: Invalid email format",
		},
		{
			Name: "missing message",
			Body: map[string]any{
				"source":      "test@gmail.com",
				"destination": "test@gmail.com",
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedError:  "validation failed: Data: This field is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := sendEmailRequest(tt.Body)
			resp := server.ServerHTTP(router, req)

			server.AssertError(t, resp, tt.ExpectedStatus, tt.ExpectedError)
		})
	}
}

func TestSendEmail_Success(t *testing.T) {
	router := startServer()

	tests := []struct {
		Name      string
		Body      map[string]any
		MessageID string
	}{
		{
			Name: "createemail1",
			Body: map[string]any{
				"source":      "test@gmail.com",
				"destination": "test@gmail.com",
				"message": map[string]any{
					"subject": map[string]any{"data": "test"},
					"body": map[string]any{
						"html": "test",
					},
				},
			},
			MessageID: "1",
		},
		{
			Name: "createemail2r",
			Body: map[string]any{
				"source":      "test@gmail.com",
				"destination": "test@gmail.com",
				"message": map[string]any{
					"subject": map[string]any{"data": "test"},
					"body": map[string]any{
						"html": "test",
					},
				},
			},
			MessageID: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := sendEmailRequest(tt.Body)
			resp := server.ServerHTTP(router, req)

			server.Assert(t, resp, http.StatusOK, map[string]any{"messageId": tt.MessageID})
		})
	}
}

func TestSendRawEmail_Validation(t *testing.T) {
	router := startServer()

	tests := []struct {
		Name           string
		Body           map[string]any
		ExpectedStatus int
		ExpectedError  string
	}{
		{
			Name: "missing data",
			Body: map[string]any{
				"test": "test",
			},
			ExpectedStatus: http.StatusBadRequest,
			ExpectedError:  "validation failed: data: Invalid value",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := sendRawEmailRequest(tt.Body)
			resp := server.ServerHTTP(router, req)

			server.AssertError(t, resp, tt.ExpectedStatus, tt.ExpectedError)
		})
	}
}

func TestSendRawEmail_Success(t *testing.T) {
	router := startServer()

	tests := []struct {
		Name      string
		Body      map[string]any
		MessageID string
	}{
		{
			Name: "createrawemail1",
			Body: map[string]any{
				"data": "test",
			},
			MessageID: "1",
		},
		{
			Name: "createrawemail2",
			Body: map[string]any{
				"data": "test",
			},
			MessageID: "2",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			req := sendRawEmailRequest(tt.Body)
			resp := server.ServerHTTP(router, req)

			server.Assert(t, resp, http.StatusOK, map[string]any{"messageId": tt.MessageID})
		})
	}
}

func TestListIdentities_Success(t *testing.T) {
	router := startServer()

	t.Run("list_identities", func(t *testing.T) {
		req := sendlistIdentitiesRequest()
		resp := server.ServerHTTP(router, req)

		assert.Equal(t, http.StatusOK, resp.Code, "wrong status code")
	})
}

func sendEmailRequest(body map[string]any) *http.Request {
	data := &bytes.Buffer{}
	json.NewEncoder(data).Encode(body)
	req, _ := http.NewRequest("POST", "/v1/sendEmail", data)
	return req
}

func sendRawEmailRequest(body map[string]any) *http.Request {
	data := &bytes.Buffer{}
	json.NewEncoder(data).Encode(body)
	req, _ := http.NewRequest("POST", "/v1/sendRawEmail", data)
	return req
}

func sendlistIdentitiesRequest() *http.Request {
	req, _ := http.NewRequest("GET", "/v1/listIdentities", nil)
	return req
}

func startServer() *gin.Engine {
	r := gin.Default()
	inMemoryStore := inmemory.NewInMemoryStore()

	ses.AttachRoutes(r, inMemoryStore)

	return r

}
