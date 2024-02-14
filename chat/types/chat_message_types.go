package types

import "meetspace_backend/user/models"


type CreateChatRequestBody struct {
    ChatRoomId string `json:"chat_room_id"`
    Content string `json:"content"`
    CurrentUserId string `json:"currentUserId,omitempty"`
}

type GetChatMessageRequestBody struct {
    ChatRoomId string `json:"chat_room_id" validate:"required,not_blank"`
}


type SingleChatMessageResponse struct {
    ID string `json:"id"`
    Content string `json:"content"`
    ChatRoomId string `json:"chat_room_id"`
    Sender *models.User `json:"sender_user"`
}

type ChatMessageResponse struct {
    TimeStamp string `json:"timestamp"`
    ChatMessage []SingleChatMessageResponse `json:"chat_message"`
}