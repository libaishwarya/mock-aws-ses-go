package ses

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/libaishwarya/mock-aws-ses-go/internal/app"
	"github.com/libaishwarya/mock-aws-ses-go/internal/serror"
	"github.com/libaishwarya/mock-aws-ses-go/internal/store"
)

type SESHandler struct {
	Store store.Store
}

func NewSESHandler(store store.Store) *SESHandler {
	return &SESHandler{
		Store: store,
	}
}

func AttachRoutes(r *gin.Engine, store store.Store) {
	sesHandler := NewSESHandler(store)
	r.POST("/v1/sendEmail", sesHandler.identities(), sesHandler.SendEmail)
	r.POST("/v1/sendRawEmail", sesHandler.identities(), sesHandler.SendRawEmail)
	r.GET("/v1/listIdentities", sesHandler.identities(), sesHandler.ListIdentities)
	r.GET("/v1/getSendQuota", sesHandler.GetSendQuota)
	// r.GET("/v1/stats", sesHandler.GetStats)

}

func (s *SESHandler) SendEmail(c *gin.Context) {
	var req app.SendEmailRequest
	// For validating Identities
	req.Identities = c.MustGet("identities").([]string)

	if err := c.ShouldBindJSON(&req); err != nil {
		serror.HandleBindError(c, err, req)
		return
	}

	if err := req.Validate(); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageID, err := s.Store.CreateEmailSend(req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// Simulate email sending
	c.JSON(http.StatusOK, app.SendEmailResponse{
		MessageId: messageID,
	})
}

func (s *SESHandler) SendRawEmail(c *gin.Context) {
	var req app.SendRawEmailRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		serror.HandleBindError(c, err, req)
		return
	}

	messageID, err := s.Store.CreateRawEmailSend(req)
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	// Simulate email sending
	c.JSON(http.StatusOK, app.SendEmailResponse{
		MessageId: messageID,
	})
}

func (s *SESHandler) ListIdentities(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Identities": c.MustGet("identities"),
	})
}

func (s *SESHandler) identities() gin.HandlerFunc {
	return func(c *gin.Context) {
		identities, err := s.Store.ListIdentities()
		if err != nil {
			c.Error(err)
			c.Status(http.StatusInternalServerError)
			return
		}

		c.Set("identities", identities)
	}
}

func (s *SESHandler) GetSendQuota(c *gin.Context) {
	count, err := s.Store.GetSentEmailCount24()
	if err != nil {
		c.Error(err)
		c.Status(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Max24HourSend":   10000,
		"MaxSendRate":     14,
		"SentLast24Hours": count,
	})
}
