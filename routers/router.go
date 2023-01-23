package routers

import (
	"net/http"

	"github.com/JuliaKozachuk/BackChat/controllers"
	"github.com/JuliaKozachuk/BackChat/migrations"

	docs "github.com/JuliaKozachuk/BackChat/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title BACKCHAT Api
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
func InitRouter() *gin.Engine {
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
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": migrations.Alex()})
	})
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return route
}
