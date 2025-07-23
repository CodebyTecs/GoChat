package postgres

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect() *sqlx.DB {
	dsn := os.Getenv("DB_DSN")
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalln("Error connecting to database:", err)
	}
	return db
}
