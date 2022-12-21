package migrations

type Users struct {
	ID_user  uint   `json:"id_user" gorm:"primary_key"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
}
