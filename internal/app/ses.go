package app

import (
	"fmt"
	"slices"
)

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

	// Used only for validation.
	Identities []string `json:"-"`
}

// Validate validates the SendEmailRequest (custom validation).
func (r SendEmailRequest) Validate() error {
	if !slices.Contains(r.Identities, r.Source) {
		return fmt.Errorf("validation failed: invalid sender")
	}

	return nil
}

type SendEmailResponse struct {
	MessageId string `json:"messageId"`
}

type SendRawEmailRequest struct {
	Data string `json:"data" binding:"base64,required"`

	// TODO: Add additional options fields later.
	ReturnPath string `json:"_"`
}
