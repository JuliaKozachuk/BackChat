package controllers

import (

	//"net/http"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthorizationUser struct {
	//ID_user           uint   `json:"id_user" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password"binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	//Verification_code string `json:"email" binding:"required"`
}

// создаем нового Юзера
func SignUpInput(context *gin.Context) {
	var InputSignUp AuthorizationUser
	if err := context.ShouldBindJSON(&InputSignUp); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("user: %v", InputSignUp.Email)
	sendUserEmail(InputSignUp.Email, "code") //получает агументы из функции "SignUp"

	pass := []byte(InputSignUp.Password)

	isCorrectPassword(pass)
	InputSignUp.Password = isCorrectPassword(pass)
	user := migrations.Users{Login: InputSignUp.Login, Email: InputSignUp.Email, Password: InputSignUp.Password}
	//migrations.DB.Create(&user)
	fmt.Printf("user: %v", user)

	//user := migrations.Users{Login: InputSignUp.Login, Email: InputSignUp.Email, Password: isCorrectPassword()}
	migrations.DB.Create(&user)

}

// отправляет код подтверждения для созданного юзера
func sendUserEmail(email string, code string) {

	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	rkey := os.Getenv("RAPID_KEY")
	rhost := os.Getenv("RAPID_HOST")

	numb := numbergenerate()
	code = strconv.Itoa(int(numb))

	payload := strings.NewReader("{\r\"personalizations\": [\r{\r\"to\": [\r{\r\"email\": \"" + email + "\"\r}\r],\r\"subject\": \"password:" + code + "\"\r}\r],\r\"from\": {\r\"email\": \"uliakozacuk649@gmail.com\"\r},\r\"content\": [\r{\r\"type\": \"text/plain\",\r\"value\": \"password:" + code + "\"\r}\r]\r}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", rkey)
	req.Header.Add("X-RapidAPI-Host", rhost)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

// генерация рандомных чисел для кода
func numbergenerate() int64 {
	safeNum, err := rand.Int(rand.Reader, big.NewInt(800000))
	if err != nil {
		fmt.Println(err)
	}
	return safeNum.Int64()

}
func isCorrectPassword(pass []byte) string {
	//inputpassword := SignUpInput(password)
	//pass = []byte(inputpassword)

	// Хеширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(hashedPassword))

	// Сравнение пароля с хэшем
	err = bcrypt.CompareHashAndPassword(hashedPassword, pass)
	fmt.Println(err) // если пароли совпадут, то err выдаст <nil>
	return string(hashedPassword)

}
