package types

import "meetspace_backend/user/models"


type CreateChatRequestBody struct {
    ChatRoomId string `json:"chat_room_id"`
    Content string `json:"content"`
    CurrentUserId string `json:"currentUserId,omitempty"`
}

type GetChatMessageRequestBody struct {
    ChatRoomId string `json:"chat_room_id"`
    CurrentUserId string `json:"current_user_id,omitempty"`
}


type SingleChatMessageResponse struct {
    Content string `json:"content"`
    ChatRoomId string `json:"chat_room_Id"`
    Sender models.User `json:"sender"`
}

type ChatMessageResponse struct {
    TimeStamp string `json:"timestamp"`
    ChatMessage []SingleChatMessageResponse `json:"chat_message"`
}