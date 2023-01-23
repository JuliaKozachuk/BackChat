package routers

import (
	"fmt"
	"log"

	"net/http"
	"os"

	"github.com/JuliaKozachuk/BackChat/controllers"
	"github.com/JuliaKozachuk/BackChat/migrations"

	docs "github.com/JuliaKozachuk/BackChat/docs"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]

func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}
func Router() {
	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := route.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
			eg.POST("/signup", controllers.SignUpInput)

		}
	}

	migrations.ConnectDB(postgresUrl())
	//migrations.Missing()

	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": migrations.Alex()})

	})

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := route.Run(":9888")
	if err != nil {
		panic("[Error] failed to start Gin server due to: " + err.Error())
		return
	}

	// route.Run()

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
