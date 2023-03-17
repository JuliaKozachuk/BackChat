package main

import (
	"github.com/JuliaKozachuk/BackChat/redisconnect"
	"github.com/JuliaKozachuk/BackChat/routers"
)

// @title BackChat Api
// @version 1.0
// @description This is a  server BackChat.

// @host localhost:9888
// @BasePath /api/v1

////@SecurityDefinitions.apikey ApiKeyAuth
////@in header
////@name Autorization

func main() {

	redisconnect.Setup()

	routers.InitRouter()

}
