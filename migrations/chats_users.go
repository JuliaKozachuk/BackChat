package migrations

type Chat_users struct {
	Chat_id int `json:"chat_id"`
	User_id int `json:"user_id"`
	Role_id int `json:"role_id"`
}
