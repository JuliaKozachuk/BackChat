package v1

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/utils"
	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8/internal/util"
)

type SignUpInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
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
	// генерация динамической строки
	//
	intermagic := uuid.New()

	magic := (intermagic.String())

	user := migrations.Users{
		Email:             get.Email,
		Password:          get.Password,
		Username:          randomstring + code,
		Verification_code: code,
		Magic:             magic,
	}

	email := get.Email

	confirm_url := fmt.Sprintf("http://localhost:9888/signupconfirm?email=%s@&magic=%s", email, magic)

	savedUser, err := user.SaveUser()
	utils.SendUserEmail(get.Email, confirm_url)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}
