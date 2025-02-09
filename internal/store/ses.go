package store

import "time"

type EmailSent struct {
	Source      string    `db:"source"`
	Destination string    `db:"destination"`
	Message     any       `db:"message"`
	CreatedAt   time.Time `db:"created_at"`
}

type RawEmailSent struct {
	Data      any       `db:"data"`
	CreatedAt time.Time `db:"created_at"`
}
