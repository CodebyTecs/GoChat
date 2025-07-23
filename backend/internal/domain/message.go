package domain

type Message struct {
	Sender   string `db:"sender"`
	Receiver string `db:"receiver"`
	Text     string `db:"text"`
}
