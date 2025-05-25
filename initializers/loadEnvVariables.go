package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	if godotenv.Load() != nil {
		log.Fatal("Failed to read environment variables")
	}
}
