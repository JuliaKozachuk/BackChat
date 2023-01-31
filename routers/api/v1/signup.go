package v1

import (
	"net/http"
	"strconv"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/utils"
	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8/internal/util"
)

type SignUpInput struct {
	//Username          string `swaggerignore:"true" json:"username" `
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	//Verification_code string `swaggerignore:"true" json:"verification_code"` //swaggerignore, чтобы исключить поле из зоны видимости сваггера
}

// @Summary writes the user to the database
// @Description  register a new user
// @Tags         Сreate a new account
// @Accept json
// @Produce json
// @Param get body SignUpInput true "user"
// @Success 201 {object} SignUpInput
// @Failure      500 "user registration failed"
// @Failure      400 "email is not unique"
// @Router /example/SignUp [post]
// // @Security ApiKeyAuth
// создает нового пользователя и возвращает сведения о сохраненном пользователе
func SignUp(context *gin.Context) {
	var get SignUpInput

	if err := context.ShouldBindJSON(&get); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	numb := utils.Numbergenerate()
	code := strconv.Itoa(int(numb))
	randomstring := utils.GenerateString()
	//get.Verification_code = code

	//get.Username = "BlackUser-" + code

	user := migrations.Users{
		Email:             get.Email,
		Password:          get.Password,
		Username:          randomstring + code,
		Verification_code: code,
	}

	savedUser, err := user.SaveUser()
	utils.SendUserEmail(get.Email, code)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}
