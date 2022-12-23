package controllers

import (

	//"net/http"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/JuliaKozachuk/BackChat/migrations"
	"github.com/gin-gonic/gin"
)

type AuthorizationUser struct {
	//ID_user           uint   `json:"id_user" binding:"required"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password"`
	Email    string `json:"email" binding:"required,email"`
	//Verification_code string `json:"email" binding:"required"`
}

func SignUp(context *gin.Context) {
	var Email AuthorizationUser
	if err := context.ShouldBindJSON(&Email); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("user: %v", Email.Email)

	user := migrations.Users{Login: Email.Login, Email: Email.Email}
	migrations.DB.Create(&user)
	fmt.Printf("user: %v", user)

	sendUserEmail(Email.Email, "77777777")

}

func sendUserEmail(email string, code string) {

	url := "https://rapidprod-sendgrid-v1.p.rapidapi.com/mail/send"
	rkey := os.Getenv("RAPID_KEY")
	rhost := os.Getenv("RAPID_HOST")
	//rapid_key := fmt.Sprintf("rkey=%s", rkey)
	//rapid_host := fmt.Sprintf("rhost=%s", rhost)
	// numb := numbergenerate()
	// autoriseUser := controllers.AuthorizationUser()
	// email := autoriseUser

	//payload:=strings.NewReader(`{"personalizations": [{"to": [{"email": "uliakozacuk649@gmail.com"}],"subject": keynumb}],"from": {"email": "sunrise3323@gmail.com"},"content": [{"type": "text/plain","value": "Hello, World!"}]}`)
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
