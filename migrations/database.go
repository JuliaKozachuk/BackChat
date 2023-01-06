package migrations

import (
	//"github.com/JuliaKozachuk/BackChat/migrations"

	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"gorm.io/gorm"
)

type inpus struct {
	username string
	email    string
}

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
	mg := db.Where("email=?", "uliakozacuk649@gmai.com").First(&Users{})
	fmt.Printf("%s", mg)

	DB = db
	//db.CreateTable(&Users{})

}
