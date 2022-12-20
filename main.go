package main

import (
	"fmt"
	"net/http"

	"log"
	"os"

	"github.com/JuliaKozachuk/BackChat/controllers"
	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/redisconnect"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

//var ctx = context.Background()

func main() {

	redisconnect.ExampleClient()

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

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // no password set
	// 	DB:       0,  // use default DB
	// })

	// err := rdb.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

	// val, err := rdb.Get(ctx, "key").Result()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("key", val)

	migrations.ConnectDB(postgres_data)

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Успешное соединение"})
	})
	route.GET("/userID", controllers.GetAllUsers)
	route.POST("/user", controllers.CreateUser)

	route.Run()
}
