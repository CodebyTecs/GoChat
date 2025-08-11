package postgres

import (
	"GoChat/internal/domain"

	"github.com/jmoiron/sqlx"
)

func SaveUser(db *sqlx.DB, u domain.User) error {
	_, err := db.Exec("INSERT INTO users (username, password) VALUES ($1, $2) ON CONFLICT DO NOTHING", u.Username, u.Password)
	return err
}

func GetUsername(db *sqlx.DB, u domain.User) (string, error) {
	var result string
	err := db.Get(&result, "SELECT username FROM users WHERE username = $1", u.Username)
	return result, err
}

func GetUserByUsername(db *sqlx.DB, username string) (domain.User, error) {
	var user domain.User
	err := db.Get(&user, "SELECT id, username, password FROM users WHERE username = $1", username)
	return user, err
}

func IsUserExist(db *sqlx.DB, username string) (bool, error) {
	var exists bool
	err := db.Get(&exists, "SELECT username FROM users WHERE username = $1", username)
	return exists, err
}
