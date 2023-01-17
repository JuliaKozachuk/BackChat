package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JuliaKozachuk/BackChat/controllers"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/redisconnect"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	redisconnect.ExampleClient()

	route := gin.Default()

	migrations.ConnectDB(postgresUrl())
	//migrations.Missing()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Успешное соединение"})
	})

	route.GET("/userID", controllers.GetAllUsers)
	//route.GET("/users:id", controllers.GetUser)
	route.POST("/user", controllers.CreateUser)

	route.POST("/signup", controllers.SignUpInput)

	route.DELETE("/del", controllers.DeleteUser)
	route.POST("/login", controllers.Login)

	err := route.Run(":9888")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}

	// route.Run()

}

// подключение к postresql
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
