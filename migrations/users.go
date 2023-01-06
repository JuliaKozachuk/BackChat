package migrations

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//var DB *gorm.DB

type Users struct {
	//gorm.Model
	ID_user           uint   `json:"id_user" gorm:"primary_key"`
	Username          string `json:"username addin"gorm:"unique"`
	Password          string `json:"password"`
	Email             string `json:"email" gorm:"unique"`
	Verification_code string `json:"verification_code"`
}
