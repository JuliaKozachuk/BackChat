package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/utils"
	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8/internal/util"
)

type GetAut struct {
	Username          string `swaggerignore:"true" json:"username" `
	Email             string `json:"email" binding:"required"`
	Password          string `json:"password" binding:"required"`
	Verification_code string `swaggerignore:"true" json:"verification_code"` //swaggerignore, чтобы исключить поле из зоны видимости сваггера
}
type LoginUp struct {
	Email             string `json:"email" binding:"required"`
	Password          string `json:"password" binding:"required"`
	Verification_code string `json:"verification_code"`
}

// @Summary writes the user to the database
// @Description  register a new user
// @Tags         Сreate a new account
// @Produce json
// @Param get body GetAut true "user"
// @Success 201 {object} GetAut
// @Failure      500 "user registration failed"
// @Failure      400 "email is not unique"
// @Router /example/getauth [post]
// // @Security ApiKeyAuth
// создает нового пользователя и возвращает сведения о сохраненном пользователе
func GetAuth(context *gin.Context) {
	var get GetAut

	if err := context.ShouldBindJSON(&get); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	numb := numbergenerate()
	code := strconv.Itoa(int(numb))
	get.Verification_code = code

	get.Username = "BlackUser-" + code

	user := migrations.Users{
		Email:             get.Email,
		Password:          get.Password,
		Username:          get.Username,
		Verification_code: get.Verification_code,
	}

	savedUser, err := user.SaveUser()
	sendUserEmail(get.Email, code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}
func Login(context *gin.Context) {
	var input LoginUp

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(&input)
	fmt.Println("helooooo")
	fmt.Println(input.Email)

	user, err := migrations.FindUserByUsername(input.Email, input.Verification_code)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utils.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}
