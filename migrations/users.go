package migrations

import (
	//"log"

	//"fmt"

	"html"

	//"BackChat/utils/token"
	"github.com/JuliaKozachuk/BackChat/token"

	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

// var DB *gorm.DB
type Token struct {
	UserId uint
	jwt.StandardClaims
}

type Users struct {
	gorm.Model
	ID_user           uint   `json:"id_user" gorm:"primary_key"`
	Username          string `json:"username addin"gorm:"unique"`
	Password          string `json:"password"`
	Email             string `json:"email" gorm:"unique"`
	Verification_code string `json:"verification_code"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (string, error) {

	var err error

	u := Users{}

	err = DB.Model(Users{}).Where("username = ?", username).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Users) SaveUser() (*Users, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}

func (u *Users) BeforeSave() error {

	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	return nil

}
