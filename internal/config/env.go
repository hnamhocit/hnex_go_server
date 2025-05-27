package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	DATABASE_URL           string
	PORT                   int
	JWT_ACCESS_EXPIRES_IN  int
	JWT_REFRESH_EXPIRES_IN int
	JWT_ACCESS_SECRET      string
	JWT_REFRESH_SECRET     string
}

func LoadEnv() *Env {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	JWT_ACCESS_EXPIRES_IN, err := strconv.Atoi(os.Getenv("JWT_ACCESS_EXPIRES_IN"))
	if err != nil {
		log.Fatal("Error parsing JWT_ACCESS_EXPIRES_IN")
	}

	JWT_REFRESH_EXPIRES_IN, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRES_IN"))
	if err != nil {
		log.Fatal("Error parsing JWT_REFRESH_EXPIRES_IN")
	}

	PORT, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Error parsing PORT")
	}

	return &Env{
		DATABASE_URL:           os.Getenv("DATABASE_URL"),
		PORT:                   PORT,
		JWT_ACCESS_EXPIRES_IN:  JWT_ACCESS_EXPIRES_IN,
		JWT_REFRESH_EXPIRES_IN: JWT_REFRESH_EXPIRES_IN,
		JWT_ACCESS_SECRET:      os.Getenv("JWT_ACCESS_SECRET"),
		JWT_REFRESH_SECRET:     os.Getenv("JWT_REFRESH_SECRET"),
	}
}
