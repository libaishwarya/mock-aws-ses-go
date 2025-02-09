package store

import "time"

const (
	EmailSentStatusSent     = "sent"
	EmailSentStatusFail     = "failed"
	EmailSentStatusRejected = "rejected"
)

type EmailSent struct {
	Source      string    `db:"source"`
	Destination string    `db:"destination"`
	Message     any       `db:"message"`
	Status      string    `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
}

type RawEmailSent struct {
	Data      any       `db:"data"`
	Status    string    `db:"status"`
	CreatedAt time.Time `db:"created_at"`
}
