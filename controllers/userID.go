package controllers

import (
	"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"

	"github.com/gin-gonic/gin"
)

func GetAllTracks(context *gin.Context) {
	var usersID []migrations.Users

	migrations.DB.Find(&usersID)

	context.JSON(http.StatusOK, gin.H{"usersID": usersID})
}
