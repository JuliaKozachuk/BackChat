package main

import (
	"github.com/JuliaKozachuk/BackChat/routers"
)

func main() {

	routers.Redis()
	routers.Router()
}
