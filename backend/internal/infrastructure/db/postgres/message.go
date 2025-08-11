package postgres

import (
	"GoChat/internal/domain"

	"github.com/jmoiron/sqlx"
)

func SaveMessage(db *sqlx.DB, msg domain.Message) error {
	_, err := db.Exec("INSERT INTO messages (sender, receiver, text) VALUES ($1, $2, $3)", msg.Sender, msg.Receiver, msg.Text)
	return err
}

func GetMessagesBetweenUsers(db *sqlx.DB, user1 string, user2 string) ([]domain.Message, error) {
	var messages []domain.Message
	err := db.Select(&messages, "SELECT * FROM messages "+
		"WHERE (sender = $1 AND receiver = $2) OR (sender = $2 AND receiver = $1) "+
		"ORDER BY created_at", user1, user2)

	if len(messages) == 0 {
		return nil, err
	}

	return messages, err
}

func GetLastMessagesForUser(db *sqlx.DB, u domain.User) ([]domain.Message, error) {
	var messages []domain.Message
	err := db.Select(&messages, "SELECT * FROM messages "+
		"WHERE receiver = $1"+
		"ORDER BY created_at DESC"+
		"LIMIT 20", u.Username)

	if len(messages) == 0 {
		return nil, err
	}

	return messages, err
}
