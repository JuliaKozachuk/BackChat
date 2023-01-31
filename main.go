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
	//  if err != nil {
	//  	panic("[Error] failed to start Gin server due to: " + err.Error())
	//  }
	// 	d := controllers.Logingenerate(150)
	// 	fmt.Println(d)
}
