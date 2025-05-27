package main

import (
	"log"
	"os"

	"hnex.com/internal/api"
	"hnex.com/internal/config"
)

func main() {
	env := config.LoadEnv()
	db, err := config.ConnectDB(env.DATABASE_URL)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal("Error getting hostname: ", err)
	}

	api.Start(env, db, hostname)
}
