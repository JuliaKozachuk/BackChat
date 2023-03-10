package controllers

import (

	//"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"

	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateUserInput struct {
	ID_user  int    `json:"user_id" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// @description getting all users
// @Produce json
// @Accept json
// @Success 200  {object} []migrations.Users
// @Router /example/userID [get]
func GetAllUsers(context *gin.Context) {
	var usersID []migrations.Users

	migrations.DB.Find(&usersID)

	context.JSON(http.StatusOK, gin.H{"ID_user": usersID})
}
func GetUser(context *gin.Context) {

	var usersID migrations.Users
	if err := migrations.DB.Where("id_user = ?", context.Param("ID_user")).First(&usersID).Error; err != nil {
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

	user := migrations.Users{Username: input.Username, Password: input.Password, Email: input.Email}
	migrations.DB.Create(&user)

	context.JSON(http.StatusOK, gin.H{"user": user})
}
func DeleteUser(context *gin.Context) {

	var user migrations.Users
	if err := migrations.DB.Where("id_user=id_user", context.Param("id_user")).First(&user).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"users": "Запись не существует"})
		return
	}

	migrations.DB.Delete(&user)

	context.JSON(http.StatusOK, gin.H{"users": true})
}

// $HOME
// /home/julia
// ~/

// /home/julia/projects

// export GOPATH=$HOME/projects:/usr/bin/go
