package GoChat

import (
	"GoChat/internal/infrastructure/db/postgres"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func main() {
	_ = godotenv.Load()

	db := postgres.Connect()
}
