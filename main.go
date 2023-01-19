package main

import (
	"github.com/JuliaKozachuk/BackChat/redisconnect"
	"github.com/JuliaKozachuk/BackChat/routers"
)

func main() {

	redisconnect.Setup()

	routers.Router()
}
