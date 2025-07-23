package domain

import "time"

type Message struct {
	ID        int64     `db:"id"`
	Sender    string    `db:"sender"`
	Receiver  string    `db:"receiver"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}
