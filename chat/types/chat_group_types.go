package types

type AddChatGroup struct{
	Title string `json:"title" validate:"required,not_blank"`
	UserIds []string `json:"user_ids" validate:"required,not_blank"`
}