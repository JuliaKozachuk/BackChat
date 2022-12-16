package main

import (
	"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	migrations.ConnectDB() // new

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Успешное соединение"})
	})

	route.Run()
}
