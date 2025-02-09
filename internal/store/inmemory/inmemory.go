package inmemory

import (
	"strconv"
	"time"

	"github.com/libaishwarya/mock-aws-ses-go/internal/app"
	"github.com/libaishwarya/mock-aws-ses-go/internal/store"
)

type InMemoryStore struct {
	EmailSent    []store.EmailSent
	RawEmailSent []store.RawEmailSent
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
	i.EmailSent = append(i.EmailSent, store.EmailSent{
		Source:      r.Source,
		Destination: r.Destination,
		Status:      store.EmailSentStatusSent,
		Message:     r.Message,
		CreatedAt:   time.Now().UTC(),
	})

	return strconv.Itoa(len(i.EmailSent)), nil
}

func (i *InMemoryStore) CreateRawEmailSend(r app.SendRawEmailRequest) (string, error) {
	i.RawEmailSent = append(i.RawEmailSent, store.RawEmailSent{
		Data:      r.Data,
		Status:    store.EmailSentStatusSent,
		CreatedAt: time.Now().UTC(),
	})
	return strconv.Itoa(len(i.RawEmailSent)), nil
}

func (i *InMemoryStore) ListIdentities() ([]string, error) {
	return i.Identities, nil
}

// Get the count of emails sent in the last 24 hours
func (i *InMemoryStore) GetSentEmailCount24() (int, error) {
	last24HourEmailCount := 0
	for _, email := range i.EmailSent {
		if email.CreatedAt.After(time.Now().Add(-24 * time.Hour)) {
			last24HourEmailCount++
		}
	}

	for _, email := range i.RawEmailSent {
		if email.CreatedAt.After(time.Now().Add(-24 * time.Hour)) {
			last24HourEmailCount++
		}
	}

	return last24HourEmailCount, nil
}

func (i *InMemoryStore) GetEmailStats() (map[string]int, error) {
	total := 0
	stats := make(map[string]int)
	for _, email := range i.EmailSent {
		stats[email.Status]++
		total++
	}

	for _, email := range i.RawEmailSent {
		stats[email.Status]++
		total++
	}

	stats["total"] = total

	return stats, nil
}
