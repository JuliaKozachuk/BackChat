package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func User(context *gin.Context) {
	access_token := context.Request.Header["Autorization"]

	fmt.Println(access_token)
	context.JSON(200, gin.H{
		"token_data": context.Request.Header["Autorization"],
	})

}
