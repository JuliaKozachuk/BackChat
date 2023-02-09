package v1

import (
	"fmt"
	"net/http"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/JuliaKozachuk/BackChat/utils"
	"github.com/gin-gonic/gin"
	//"github.com/go-redis/redis/v8/internal/util"
)

type SignInInput struct {
	Email             string `json:"email" binding:"required"`
	Password          string `json:"password" binding:"required"`
	Verification_code string `json:"verification_code"`
}

// @Summary User login
// @Description  User login to the system by mail, password, verification code
// @ID login
// @Tags         auth
// // @Accept json
// @Produce json
// @Param input body SignInInput true " login user"
// @Success 200 {string} string "jwt"
// @Router /example/SingIn [post]
func SingIn(context *gin.Context) {
	var input SignInInput

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
