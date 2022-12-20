package controllers

import (
	"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	//ID_user  int    `json:"user_id" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

func GetAllUsers(context *gin.Context) {
	var usersID []migrations.Users

	migrations.DB.Find(&usersID)

	context.JSON(http.StatusOK, gin.H{"ID_user": usersID})
}
func GetUser(context *gin.Context) {
	// Проверяем имеется ли запись
	var usersID migrations.Users
	if err := migrations.DB.Where("id = ?", context.Param("ID_user")).First(&usersID).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"usersID": usersID})
}

func CreateUser(context *gin.Context) {
	var input CreateUserInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := migrations.Users{Login: input.Login, Password: input.Password, Email: input.Email}
	migrations.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})
}
