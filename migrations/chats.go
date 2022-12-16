package migrations

type chats struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Name_chat string `json:"name_chat"`
}
