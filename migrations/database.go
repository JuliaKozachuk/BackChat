package migrations

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=backchat password=qwerty sslmode=disable")
	if err != nil {
		panic("Не удалось подключиться к базе данных")
	}
	db.AutoMigrate(&Chat_users{})
	db.AutoMigrate(&chats{})
	db.AutoMigrate(&Messages{})
	db.AutoMigrate(&Roles{})
	db.AutoMigrate(&Users{})

	return db
}
