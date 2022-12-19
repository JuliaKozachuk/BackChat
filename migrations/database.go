package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB(postgres_data string) {

	db, err := gorm.Open("postgres", postgres_data)
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Chat_users{})
	db.AutoMigrate(&chats{})
	db.AutoMigrate(&Messages{})
	db.AutoMigrate(&Roles{})
	db.AutoMigrate(&Users{})

	DB = db
}
