package migrations

import (
	//"log"

	//"fmt"

	//"html"

	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Model struct {
	ID        uint       `gorm:"primary_key;column:id" json:"id"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}
type Users struct {
	Model
	ID_user           uint   `json:"id_user" gorm:"primary_key"`
	Username          string `json:"username addin"gorm:"unique"`
	Password          string `json:"password"`
	Email             string `json:"email" gorm:"unique"`
	Verification_code string `json:"verification_code"`
}

// func (u *Users) SaveUser() (*Users, error) {

// 	var err error
// 	err = DB.Create(&u).Error
// 	if err != nil {
// 		return &Users{}, err
// 	}
// 	return u, nil
// }
// func VerifyPassword(password, hashedPassword string) error {
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
// }

// func LoginCheck(username string, password string) (string, error) {

// 	var err error

// 	u := Users{}

// 	err = DB.Model(Users{}).Where("username = ?", username).Take(&u).Error

// 	if err != nil {
// 		return "", err
// 	}

// 	err = VerifyPassword(password, u.Password)

// 	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
// 		return "", err
// 	}

// 	token, err := token.GenerateToken(u.ID)

// 	if err != nil {
// 		return "", err
// 	}

// 	return token, nil

// }
