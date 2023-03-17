package utils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

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

// функция гарантирует, что входящий запрос содержит допустимый токен в заголовке запроса.
func ValidateJWT(context *gin.Context) error {
	token, err := getToken(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

// будет использоваться для получения пользователя, связанного с предоставленным JWT, путем извлечения ключа идентификатора из проанализированного JWT и извлечения соответствующего пользователя из базы данных.
func CurrentUser(context *gin.Context) (migrations.Users, error) {
	err := ValidateJWT(context)
	if err != nil {
		return migrations.Users{}, err
	}
	token, _ := getToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))

	user, err := migrations.FindUserById(userId)
	if err != nil {
		return migrations.Users{}, err
	}
	return user, nil
}

// использует возвращенную строку маркера для анализа JWT с использованием закрытого ключа, указанного в env
func getToken(context *gin.Context) (*jwt.Token, error) {
	tokenString := getTokenFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

// извлекает маркер носителя из запроса.
func getTokenFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
