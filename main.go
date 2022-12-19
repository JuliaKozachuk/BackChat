package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/JuliaKozachuk/BackChat/controllers"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	route := gin.Default()

	migrations.ConnectDB()

	//route.GET("/", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{"message": "Успешное соединение"})

	//})
	route.GET("/userID", controllers.GetAllUsers)
	route.POST("/user", controllers.CreateUser)

	route.Run()
}
