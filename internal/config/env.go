package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DATABASE_URL string
	PORT         int
}

func LoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error parsing PORT")
	}

	return &Env{
		DATABASE_URL: os.Getenv("DATABASE_URL"),
		PORT:         PORT,
	}
}
