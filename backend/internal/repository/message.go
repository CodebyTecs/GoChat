package repository

import (
	"GoChat/internal/domain"
	"github.com/jmoiron/sqlx"
)

type MessageRepository interface {
	SaveMessage(db *sqlx.DB, msg domain.Message) error
	GetMessagesBetweenUsers(db *sqlx.DB, user1 string, user2 string) ([]domain.Message, error)
	GetLastMessagesForUser(db *sqlx.DB, u domain.User) ([]domain.Message, error)
}
