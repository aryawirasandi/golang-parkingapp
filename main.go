package main

import (
	"log"
	"os"

	"github.com/aryawirasandi/parking-app/database"
	"github.com/aryawirasandi/parking-app/service"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	db := os.Getenv("DB_DATABASENAME")
	port := os.Getenv("DB_PORT")
	portRunning := os.Getenv("PORT_RUNNING")
	env := &database.Env{
		Host:     host,
		Username: username,
		Password: password,
		Database: db,
		Port:     port,
	}

	service.Mapper(env).Run(portRunning)
}
