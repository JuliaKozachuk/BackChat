package routers

import (
	"github.com/JuliaKozachuk/BackChat/redisconnect"
)

func Redis() {
	redisconnect.Setup()
}
