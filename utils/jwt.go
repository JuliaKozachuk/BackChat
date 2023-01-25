package utils

import (
	"os"
	"strconv"
	"time"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/golang-jwt/jwt/v4"
)

var privateKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

// Эта функция берет модель пользователя и создает JWT, содержащий идентификатор пользователя ( id),
// время выпуска токена ( iat) и дату истечения срока действия токена ( eat).
// При использовании JWT_PRIVATE_KEYпеременной среды подписанный JWT возвращается в виде строки.
func GenerateJWT(user migrations.Users) (string, error) {
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}
