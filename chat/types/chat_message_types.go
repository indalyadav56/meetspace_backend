package types

import "meetspace_backend/user/models"


type CreateChatRequestBody struct {
    Content string `json:"content"`
    RoomID string `json:"room_id"`
    RecieverUserID string `json:"receiver_user_id,ommitempty"`
}

type GetChatMessageRequestBody struct {
    ChatRoomId string `json:"chat_room_id" validate:"required,not_blank"`
}


type SingleChatMessageResponse struct {
    ID string `json:"id"`
    Content string `json:"content"`
    ChatRoomId string `json:"chat_room_id"`
    Sender *models.User `json:"sender_user"`
    CreatedAt string `json:"created_at"`
}

type ChatMessageResponse struct {
    TimeStamp string `json:"timestamp"`
    ChatMessage []SingleChatMessageResponse `json:"chat_message"`
}