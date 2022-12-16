package migrations

type Messages struct {
	ID_message  int    `json:"id_message"`
	ID_chat     int    `json:"id_chat"`
	ID_user     int    `json: "id_user"`
	Content     string `json:"content"`
	Date_create int    `json:"data_create"`
}
