package ses_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/libaishwarya/mock-aws-ses-go/internal/server/ses"

	"github.com/libaishwarya/mock-aws-ses-go/internal/server"
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

			server.AssertError(t, resp, http.StatusBadRequest, tt.ExpectedError)
		})
	}
}

func sendEmailRequest(body map[string]any) *http.Request {
	data := &bytes.Buffer{}
	json.NewEncoder(data).Encode(body)
	req, _ := http.NewRequest("POST", "/v1/sendEmail", data)
	return req
}

func startServer() *gin.Engine {
	r := gin.Default()
	ses.AttachRoutes(r)

	return r

}
