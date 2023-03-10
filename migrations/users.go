package migrations

import (
	//"log"

	//"fmt"

	"html"

	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

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
	Username          string `gorm: not null; unique" json:"username"`
	Password          string `gorm:"size:255;not null;" json:"-"` //Опривязка JSON для Password поля — -. Это гарантирует, что пароль пользователя не будет возвращен в ответе JSON.
	Email             string `gorm:"size:255;not null;unique" json:"email"`
	Verification_code string `json:"-"`
}

// добавляет в базу нового пользователя
func (u *Users) SaveUser() (*Users, error) {

	var err error
	err = DB.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}
	return u, nil
}

// хэширует пароль, перед сохранением, предварительно обрезав все возможные пробелы
func (user *Users) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	//nameuser,err := user.Username
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	//user.Username = string(user.Username)
	return nil
}

// генерируется хеш для предоставленного открытого пароля и сравнивается с хэшем пароля пользователя. Если они не совпадают, возвращается ошибка.
func (user *Users) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

// принимает  email пользователя и запрашивает базу данных, чтобы найти соответствующего пользователя.
func FindUserByUsername(email, verification_code string) (Users, error) {
	var user Users
	err := DB.Where("email=?", email).Find(&user).Error
	if err != nil {
		return Users{}, err
	}
	ver := DB.Where("verification_code=?", verification_code).Find(&user).Error
	if ver != nil {
		return Users{}, err
	}
	return user, nil
}
