package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvirontmentVariable() {
		err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}