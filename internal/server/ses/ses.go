package ses

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/gin-gonic/gin"
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
}

func (s *SESHandler) SendEmail(c *gin.Context) {
	var req SendEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		serror.HandleBindError(c, err, req)
		return
	}

	// Simulate email sending
	c.JSON(http.StatusOK, ses.SendEmailOutput{
		MessageId: aws.String("mock-message-id"),
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
