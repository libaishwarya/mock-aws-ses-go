package ses

import (
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/libaishwarya/mock-aws-ses-go/internal/serror"
)

type SESHandler struct {
}

func NewSESHandler() *SESHandler {
	return &SESHandler{}
}

func AttachRoutes(r *gin.Engine) {
	sesHandler := NewSESHandler()
	r.POST("/v1/sendEmail", sesHandler.SendEmail)
	r.POST("/v1/sendRawEmail", sesHandler.SendRawEmail)
	// r.GET("/v1/listIdentities", sesHandler.ListIdentities)
	// r.GET("/v1/getSendQuota", sesHandler.GetSendQuota)
	// r.GET("/v1/stats", sesHandler.GetStats)

}

func (s *SESHandler) SendEmail(c *gin.Context) {
	var req SendEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		serror.HandleBindError(c, err, req)
		return
	}

	// TODO: Implement store.

	// Simulate email sending
	c.JSON(http.StatusOK, SendEmailResponse{
		MessageId: "mock-message-id",
	})
}

func (s *SESHandler) SendRawEmail(c *gin.Context) {
	var req SendRawEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		serror.HandleBindError(c, err, req)
		return
	}

	// TODO: Implement store.

	// Simulate email sending
	c.JSON(http.StatusOK, SendEmailResponse{
		MessageId: "mock-message-id",
	})
}

type Body struct {
	HTML string `json:"html,omitempty"`
	Text string `json:"text,omitempty"`
}

type Subject struct {
	Data    string `json:"data" binding:"required"`
	Charset string `json:"charset" binding:"omitempty"`
}

type Message struct {
	Subject Subject `json:"subject" binding:"required"`
	Body    Body    `json:"body" binding:"required"`
}

type SendEmailRequest struct {
	Source      string  `json:"source" binding:"email,required"`
	Destination string  `json:"destination" binding:"email,required"`
	Message     Message `json:"message" binding:"required"`

	// TODO: Add additional options fields later.
	ReturnPath string `json:"_"`
}

type SendEmailResponse struct {
	MessageId string `json:"messageId"`
}

type SendRawEmailRequest struct {
	Data string `json:"data" binding:"base64,required"`

	// TODO: Add additional options fields later.
	ReturnPath string `json:"_"`
}

// Custom Base64 validator
func base64Validator(fl validator.FieldLevel) bool {
	_, err := base64.StdEncoding.DecodeString(fl.Field().String())
	return err == nil
}
