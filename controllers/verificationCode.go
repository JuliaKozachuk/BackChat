package controllers

import (

	//"net/http"

	"crypto/rand"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"

	_ "github.com/JuliaKozachuk/BackChat/docs"
)

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
	safeNum, err := rand.Int(rand.Reader, big.NewInt(80000000000000))
	if err != nil {
		fmt.Println(err)
	}
	return safeNum.Int64()

}
