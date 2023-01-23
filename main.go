package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/redisconnect"
	"github.com/JuliaKozachuk/BackChat/routers"
)

// @title BACKCHAT Api
func main() {

	redisconnect.Setup()

	err := routers.InitRouter().Run(":9888")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())

	}

	migrations.ConnectDB(postgresUrl())
}
func postgresUrl() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	sslmode := os.Getenv("POSTRES_SSLMODE")

	postgres_data := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s ", host, port, user, dbname, password, sslmode)

	return postgres_data
}
