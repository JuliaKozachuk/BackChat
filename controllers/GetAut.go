package controllers

import (
	"fmt"
	"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/utils"
	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8/internal/util"
)

type GetAut struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// создает нового пользователя и возвращает сведения о сохраненном пользователе
func GetAuth(context *gin.Context) {
	var get GetAut

	if err := context.ShouldBindJSON(&get); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := migrations.Users{
		Email:    get.Email,
		Password: get.Password,
	}

	savedUser, err := user.SaveUser()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}
func Login(context *gin.Context) {
	var input GetAut

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(&input)
	fmt.Println("helooooo")
	fmt.Println(input.Email)

	user, err := migrations.FindUserByUsername(input.Email)

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
