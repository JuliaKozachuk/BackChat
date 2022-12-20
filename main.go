package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JuliaKozachuk/BackChat/connectredis"
	"github.com/joho/godotenv"

	"github.com/JuliaKozachuk/BackChat/controllers"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	connectredis.ExampleClient()

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

	route := gin.Default()

	migrations.ConnectDB(postgres_data)

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Успешное соединение"})

	})
	route.GET("/userID", controllers.GetAllUsers)
	route.POST("/user", controllers.CreateUser)

	route.Run()
}
