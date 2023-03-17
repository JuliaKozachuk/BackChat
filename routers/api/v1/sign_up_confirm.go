package v1

import (
	"log"
	"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/gin-gonic/gin"
)

type SignUpConfirmInput struct {
	Email string `form:"email"`
	Magic string `form:"magic"`
}

func SignUpConfirm(context *gin.Context) {

	var input SignUpConfirmInput
	if context.Bind(&input) == nil {
		log.Println(input.Email)
		log.Println(input.Magic)
		log.Println("Binding success...............")
	} else {
		log.Println("Binding failed...............")
	}
	/////////////////////////////////////////////////////

	user, err := migrations.FindUserByEmail(input.Email, input.Magic)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updateUser, err := user.UpdateUser()
	context.JSON(http.StatusCreated, gin.H{"user": updateUser})

	context.String(200, "Success")
}
