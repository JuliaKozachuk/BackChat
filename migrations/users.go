package migrations

import (
	//"log"

	//"fmt"

	//"html"

	//"BackChat/utils/token"

	token "github.com/JuliaKozachuk/BackChat/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	gorm.Model
	ID_user           uint   `json:"id_user" gorm:"primary_key"`
	Username          string `json:"username addin"gorm:"unique"`
	Password          string `json:"password"`
	Email             string `json:"email" gorm:"unique"`
	Verification_code string `json:"verification_code"`
}

func (u *Users) SaveUser() (*Users, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
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
