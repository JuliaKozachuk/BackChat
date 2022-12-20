package connectDB

import (
	"fmt"
	"log"
	"os"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/joho/godotenv"
)

func connectBackChat() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	sslmode := os.Getenv("DB_SSLMODE")

	postgres_data := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s ", host, port, user, dbname, password, sslmode)
	migrations.ConnectDB(postgres_data)
}
