package repository

import (
	"GoChat/internal/domain"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	SaveUser(db *sqlx.DB, u domain.User) error
	GetUsername(db *sqlx.DB, u domain.User) (string, error)
	IsUserExist(db *sqlx.DB, u domain.User) (bool, error)
}
