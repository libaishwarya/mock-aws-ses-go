package inmemory

import (
	"strconv"

	"github.com/libaishwarya/mock-aws-ses-go/internal/app"
)

type InMemoryStore struct {
	EmailSent    []app.SendEmailRequest
	RawEmailSent []app.SendRawEmailRequest
	Identities   []string
}

func NewInMemoryStore() *InMemoryStore {
	return &InMemoryStore{
		Identities: []string{
			"test@demtech.ai",
			"test@test.com",
			"test@gmail.com",
		},
	}
}

func (i *InMemoryStore) CreateEmailSend(r app.SendEmailRequest) (string, error) {
	i.EmailSent = append(i.EmailSent, r)
	return strconv.Itoa(len(i.EmailSent)), nil
}

func (i *InMemoryStore) CreateRawEmailSend(r app.SendRawEmailRequest) (string, error) {
	i.RawEmailSent = append(i.RawEmailSent, r)
	return strconv.Itoa(len(i.RawEmailSent)), nil
}

func (i *InMemoryStore) ListIdentities() ([]string, error) {
	return i.Identities, nil
}
