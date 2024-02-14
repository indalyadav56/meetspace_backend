package types

type AddChatGroup struct{
	Title string `json:"title" validate:"required,not_blank"`
	UserIds []string `json:"user_ids" validate:"required,not_blank"`
}

type UpdateChatGroup struct{
	RoomID string `json:"room_id" validate:"required,not_blank"`
	Title string `json:"title" validate:"omitempty,not_blank"`
	UserIds []string `json:"user_ids" validate:"omitempty,not_blank"`
}