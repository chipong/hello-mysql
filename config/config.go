package config

import (
	"os"
	"log"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("load failed env file")
	}
	return os.Getenv(key)
}