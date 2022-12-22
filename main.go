package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

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
	generate.numbergenerate()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Успешное соединение"})
	})
	route.GET("/userID", controllers.GetAllUsers)
	route.GET("/users:id", controllers.GetUser)
	route.POST("/user", controllers.CreateUser)
	route.POST("/Post", sendEmail)
	route.DELETE("/del", controllers.DeleteUser)

	route.Run()
}

func postgresUrl() string {
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

	return postgres_data
}
func sendEmail(context *gin.Context) {
	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	rkey := os.Getenv("RAPID_KEY")
	rhost := os.Getenv("RAPID_HOST")
	//rapid_key := fmt.Sprintf("rkey=%s", rkey)
	//rapid_host := fmt.Sprintf("rhost=%s", rhost)

	payload := strings.NewReader("{\r\"personalizations\": [\r {\r\"to\": [\r{\r  \"email\": \"sunrise3323@gmail.com\"\r  }\r ],\r \"subject\": \"generation.numbergeneration=generate.numbergenerate\"\r }\r],\r \"from\":{\r \"email\": \"uliakozacuk649@gmail.com\"\r},\r \"content\": [\r {\r  \"type\": \"text/plain\",\r  \"value\": \"generation.numbergeneration=generate.numbergenerate\"\r }\r]\r}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", rkey)
	req.Header.Add("X-RapidAPI-Host", rhost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
