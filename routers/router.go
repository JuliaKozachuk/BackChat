package routers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/JuliaKozachuk/BackChat/controllers"
	"github.com/JuliaKozachuk/BackChat/docs"
	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/joho/godotenv"

	//	_ "github.com/JuliaKozachuk/BackChat/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
func InitRouter() {
	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
			eg.POST("/signup", controllers.SignUpInput)
			eg.GET("/userID", controllers.GetAllUsers)
		}

	}
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": migrations.Alex()})
	})
	route.GET("/userID", controllers.GetAllUsers)

	migrations.ConnectDB(postgresUrl())

	//route.POST("/signup", controllers.SignUpInput)
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	//route.POST("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	err := route.Run(":9888")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}

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
