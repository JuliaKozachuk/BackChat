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

	_ "github.com/JuliaKozachuk/BackChat/docs"
	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthorizationUser struct {
	//ID_user           uint   `json:"id_user" binding:"required"`
	Username          string `json:"username" binding:"required"`
	Password          string `json:"password"binding:"required"`
	Email             string `json:"email" binding:"required,email"`
	Verification_code string `json:"verification_code"`
}

// @Summary writes the user to the database
// @Produce json
// @Param SignUpInput body AuthorizationUser true "user"
// @Success 200  {object} migrations.Users
// @Router /example/signup [post]
// // @Security ApiKeyAuth
func SignUpInput(context *gin.Context) { // создаем нового Юзера
	var InputSignUp AuthorizationUser
	if err := context.ShouldBindJSON(&InputSignUp); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	numb := numbergenerate()
	code := strconv.Itoa(int(numb))
	InputSignUp.Verification_code = code

	fmt.Printf("user: %v", InputSignUp.Email)
	sendUserEmail(InputSignUp.Email, code) //получает агументы из функции "SignUp"

	pass := []byte(InputSignUp.Password)

	IsCorrectPassword(pass)
	InputSignUp.Password = IsCorrectPassword(pass)
	//InputSignUp.Verification_code= isCorrectVerificode(code)

	//numb := numbergenerate()
	//code := strconv.Itoa(int(numb))

	user := migrations.Users{Username: InputSignUp.Username, Email: InputSignUp.Email, Password: InputSignUp.Password, Verification_code: InputSignUp.Verification_code}
	//migrations.DB.Create(&user)
	fmt.Printf("user: %v", user)

	//user := migrations.Users{Login: InputSignUp.Login, Email: InputSignUp.Email, Password: isCorrectPassword()}
	migrations.DB.Create(&user)

}

// отправляет код подтверждения для созданного юзера
func sendUserEmail(email string, code string) {

	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/alerts/%7Balert_id%7D"
	rkey := os.Getenv("RAPID_KEY")
	rhost := os.Getenv("RAPID_HOST")

	// numb := numbergenerate()
	// code = strconv.Itoa(int(numb))

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
func IsCorrectPassword(pass []byte) string {

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
