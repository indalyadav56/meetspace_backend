package types

type AddChatGroup struct{
	Title string `json:"title"`
	UserIds []string `json:"user_ids"`
}