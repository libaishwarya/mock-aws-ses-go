package store

import "github.com/libaishwarya/mock-aws-ses-go/internal/app"

type Store interface {
	CreateEmailSend(app.SendEmailRequest) (string, error)
	CreateRawEmailSend(app.SendRawEmailRequest) (string, error)
	ListIdentities() ([]string, error)
}
