package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

type SignUpConfirmInput struct {
	Email string `form:"email"`
	Magic string `form:"magic"`
}

// HTTP GET with query params
func SignUpConfirm(context *gin.Context) {
	// Step 1. Получаем Email и Magic
	var input SignUpConfirmInput
	if context.Bind(&input) == nil {
		log.Println(input.Email)
		log.Println(input.Magic)
		log.Println("Binding success...............")
	} else {
		log.Println("Binding failed...............")
	}
	/////////////////////////////////////////////////////

	// Step 2. Ищем в базе юзера по емейл и magic
	// если нет то выход ошибка

	// Step 3. Обновляем юзеру статус status="active"

	context.String(200, "Success")
}
