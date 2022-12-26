package controllers

import (
	"fmt"

	"github.com/JuliaKozachuk/BackChat/controllers"
	"golang.org/x/crypto/bcrypt"
)

func isCorrectPassword(password []byte) {
	password = []byte(controllers.SignUpInput(password))

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Сравнение пароля с хэшем
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // если пароли совпадут, то err выдаст <nil>
}
